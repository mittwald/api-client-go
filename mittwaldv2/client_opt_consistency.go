package mittwaldv2

import (
	"context"
	"github.com/mittwald/api-client-go/mittwaldv2/httpclient"
)

// WithEventualConsistency enables client-side support for the eventual-consistency
// behaviour as documented in [1].
//
// [1]: https://developer.mittwald.de/docs/v2/api/intro/#eventual-consistency
func WithEventualConsistency() ClientOption {
	return func(_ context.Context, inner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		return httpclient.NewEventualConsistentClient(inner), nil
	}
}
