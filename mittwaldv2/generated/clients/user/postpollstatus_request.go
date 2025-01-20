package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type PostPollStatusRequest struct {
	Body   PostPollStatusRequestBody
	UserID string
}

func (r *PostPollStatusRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *PostPollStatusRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *PostPollStatusRequest) url() string {
	return fmt.Sprintf("/v2/poll-settings/%s", url.PathEscape(r.UserID))
}

func (r *PostPollStatusRequest) query() url.Values {
	return nil
}