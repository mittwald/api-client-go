package relocation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/httpclient"
	"github.com/mittwald/api-client-go/mittwaldv2/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	CreateLegacyTariffChange(
		ctx context.Context,
		req CreateLegacyTariffChangeRequest,
	) (*CreateLegacyTariffChangeResponse, *http.Response, error)
	CreateRelocation(
		ctx context.Context,
		req CreateRelocationRequest,
	) (*http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

//Register a tariff change for a legacy tariff.
func (c *clientImpl) CreateLegacyTariffChange(
	ctx context.Context,
	req CreateLegacyTariffChangeRequest,
) (*CreateLegacyTariffChangeResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
	if err != nil {
		return nil, nil, err
	}

	httpRes, err := c.client.Do(httpReq.WithContext(ctx))
	if err != nil {
		return nil, httpRes, err
	}

	if httpRes.StatusCode >= 400 {
		err := &httperr.ErrUnexpectedResponse{Response: httpRes}
		return nil, httpRes, err
	}

	var response CreateLegacyTariffChangeResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Relocate an external Project to mittwald.
//
//Give mittwald access to your Provider and let them move your Project to mittwald.
func (c *clientImpl) CreateRelocation(
	ctx context.Context,
	req CreateRelocationRequest,
) (*http.Response, error) {
	httpReq, err := req.BuildRequest()
	if err != nil {
		return nil, err
	}

	httpRes, err := c.client.Do(httpReq.WithContext(ctx))
	if err != nil {
		return httpRes, err
	}

	if httpRes.StatusCode >= 400 {
		err := &httperr.ErrUnexpectedResponse{Response: httpRes}
		return httpRes, err
	}

	return httpRes, nil
}
