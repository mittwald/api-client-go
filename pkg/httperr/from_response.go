package httperr

import (
	"net/http"
)

// ErrFromResponse maps an HTTP response (with an error status code) to a (more
// or less) specific error type.
func ErrFromResponse(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusNotFound:
		return &ErrNotFound{Response: res}
	case http.StatusForbidden:
		return &ErrPermissionDenied{Response: res}
	case http.StatusBadRequest:
		if validationError, isValidationError := IsValidationErrorResponse(res); isValidationError {
			return &ErrValidation{Response: res, ValidationError: validationError}
		}
		return &ErrBadRequest{Response: res}
	default:
		if defaultError, isDefaultError := IsDefaultErrorResponse(res); isDefaultError {
			return &ErrDefault{Response: res, DefaultError: defaultError}
		}

		return &ErrUnexpectedResponse{Response: res}
	}
}
