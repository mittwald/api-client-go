package httpclient

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
)

type loggingClient struct {
	inner             RequestRunner
	logger            *slog.Logger
	logRequestBodies  bool
	logResponseBodies bool
}

func NewLoggingClient(inner RequestRunner, logger *slog.Logger, logRequestBodies, logResponseBodies bool) RequestRunner {
	return &loggingClient{
		inner:             inner,
		logger:            logger,
		logRequestBodies:  logRequestBodies,
		logResponseBodies: logResponseBodies,
	}
}

func (c *loggingClient) Do(request *http.Request) (*http.Response, error) {
	l := c.logger.With("req.method", request.Method, "req.url", request.URL.String())

	if c.logRequestBodies && request.Body != nil {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}

		request.Body = io.NopCloser(bytes.NewBuffer(body))
		l = l.With("req.body", string(body))
	}

	l.Debug("executing request")

	response, err := c.inner.Do(request)

	if response != nil {
		l = l.With("res.status", response.StatusCode)
		if c.logResponseBodies && response.Body != nil {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}

			response.Body = io.NopCloser(bytes.NewBuffer(body))
			l = l.With("res.body", string(body))
		}
	}

	l.Debug("received response")

	return response, err
}
