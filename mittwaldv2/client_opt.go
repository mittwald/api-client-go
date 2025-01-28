package mittwaldv2

import (
	"context"
	"github.com/mittwald/api-client-go/pkg/httpclient"
)

// ClientOption defines a generic data type for modifying the behaviour of the
// mittwald v2 API client.
//
// Client behaviours are implemented as a chain of middlewares that are wrapped
// around a HTTP client implementation (or more precisely, any implementation of
// the httpclient.RequestRunner interface, which is also implemented by the
// default *http.Client type).
type ClientOption func(ctx context.Context, runner httpclient.RequestRunner) (httpclient.RequestRunner, error)

// WithHTTPClient allows you to override the base HTTP client to use for all
// requests.
//
// In the simplest case, this may be an *http.Client, which implements the
// httpclient.RequestRunner interface. However, the interface also allows you to
// supply your own implementation.
//
// NOTE: When used with other ClientOption's, this one must be the first.
func WithHTTPClient(client httpclient.RequestRunner) ClientOption {
	return func(_ context.Context, _ httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		return client, nil
	}
}

// WithBaseURL allows you to override the base URL that is used for HTTP
// requests.
//
// If omitted, this will default to "https://api.mittwald.de/v2" as base URL.
// During regular usage, there should be no need to change the base URL, but it
// may be useful during testing.
func WithBaseURL(baseURL string) ClientOption {
	return func(_ context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		return httpclient.NewClientWithBaseURL(inner, baseURL)
	}
}
