package mittwaldv2

import (
	"context"
	"net/http"

	generatedv2 "github.com/mittwald/api-client-go/mittwaldv2/generated/clients"
	"github.com/mittwald/api-client-go/pkg/httpclient"
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

	// caution, this is counter-intuitive: since the options are basically a chain of wrappers around the actual
	// request runner, they run inside-out and the options added _last_ will be applied _first_. This is why the default
	// options need to come last.
	allOpts := append(
		opts,
		defaultOpts[:]...,
	)

	for _, o := range allOpts {
		runner, err = o(ctx, runner)
		if err != nil {
			return nil, err
		}
	}

	return generatedv2.NewClient(runner), nil
}
