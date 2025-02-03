package mittwaldv2_test

import (
	"context"
	"net/http"
	"os"

	"github.com/mittwald/api-client-go/mittwaldv2"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/userclientv2"
	"github.com/mittwald/api-client-go/pkg/httpclient_mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client authentication", func() {
	Describe("WithAccessToken", func() {
		It("should append the provided access token to all requests", func() {
			ctx := context.Background()

			runner := &httpclient_mock.MockRequestRunner{}
			runner.ExpectRequest(http.MethodGet, "/v2/users/self", httpclient_mock.WithJSONResponse(map[string]any{}))

			client, err := mittwaldv2.New(ctx, mittwaldv2.WithHTTPClient(runner), mittwaldv2.WithAccessToken("FOOBAR"))

			Expect(err).NotTo(HaveOccurred())

			_, _, err = client.User().GetUser(ctx, userclientv2.GetUserRequest{UserID: "self"})

			Expect(err).NotTo(HaveOccurred())
			Expect(runner.Requests).To(HaveLen(1))
			Expect(runner.Requests[0].Header.Get("X-Access-Token")).To(Equal("FOOBAR"))
		})
	})

	Describe("WithAccessTokenFromEnv", func() {
		It("should retrieve the access token from the environment", func() {
			Expect(os.Setenv("MITTWALD_API_TOKEN", "FOOBAR")).To(Succeed())

			ctx := context.Background()

			runner := &httpclient_mock.MockRequestRunner{}
			runner.ExpectRequest(http.MethodGet, "/v2/users/self", httpclient_mock.WithJSONResponse(map[string]any{}))

			client, err := mittwaldv2.New(ctx, mittwaldv2.WithHTTPClient(runner), mittwaldv2.WithAccessTokenFromEnv())

			Expect(err).NotTo(HaveOccurred())

			_, _, err = client.User().GetUser(ctx, userclientv2.GetUserRequest{UserID: "self"})

			Expect(err).NotTo(HaveOccurred())
			Expect(runner.Requests).To(HaveLen(1))
			Expect(runner.Requests[0].Header.Get("X-Access-Token")).To(Equal("FOOBAR"))
		})
	})

	Describe("WithUsernamePassword", func() {
		It("should retrieve the access token from an actual login", func() {
			ctx := context.Background()

			runner := &httpclient_mock.MockRequestRunner{}
			runner.ExpectRequest(http.MethodPost, "/v2/authenticate", httpclient_mock.WithJSONResponse(map[string]any{"token": "FOOBAR"}))
			runner.ExpectRequest(http.MethodGet, "/v2/users/self", httpclient_mock.WithJSONResponse(map[string]any{}))

			client, err := mittwaldv2.New(ctx, mittwaldv2.WithHTTPClient(runner), mittwaldv2.WithUsernamePassword("martin@foo.example", "secret"))

			Expect(err).NotTo(HaveOccurred())

			_, _, err = client.User().GetUser(ctx, userclientv2.GetUserRequest{UserID: "self"})

			Expect(err).NotTo(HaveOccurred())
			Expect(runner.Requests).To(HaveLen(2))
			Expect(runner.Requests[0].Header.Get("X-Access-Token")).To(Equal(""))
			Expect(runner.Requests[1].Header.Get("X-Access-Token")).To(Equal("FOOBAR"))
		})

		It("should return an error when 2FA is required", func() {
			ctx := context.Background()

			runner := &httpclient_mock.MockRequestRunner{}
			runner.ExpectRequest(
				http.MethodPost,
				"/v2/authenticate",
				httpclient_mock.WithStatus(http.StatusAccepted),
				httpclient_mock.WithJSONResponse(map[string]any{"name": "SecondFactorRequired"}))

			_, err := mittwaldv2.New(ctx, mittwaldv2.WithHTTPClient(runner), mittwaldv2.WithUsernamePassword("martin@foo.example", "secret"))

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("second factor required; use an API token instead"))
		})
	})
})
