package httperr

import (
	"fmt"
	"net/http"
)

// ErrUnexpectedResponse represents the situation in which an unexpected HTTP
// status code was returned from an API call.
type ErrUnexpectedResponse struct {
	Response *http.Response
}

func (e *ErrUnexpectedResponse) Error() string {
	if e == nil || e.Response == nil {
		return "unexpected http response: <nil>"
	}
	return fmt.Sprintf("unexpected http response %d", e.Response.StatusCode)
}
