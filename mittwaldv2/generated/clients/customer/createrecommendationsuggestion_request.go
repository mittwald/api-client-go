package customer

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

type CreateRecommendationSuggestionRequest struct {
	Body       CreateRecommendationSuggestionRequestBody
	CustomerID string
}

func (r *CreateRecommendationSuggestionRequest) BuildRequest() (*http.Request, error) {
	body, err := r.body()
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, r.url(), body)
}

func (r *CreateRecommendationSuggestionRequest) body() (io.Reader, error) {
	out, err := json.Marshal(&r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling JSON: %w", err)
	}
	return bytes.NewReader(out), nil
}

func (r *CreateRecommendationSuggestionRequest) url() string {
	return fmt.Sprintf("/v2/customers/%s/recommendation-suggestions", url.PathEscape(r.CustomerID))
}

func (r *CreateRecommendationSuggestionRequest) query() url.Values {
	return nil
}
