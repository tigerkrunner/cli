package ccv3_test

import (
	"net/http"

	. "code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Curl", func() {
	var (
		client          *Client
		rootRespondWith http.HandlerFunc
		v3RespondWith   http.HandlerFunc
		executeErr      error
	)

	BeforeEach(func() {
		rootRespondWith = nil
		v3RespondWith = nil
	})

	JustBeforeEach(func() {
		client, _ = NewTestClient()

		server.AppendHandlers(
			CombineHandlers(
				VerifyRequest(http.MethodGet, "/"),
				rootRespondWith,
			),
			CombineHandlers(
				VerifyRequest(http.MethodGet, "/v3"),
				v3RespondWith,
			))

		_, _, _, executeErr = client.GetInfo()
	})

	Describe("MakeRawRequest", func() {
		It("WIP", func() {
			Expect(executeErr).To(Equal(nil))
			Expect(true).To(Equal(false))
		})
	})
})
