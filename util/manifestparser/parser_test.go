package manifestparser_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	. "code.cloudfoundry.org/cli/util/manifestparser"

	"github.com/cloudfoundry/bosh-cli/director/template"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var parser *Parser

	BeforeEach(func() {
		parser = NewParser()
	})

	Describe("NewParser", func() {
		It("returns a parser", func() {
			Expect(parser).ToNot(BeNil())
		})
	})

	Describe("Parse", func() {
		var (
			manifestPath string
			manifest     map[string]interface{}

			executeErr error
		)

		JustBeforeEach(func() {
			tmpfile, err := ioutil.TempFile("", "")
			Expect(err).ToNot(HaveOccurred())
			manifestPath = tmpfile.Name()
			Expect(tmpfile.Close()).ToNot(HaveOccurred())

			WriteManifest(manifestPath, manifest)

			executeErr = parser.Parse(manifestPath)
		})

		AfterEach(func() {
			Expect(os.RemoveAll(manifestPath)).ToNot(HaveOccurred())
		})

		When("given a valid manifest file", func() {
			BeforeEach(func() {
				manifest = map[string]interface{}{
					"applications": []map[string]string{
						{
							"name": "app-1",
						},
						{
							"name": "app-2",
						},
					},
				}
			})

			It("returns nil and sets the applications", func() {
				Expect(executeErr).ToNot(HaveOccurred())
			})
		})

		When("given an invalid manifest file", func() {
			BeforeEach(func() {
				manifest = map[string]interface{}{}
			})

			It("returns an error", func() {
				Expect(executeErr).To(MatchError("must have at least one application"))
			})
		})
	})

	Describe("InterpolateAndParse", func() {
		var (
			pathToManifest   string
			pathsToVarsFiles []string
			vars             []template.VarKV

			executeErr error

			rawManifest []byte
		)

		BeforeEach(func() {
			tempFile, err := ioutil.TempFile("", "manifest-test-")
			Expect(err).ToNot(HaveOccurred())
			Expect(tempFile.Close()).ToNot(HaveOccurred())
			pathToManifest = tempFile.Name()
			vars = nil

			pathsToVarsFiles = nil
		})

		AfterEach(func() {
			Expect(os.RemoveAll(pathToManifest)).ToNot(HaveOccurred())
			for _, path := range pathsToVarsFiles {
				Expect(os.RemoveAll(path)).ToNot(HaveOccurred())
			}
		})

		JustBeforeEach(func() {
			executeErr = parser.InterpolateAndParse(pathToManifest, pathsToVarsFiles, vars)
		})

		Context("regardless of whether the manifest needs interpolation", func() {
			BeforeEach(func() {
				rawManifest = []byte(`---
applications:
- name: spark
  memory: 1G
  instances: 2
- name: flame
  memory: 1G
  instances: 2
`)
				err := ioutil.WriteFile(pathToManifest, rawManifest, 0666)
				Expect(err).ToNot(HaveOccurred())
			})

			It("parses the manifest properly", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(parser.AppNames()).To(ConsistOf("spark", "flame"))
				Expect(parser.PathToManifest).To(Equal(pathToManifest))
				Expect(parser.RawManifest("doesn't, matter")).To(MatchYAML(rawManifest))
			})
		})

		When("the manifest contains variables that need interpolation", func() {
			BeforeEach(func() {
				rawManifest = []byte(`---
applications:
- name: ((var1))
- name: ((var2))
`)
				err := ioutil.WriteFile(pathToManifest, rawManifest, 0666)
				Expect(err).ToNot(HaveOccurred())
			})

			When("only vars files are provided", func() {
				var (
					varsDir string
				)

				BeforeEach(func() {
					var err error
					varsDir, err = ioutil.TempDir("", "vars-test")
					Expect(err).ToNot(HaveOccurred())

					varsFilePath1 := filepath.Join(varsDir, "vars-1")
					err = ioutil.WriteFile(varsFilePath1, []byte("var1: spark"), 0666)
					Expect(err).ToNot(HaveOccurred())

					varsFilePath2 := filepath.Join(varsDir, "vars-2")
					err = ioutil.WriteFile(varsFilePath2, []byte("var2: flame"), 0666)
					Expect(err).ToNot(HaveOccurred())

					pathsToVarsFiles = append(pathsToVarsFiles, varsFilePath1, varsFilePath2)
				})

				AfterEach(func() {
					Expect(os.RemoveAll(varsDir)).ToNot(HaveOccurred())
				})

				When("multiple values for the same variable(s) are provided", func() {
					BeforeEach(func() {
						varsFilePath1 := filepath.Join(varsDir, "vars-1")
						err := ioutil.WriteFile(varsFilePath1, []byte("var1: garbageapp\nvar1: spark\nvar2: doesn't matter"), 0666)
						Expect(err).ToNot(HaveOccurred())

						varsFilePath2 := filepath.Join(varsDir, "vars-2")
						err = ioutil.WriteFile(varsFilePath2, []byte("var2: flame"), 0666)
						Expect(err).ToNot(HaveOccurred())

						pathsToVarsFiles = append(pathsToVarsFiles, varsFilePath1, varsFilePath2)
					})

					It("interpolates the placeholder values", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(parser.AppNames()).To(ConsistOf("spark", "flame"))
					})
				})

				When("the provided files exists and contain valid yaml", func() {
					It("interpolates the placeholder values", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(parser.AppNames()).To(ConsistOf("spark", "flame"))
					})
				})

				When("a variable in the manifest is not provided in the vars file", func() {
					BeforeEach(func() {
						varsFilePath := filepath.Join(varsDir, "vars-1")
						err := ioutil.WriteFile(varsFilePath, []byte("notvar: foo"), 0666)
						Expect(err).ToNot(HaveOccurred())

						pathsToVarsFiles = []string{varsFilePath}
					})

					It("returns an error", func() {
						Expect(executeErr.Error()).To(Equal("Expected to find variables: var1, var2"))
					})
				})

				When("the provided file path does not exist", func() {
					BeforeEach(func() {
						pathsToVarsFiles = []string{"garbagepath"}
					})

					It("returns an error", func() {
						Expect(executeErr).To(HaveOccurred())
						Expect(os.IsNotExist(executeErr)).To(BeTrue())
					})
				})

				When("the provided file is not a valid yaml file", func() {
					BeforeEach(func() {
						varsFilePath := filepath.Join(varsDir, "vars-1")
						err := ioutil.WriteFile(varsFilePath, []byte(": bad"), 0666)
						Expect(err).ToNot(HaveOccurred())

						pathsToVarsFiles = []string{varsFilePath}
					})

					It("returns an error", func() {
						Expect(executeErr).To(HaveOccurred())
						Expect(executeErr).To(MatchError(InvalidYAMLError{
							Err: errors.New("yaml: did not find expected key"),
						}))
					})
				})
			})

			When("only vars are provided", func() {
				BeforeEach(func() {
					vars = []template.VarKV{
						{Name: "var1", Value: "spark"},
						{Name: "var2", Value: "flame"},
					}
				})

				It("interpolates the placeholder values", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(parser.AppNames()).To(ConsistOf("spark", "flame"))
				})
			})

			When("vars and vars files are provided", func() {
				var varsFilePath string
				BeforeEach(func() {
					tmp, err := ioutil.TempFile("", "util-manifest-varsilfe")
					Expect(err).NotTo(HaveOccurred())
					Expect(tmp.Close()).NotTo(HaveOccurred())

					varsFilePath = tmp.Name()
					err = ioutil.WriteFile(varsFilePath, []byte("var1: spark\nvar2: 12345"), 0666)
					Expect(err).ToNot(HaveOccurred())

					pathsToVarsFiles = []string{varsFilePath}
					vars = []template.VarKV{
						{Name: "var2", Value: "flame"},
					}
				})

				AfterEach(func() {
					Expect(os.RemoveAll(varsFilePath)).ToNot(HaveOccurred())
				})

				It("interpolates the placeholder values, prioritizing the vars flag", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(parser.AppNames()).To(ConsistOf("spark", "flame"))
				})
			})
		})
	})

	Describe("AppNames", func() {
		When("given a valid manifest file", func() {
			BeforeEach(func() {
				parser.Applications = []Application{{Name: "app-1"}, {Name: "app-2"}}
			})

			It("gets the app names", func() {
				appNames := parser.AppNames()
				Expect(appNames).To(ConsistOf("app-1", "app-2"))
			})
		})
	})

	Describe("RawManifest", func() {
		When("given an app name", func() {
			When("app is successfully marshalled", func() {
				var manifestPath string

				BeforeEach(func() {
					tmpfile, err := ioutil.TempFile("", "")
					Expect(err).ToNot(HaveOccurred())
					manifestPath = tmpfile.Name()
					Expect(tmpfile.Close()).ToNot(HaveOccurred())

					manifest := map[string]interface{}{
						"applications": []map[string]string{
							{
								"name": "app-1",
							},
							{
								"name": "app-2",
							},
						},
					}
					WriteManifest(manifestPath, manifest)

					executeErr := parser.Parse(manifestPath)
					Expect(executeErr).ToNot(HaveOccurred())
				})

				AfterEach(func() {
					Expect(os.RemoveAll(manifestPath)).ToNot(HaveOccurred())
				})

				It("gets the app's manifest", func() {
					rawManifest, err := parser.RawManifest("app-1")
					Expect(err).ToNot(HaveOccurred())
					Expect(rawManifest).To(MatchYAML(`---
applications:
- name: app-1

- name: app-2
`))

					rawManifest, err = parser.RawManifest("app-2")
					Expect(err).ToNot(HaveOccurred())
					Expect(rawManifest).To(MatchYAML(`---
applications:
- name: app-1

- name: app-2
`))
				})
			})

			PWhen("app marshalling errors", func() {
				It("returns an error", func() {})
			})
		})
	})
})
