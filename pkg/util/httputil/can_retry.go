package httputil

import "net/http"

// CanRetry tests if an *http.Request can be safely retried.
//
// This relies on the definition of "idempotent methods" from RFC 7231, section 4.2.2 [1]
//
// [1]: https://www.rfc-editor.org/rfc/rfc7231#section-4.2.2
func CanRetry(req *http.Request) bool {
	return req.Method == http.MethodGet ||
		req.Method == http.MethodHead ||
		req.Method == http.MethodDelete ||
		req.Method == http.MethodPut
}
