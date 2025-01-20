package domain

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type SuggestRequest struct {
	Prompt      string
	DomainCount *int64
	Tlds        []string
}

func (r *SuggestRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *SuggestRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *SuggestRequest) url() string {
	return "/v2/domain-suggestions"
}

func (r *SuggestRequest) query() url.Values {
	q := make(url.Values)
	q.Set("prompt", r.Prompt)
	if r.DomainCount != nil {
		q.Set("domainCount", fmt.Sprintf("%d", *r.DomainCount))
	}
	for _, val := range r.Tlds {
		q.Add("tlds", val)
	}
	return q
}
