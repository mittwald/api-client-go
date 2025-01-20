package mittwaldv2

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	generatedv2 "github.com/mittwald/api-client-go/mittwaldv2/generated/clients"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/marketplace"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/user"
	"github.com/mittwald/api-client-go/mittwaldv2/httpclient"
	"os"
)

const apiTokenEnvVar = "MITTWALD_API_TOKEN"

// WithAccessToken adds a static access token to all executed requests.
// The access token needs to be obtained ahead-of-time.
func WithAccessToken(token string) ClientOption {
	return func(_ context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		return httpclient.NewAuthenticatedClient(inner, token), nil
	}
}

// WithAccessTokenFromEnv is a convenience wrapper around WithAccessToken that
// automatically retrieves the API token from the process environment variables.
func WithAccessTokenFromEnv() ClientOption {
	return func(_ context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		token, ok := os.LookupEnv(apiTokenEnvVar)
		if !ok {
			return nil, fmt.Errorf("environment variable '%s' must be set", apiTokenEnvVar)
		}

		return httpclient.NewAuthenticatedClient(inner, token), nil
	}
}

// WithUsernamePassword authenticates the API user using their actual email
// address+password combination.
//
// NOTE: This will only work then MFA is *not* activated for a user. In general,
// it is recommended to use API tokens, instead of username+password.
func WithUsernamePassword(email, password string) ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		req := user.AuthenticateRequest{
			Body: user.AuthenticateRequestBody{
				Email:    email,
				Password: password,
			},
		}

		resp, _, err := generatedv2.NewClient(inner).User().Authenticate(ctx, req)
		if err != nil {
			return nil, fmt.Errorf("error while authenticating user '%s': %w", email, err)
		}

		if r := resp.AlternativeAuthenticateOKResponse; r != nil {
			return WithAccessToken(r.Token)(ctx, inner)
		}

		return nil, fmt.Errorf("second factor required; use an API token instead")
	}
}

// WithAccessTokenRetrievalKey authenticates a user using an "access token retrieval key".
//
// This is a special mechanism for one-click-authenticating users that are
// redirected to an mStudio extension [1].
//
// [1]: https://developer.mittwald.de/docs/v2/contribution/overview/concepts/authentication/
func WithAccessTokenRetrievalKey(userID uuid.UUID, accessTokenRetrievalKey string) ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		req := user.AuthenticateWithAccessTokenRetrievalKeyRequest{
			Body: user.AuthenticateWithAccessTokenRetrievalKeyRequestBody{
				AccessTokenRetrievalKey: accessTokenRetrievalKey,
				UserId:                  userID,
			},
		}

		resp, _, err := generatedv2.NewClient(inner).User().AuthenticateWithAccessTokenRetrievalKey(ctx, req)
		if err != nil {
			return nil, fmt.Errorf("error while authenticating user '%s': %w", userID, err)
		}

		return WithAccessToken(resp.Token)(ctx, inner)
	}
}

// WithExtensionSecret authenticates an mStudio extension (i.e. not a user).
func WithExtensionSecret(extensionInstanceID uuid.UUID, extensionSecret string) ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		req := marketplace.AuthenticateInstanceRequest{
			ExtensionInstanceID: extensionInstanceID,
			Body: marketplace.AuthenticateInstanceRequestBody{
				ExtensionInstanceSecret: extensionSecret,
			},
		}

		resp, _, err := generatedv2.NewClient(inner).Marketplace().AuthenticateInstance(ctx, req)
		if err != nil {
			return nil, fmt.Errorf("error while authenticating extension instance: %w", err)
		}

		return WithAccessToken(resp.PublicToken)(ctx, inner)
	}
}
