package mittwaldv2

import (
	"context"
	"github.com/mittwald/api-client-go/mittwaldv2/httpclient"
)

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

func WithBaseURL(baseURL string) ClientOption {
	return func(_ context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		return httpclient.NewClientWithBaseURL(inner, baseURL)
	}
}
