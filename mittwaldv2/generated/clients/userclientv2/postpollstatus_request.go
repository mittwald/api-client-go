package userclientv2

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

// PostPollStatusRequest models a request for the 'user-post-poll-status'
// operation. See [1] for more information.
//
// Store new or update poll settings.
//
// [1]: https://developer.mittwald.de/docs/v2/reference/user/user-post-poll-status
type PostPollStatusRequest struct {
	Body   PostPollStatusRequestBody
	UserID string
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *PostPollStatusRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *PostPollStatusRequest) body() (io.Reader, string, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), "application/json", nil
}

func (r *PostPollStatusRequest) url() string {
	return fmt.Sprintf("/v2/poll-settings/%s", url.PathEscape(r.UserID))
}

func (r *PostPollStatusRequest) query() url.Values {
	return nil
}
