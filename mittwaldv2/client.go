package mittwaldv2

import (
	"context"
	"net/http"

	generatedv2 "github.com/mittwald/api-client-go/mittwaldv2/generated/clients"
	"github.com/mittwald/api-client-go/mittwaldv2/httpclient"
)

var defaultOpts = [...]ClientOption{
	WithBaseURL("https://api.mittwald.de/v2"),
}

// New creates a new mittwaldv2 API client.
//
// Use the "opts" parameter to configure things like authentication and other
// specific client-side behaviours.
//
// Please note that this function may fail with an error, because some authentication mechanisms actively communicate with the API
func New(ctx context.Context, opts ...ClientOption) (generatedv2.Client, error) {
	var runner httpclient.RequestRunner = http.DefaultClient
	var err error

	allOpts := append(
		defaultOpts[:],
		opts...,
	)

	for _, o := range allOpts {
		runner, err = o(ctx, runner)
		if err != nil {
			return nil, err
		}
	}

	return generatedv2.NewClient(runner), nil
}
