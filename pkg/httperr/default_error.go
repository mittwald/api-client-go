package httperr

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/commonsv2"
)

func IsDefaultErrorResponse(res *http.Response) (*commonsv2.Error, bool) {
	target := commonsv2.Error{}

	if err := json.NewDecoder(res.Body).Decode(&target); err != nil {
		return nil, false
	}

	return &target, true
}

// ErrDefault represents the default fallback error used in the mittwald api.
type ErrDefault struct {
	Response     *http.Response
	DefaultError *commonsv2.Error
}

func (e *ErrDefault) Error() string {
	if e.DefaultError != nil {
		return fmt.Sprintf("%s: %s", e.DefaultError.Type, e.DefaultError.Message)
	}
	return "unexpected default error: <nil>"
}
