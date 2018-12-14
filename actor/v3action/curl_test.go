package v3action_test

import (
	. "code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/actor/v3action/v3actionfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Curl Actor", func() {

	var (
		actor                     *Actor
		fakeCloudControllerClient *v3actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v3actionfakes.FakeCloudControllerClient)
		actor = NewActor(fakeCloudControllerClient, nil, nil, nil)
	})

	Describe("MakeRequest", func() {
		When("path is valid", func() {
			It("makes the request and returns the headers and response JSON", func() {
				actor.MakeRequest("")
				Expect(true).To(Equal(false))
			})
		})
	})
})
