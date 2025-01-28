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
	return fmt.Sprintf("resource not found: %s", e.Response.Request.URL)
}
