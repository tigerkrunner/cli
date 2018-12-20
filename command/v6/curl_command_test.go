package v6_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/translatableerror"
	. "code.cloudfoundry.org/cli/command/v6"
	"code.cloudfoundry.org/cli/command/v6/v6fakes"
	"code.cloudfoundry.org/cli/util/ui"
)

var _ = FDescribe("CurlCommand", func() {
	var (
		cmd        CurlCommand
		testUI     *ui.UI
		fakeConfig *commandfakes.FakeConfig
		fakeActor  *v6fakes.FakeCurlActor
		extraArgs  []string

		executeErr error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeActor = new(v6fakes.FakeCurlActor)
		extraArgs = nil

		cmd = CurlCommand{
			Config: fakeConfig,
			Actor:  fakeActor,
			UI:     testUI,
		}
		cmd.RequiredArgs.Path = "/some/api/path"

		fakeConfig.TargetReturns("https://my.fake.api")
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(extraArgs)
	})

	When("the refactor is incomplete", func() {
		When("CF_CLI_CURL_EXPERIMENTAL is false", func() {
			BeforeEach(func() {
				fakeConfig.ExperimentalReturns(false)
			})

			It("returns an UnrefactoredCommandError", func() {
				Expect(executeErr).To(MatchError(translatableerror.UnrefactoredCommandError{}))
			})
		})

		When("CF_CLI_CURL_EXPERIMENTAL is true", func() {
			BeforeEach(func() {
				fakeConfig.CurlExperimentalReturns(true)
			})

			When("the server returns a response", func() {
				BeforeEach(func() {
					buf := bytes.NewBufferString("my fancy response")
					body := ioutil.NopCloser(buf)

					h := http.Header{}
					h.Add("Lol", "WUT")

					resp := &http.Response{
						Body:       body,
						Header:     h,
						Status:     "200 OK",
						StatusCode: 200,
						Proto:      "1.1",
					}
					fakeActor.DoReturns(resp, nil)
				})

				// TODO: refactor this so that it makes sense for the When, or move it out.
				It("makes a GET request to the provided endpoint on the targeted API", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(fakeActor.DoCallCount()).To(Equal(1))

					req := fakeActor.DoArgsForCall(0)
					Expect(req.Method).To(Equal(http.MethodGet))
					Expect(req.URL.Path).To(Equal("/some/api/path"))
					Expect(req.URL.Host).To(Equal("my.fake.api"))
				})

				It("prints the response body", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(testUI.Out).To(Say("my fancy response"))
				})

				When("verbose logging is turned on", func() {
					BeforeEach(func() {
						fakeConfig.VerboseReturns(true, nil)
					})

					It("prints the request", func() {
						Expect(testUI.Out).To(Say(`REQUEST:\s+\[\d{4}-\d{1,2}-\d{1,2}T.*\]`))
						Expect(testUI.Out).To(Say("GET /some/api/path HTTP/1.1"))
						Expect(testUI.Out).To(Say("Host: my.fake.api"))
					})

					It("prints the response headers", func() {
						Expect(testUI.Out).To(Say(`RESPONSE:\s+\[\d{4}-\d{1,2}-\d{1,2}T.*\]`))
						Expect(testUI.Out).To(Say("HTTP/1.1 200 OK"))
						Expect(testUI.Out).To(Say("LOL: WUT"))
					})
				})
			})

			When("making the request fails", func() {
				BeforeEach(func() {
					fakeActor.DoReturns(nil, errors.New("whoops"))
				})

				It("returns the error", func() {
					Expect(executeErr).To(MatchError("whoops"))
				})
			})

			When("too many positional args are passed", func() {
				BeforeEach(func() {
					extraArgs = []string{"foo"}
				})

				It("returns an error", func() {
					Expect(executeErr).To(MatchError(translatableerror.TooManyArgumentsError{ExtraArgument: extraArgs[0]}))
				})
			})
		})
	})
})
