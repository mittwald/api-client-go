package mittwaldv2

import (
	"context"
	"log/slog"

	"github.com/mittwald/api-client-go/pkg/httpclient"
)

// WithRequestLogging adds a logging middleware to the request runner chain
// allowing you to log all executed HTTP requests in a slog.Logger of your
// choice.
//
// Be mindful of the log{Request,Response}Bodies parameters; these will cause
// the logger to print the full request bodies without redaction, which may
// easily leak sensitive data.
func WithRequestLogging(logger *slog.Logger, logRequestBodies, logResponseBodies bool) ClientOption {
	return func(ctx context.Context, runner httpclient.RequestRunner) (httpclient.RequestRunner, error) {
		return httpclient.NewLoggingClient(runner, logger, logRequestBodies, logResponseBodies), nil
	}
}
