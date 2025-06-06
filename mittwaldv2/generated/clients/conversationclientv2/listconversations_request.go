package conversationclientv2

import (
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListConversationsRequest models a request for the
// 'conversation-list-conversations' operation. See [1] for more information.
//
// Get all conversation the authenticated user has created or has access to.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/conversation/conversation-list-conversations
type ListConversationsRequest struct {
	Sort  []ListConversationsRequestQuerySortItem
	Order []ListConversationsRequestQueryOrderItem
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListConversationsRequest) BuildRequest(reqEditors ...func(req *http.Request) error) (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	for _, editor := range reqEditors {
		if err := editor(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

func (r *ListConversationsRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListConversationsRequest) url() string {
	u := url.URL{
		Path:     "/v2/conversations",
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListConversationsRequest) query() url.Values {
	q := make(url.Values)
	for _, val := range r.Sort {
		q.Add("sort", string(val))
	}
	for _, val := range r.Order {
		q.Add("order", string(val))
	}
	return q
}
