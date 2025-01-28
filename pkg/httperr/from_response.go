package httperr

import "net/http"

// ErrFromResponse maps an HTTP response (with an error status code) to a (more
// or less) specific error type.
func ErrFromResponse(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusNotFound:
		return &ErrNotFound{Response: res}
	case http.StatusForbidden:
		return &ErrPermissionDenied{Response: res}
	default:
		return &ErrUnexpectedResponse{Response: res}
	}
}
