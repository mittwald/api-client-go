package httperr

import (
	"fmt"
	"net/http"
)

// ErrNotFound represents the error of trying to access a resource that does
// not exist.
type ErrNotFound struct {
	Response *http.Response
}

func (e *ErrNotFound) Error() string {
	if e == nil || e.Response == nil {
		return "resource not found: <nil>"
	}
	return fmt.Sprintf("resource not found: %s", e.Response.Request.URL)
}
