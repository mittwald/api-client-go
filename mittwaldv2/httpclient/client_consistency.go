package httpclient

import (
	"github.com/mittwald/api-client-go/pkg/util/httputil"
	"net/http"
)

type eventualConsistencyClient struct {
	inner       RequestRunner
	lastEventId string
}

// NewEventualConsistentClient adds support for the eventual consistency behaviour
// as documented in [1] to a client.
//
// [1]: https://developer.mittwald.de/docs/v2/api/intro/#eventual-consistency
func NewEventualConsistentClient(inner RequestRunner) RequestRunner {
	return &eventualConsistencyClient{inner: inner}
}

func (e *eventualConsistencyClient) addEventReachedToRequest(req *http.Request) {
	req.Header.Set("If-Event-Reached", e.lastEventId)
}

func (e *eventualConsistencyClient) Do(request *http.Request) (*http.Response, error) {
	if e.lastEventId != "" && !httputil.CanRetry(request) {
		e.addEventReachedToRequest(request)
	}

	response, err := e.inner.Do(request)
	if err != nil {
		return response, err
	}

	if e.lastEventId != "" && (response.StatusCode == 403 || response.StatusCode == 404) && httputil.CanRetry(request) {
		e.addEventReachedToRequest(request)

		response, err = e.inner.Do(request)
		if err != nil {
			return response, err
		}
	}

	if etag := response.Header.Get("ETag"); etag != "" {
		e.lastEventId = etag
	}

	return response, err
}
