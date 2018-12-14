package v6_test

import (
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
		executeErr error
		extraArgs  []string
		outBuffer  *Buffer
	)

	BeforeEach(func() {
		outBuffer = NewBuffer()
		testUI = ui.NewTestUI(NewBuffer(), outBuffer, NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeActor = new(v6fakes.FakeCurlActor)
		extraArgs = nil

		cmd = CurlCommand{
			Config: fakeConfig,
			UI:     testUI,
			Actor:  fakeActor,
		}
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

			When("too many positional args are passed", func() {
				BeforeEach(func() {
					extraArgs = []string{"foo"}
				})

				It("returns an error and displays", func() {
					Expect(executeErr).To(MatchError(translatableerror.TooManyArgumentsError{ExtraArgument: extraArgs[0]}))
				})
			})

			When("The APIPath is valid", func() {
				var expectedJSONResponse, expectedRequestHeaders, expectedResponseHeaders string

				BeforeEach(func() {
					expectedJSONResponse = `{
					"key1": "value1",
					"key2": "value2"
			}`
					expectedRequestHeaders = "Request: test\n X-Foo: foo"
					expectedResponseHeaders = "Response: test\n X-Bar: bar"

					fakeActor.MakeRequestReturns(expectedRequestHeaders, expectedResponseHeaders, expectedJSONResponse)
				})

				It("should make the request and display the JSON response", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(outBuffer.Contents()).To(MatchJSON(expectedJSONResponse))
				})

				XWhen("the -v flag is not set", func() {
					BeforeEach(func() {
						fakeConfig.VerboseReturns(false, nil)
					})

					It("makes a request and displays the JSON response", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(testUI.Out).ToNot(Say(expectedRequestHeaders))
						Expect(testUI.Out).ToNot(Say(expectedResponseHeaders))
						Expect(testUI.Out).To(Say(expectedJSONResponse))
					})
				})

				XWhen("-v flag is set", func() {
					BeforeEach(func() {
						fakeConfig.VerboseReturns(true, nil)
					})
					It("displays the request and response headers", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(testUI.Out).To(Say(expectedRequestHeaders))
						Expect(testUI.Out).To(Say(expectedResponseHeaders))
						Expect(testUI.Out).To(Say(expectedJSONResponse))
					})
				})
			})
		})
	})

})
