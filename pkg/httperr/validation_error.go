package httperr

import (
	"encoding/json"
	"fmt"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/commonsv2"
	"net/http"
)

func IsValidationErrorResponse(res *http.Response) (*commonsv2.ValidationErrors, bool) {
	target := commonsv2.ValidationErrors{}

	if err := json.NewDecoder(res.Body).Decode(&target); err != nil {
		return nil, false
	}

	if target.Type != commonsv2.ValidationErrorsTypeValidationError {
		return nil, false
	}

	return &target, true
}

// ErrValidation represents a validation error response from the server.
type ErrValidation struct {
	Response        *http.Response
	ValidationError *commonsv2.ValidationErrors
}

func (e *ErrValidation) Error() string {
	if e.ValidationError != nil && e.ValidationError.Message != nil {
		return fmt.Sprintf("invalid request: %s", *e.ValidationError.Message)
	}
	return "invalid request: <nil>"
}
