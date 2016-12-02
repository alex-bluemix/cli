package uaa_test

import (
	"net/http"

	. "code.cloudfoundry.org/cli/api/uaa"
	"code.cloudfoundry.org/cli/api/uaa/uaafakes"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("User", func() {
	var (
		client    *Client
		fakeStore *uaafakes.FakeAuthenticationStore
	)

	BeforeEach(func() {
		client, fakeStore = NewTestUAAClientAndStore()
	})

	PDescribe("NewUser", func() {
		BeforeEach(func() {
			response := `{
			}`
			server.AppendHandlers(
				CombineHandlers(
					VerifyRequest(http.MethodPost, "/Users"),
					VerifyHeaderKV("Accept", "application/json"),
					VerifyHeaderKV("Content-Type", "application/json"),
					VerifyBody([]byte("")),
					RespondWith(http.StatusOK, response),
				))
		})

		It("refreshes the token", func() {
		})
	})
})
