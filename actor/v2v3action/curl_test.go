package v2v3action_test

import (
	. "code.cloudfoundry.org/cli/actor/v2v3action"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("curl", func() {
	var (
		requestHeaders, responseHeaders, responseBody string
		actor                                         Actor
		path                                          string
	)

	JustBeforeEach(func() {
		requestHeaders, responseHeaders, responseBody = actor.MakeRequest(path)
	})
	When("the request succeeds", func() {

		It("should make the request", func() {
			Expect(responseHeaders).To(Equal("stuff"))
			Expect(requestHeaders).To(Equal("stuff"))
			Expect(responseBody).To(Equal("stuff"))
		})
	})
})
