package appclientv2

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// ListAppinstallationsForUserRequest models a request for the
// 'app-list-appinstallations-for-user' operation. See [1] for more information.
//
// List AppInstallations that a user has access to.
//
// [1]:
// https://developer.mittwald.de/docs/v2/reference/app/app-list-appinstallations-for-user
type ListAppinstallationsForUserRequest struct {
	AppIDs      []string
	SearchTerm  *string
	PhpVersions []string
	Limit       *int64
	Skip        *int64
	Page        *int64
}

// BuildRequest builds an *http.Request instance from this request that may be used
// with any regular *http.Client instance.
func (r *ListAppinstallationsForUserRequest) BuildRequest() (*http.Request, error) {
	body, contentType, err := r.body()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, r.url(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

func (r *ListAppinstallationsForUserRequest) body() (io.Reader, string, error) {
	return nil, "", nil
}

func (r *ListAppinstallationsForUserRequest) url() string {
	u := url.URL{
		Path:     "/v2/app-installations",
		RawQuery: r.query().Encode(),
	}
	return u.String()
}

func (r *ListAppinstallationsForUserRequest) query() url.Values {
	q := make(url.Values)
	for _, val := range r.AppIDs {
		q.Add("appIds", val)
	}
	if r.SearchTerm != nil {
		q.Set("searchTerm", *r.SearchTerm)
	}
	for _, val := range r.PhpVersions {
		q.Add("phpVersions", val)
	}
	if r.Limit != nil {
		q.Set("limit", fmt.Sprintf("%d", *r.Limit))
	}
	if r.Skip != nil {
		q.Set("skip", fmt.Sprintf("%d", *r.Skip))
	}
	if r.Page != nil {
		q.Set("page", fmt.Sprintf("%d", *r.Page))
	}
	return q
}
