package httperr

import (
	"fmt"
	"net/http"
)

// ErrBadRequest represents the error of an invalid request to the server.
type ErrBadRequest struct {
	Response *http.Response
}

func (e *ErrBadRequest) Error() string {
	if e == nil || e.Response == nil {
		return "bad request: <nil>"
	}
	return fmt.Sprintf("bad request: %s", e.Response.Request.URL)
}
