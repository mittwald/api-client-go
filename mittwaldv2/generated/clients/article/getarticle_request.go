package article

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type GetArticleRequest struct {
	ArticleID  string
	CustomerID *string
}

func (r *GetArticleRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, r.url(), body)
}

func (r *GetArticleRequest) body() (io.Reader, error) {
	return nil, nil
}

func (r *GetArticleRequest) url() string {
	return fmt.Sprintf("/v2/articles/%s", url.PathEscape(r.ArticleID))
}

func (r *GetArticleRequest) query() url.Values {
	q := make(url.Values)
	if r.CustomerID != nil {
		q.Set("customerId", *r.CustomerID)
	}
	return q
}
