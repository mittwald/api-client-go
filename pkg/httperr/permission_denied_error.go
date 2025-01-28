package httperr

import (
	"fmt"
	"net/http"
)

// ErrPermissionDenied represents the error of trying to access a resource that
// the currently authenticated user is not authorized to access.
type ErrPermissionDenied struct {
	Response *http.Response
}

func (e *ErrPermissionDenied) Error() string {
	return fmt.Sprintf("permission to resource denied: %s", e.Response.Request.URL)
}
