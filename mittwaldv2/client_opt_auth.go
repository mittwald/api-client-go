package mittwaldv2

import (
	"context"
	"fmt"
	"os"
	"time"

	generatedv2 "github.com/mittwald/api-client-go/mittwaldv2/generated/clients"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/marketplaceclientv2"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/userclientv2"
	"github.com/mittwald/api-client-go/pkg/httpclient"
)

const apiTokenEnvVar = "MITTWALD_API_TOKEN"
const apiTokenHeader = "X-Access-Token"

// WithAccessToken adds a static access token to all executed requests.
// The access token needs to be obtained ahead-of-time.
func WithAccessToken(token string) ClientOption {
	return func(_ context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		return httpclient.NewAuthenticatedClient(inner, token, apiTokenHeader), nil
	}
}

// WithAccessTokenFromEnv is a convenience wrapper around WithAccessToken that
// automatically retrieves the API token from the process environment variables.
func WithAccessTokenFromEnv() ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		token, ok := os.LookupEnv(apiTokenEnvVar)
		if !ok {
			return nil, fmt.Errorf("environment variable '%s' must be set", apiTokenEnvVar)
		}

		return WithAccessToken(token)(ctx, inner)
	}
}

// WithAccessTokenAndRefresh supports automatically refreshing API tokens.
//
// Clients with this option will refresh API tokens automatically when they expire.
func WithAccessTokenAndRefresh(token, refreshToken string, expiration time.Time) ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		refreshToken := refreshToken
		refreshFunc := func() (string, time.Time, error) {
			req := userclientv2.RefreshSessionRequest{
				Body: userclientv2.RefreshSessionRequestBody{
					RefreshToken: refreshToken,
				},
			}

			resp, _, err := generatedv2.NewClient(inner).User().RefreshSession(ctx, req)
			if err != nil {
				return "", time.Time{}, err
			}

			refreshToken = resp.RefreshToken

			return resp.Token, resp.ExpiresAt, nil
		}

		return httpclient.NewAuthenticatedClientWithRefresh(inner, token, apiTokenHeader, expiration, refreshFunc), nil
	}
}

// WithAccessTokenAndRefreshFunc supports automatically refreshing API tokens.
//
// Clients with this option will refresh API tokens automatically when they expire.
func WithAccessTokenAndRefreshFunc(token string, expiration time.Time, refreshFunc httpclient.RefreshFunc) ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		return httpclient.NewAuthenticatedClientWithRefresh(inner, token, apiTokenHeader, expiration, refreshFunc), nil
	}
}

// WithUsernamePassword authenticates the API user using their actual email
// address+password combination.
//
// NOTE: This will only work then MFA is *not* activated for a user. In general,
// it is recommended to use API tokens, instead of username+password.
func WithUsernamePassword(email, password string) ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		req := userclientv2.AuthenticateRequest{
			Body: userclientv2.AuthenticateRequestBody{
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
// Clients authenticated with this method will refresh their API tokens
// automatically.
//
// [1]: https://developer.mittwald.de/docs/v2/contribution/overview/concepts/authentication/
func WithAccessTokenRetrievalKey(userID string, accessTokenRetrievalKey string) ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		req := userclientv2.AuthenticateWithAccessTokenRetrievalKeyRequest{
			Body: userclientv2.AuthenticateWithAccessTokenRetrievalKeyRequestBody{
				AccessTokenRetrievalKey: accessTokenRetrievalKey,
				UserId:                  userID,
			},
		}

		resp, _, err := generatedv2.NewClient(inner).User().AuthenticateWithAccessTokenRetrievalKey(ctx, req)
		if err != nil {
			return nil, fmt.Errorf("error while authenticating user '%s': %w", userID, err)
		}

		return WithAccessTokenAndRefresh(resp.Token, resp.RefreshToken, resp.ExpiresAt)(ctx, inner)
	}
}

// WithExtensionSecret authenticates an mStudio extension (i.e. not a user).
//
// Clients authenticated with this method will refresh their API tokens
// automatically.
func WithExtensionSecret(extensionInstanceID string, extensionSecret string) ClientOption {
	return func(ctx context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		refreshFunc := func() (string, time.Time, error) {
			req := marketplaceclientv2.AuthenticateInstanceRequest{
				ExtensionInstanceID: extensionInstanceID,
				Body: marketplaceclientv2.AuthenticateInstanceRequestBody{
					ExtensionInstanceSecret: extensionSecret,
				},
			}

			resp, _, err := generatedv2.NewClient(inner).Marketplace().AuthenticateInstance(ctx, req)
			if err != nil {
				return "", time.Time{}, fmt.Errorf("error while authenticating extension instance: %w", err)
			}

			return resp.PublicToken, resp.Expiry, nil
		}

		token, expiration, err := refreshFunc()
		if err != nil {
			return nil, err
		}

		return WithAccessTokenAndRefreshFunc(token, expiration, refreshFunc)(ctx, inner)
	}
}
