package relocationclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

type Client interface {
	CreateLegacyTariffChange(
		ctx context.Context,
		req CreateLegacyTariffChangeRequest,
		reqEditors ...func(req *http.Request) error,
	) (*CreateLegacyTariffChangeResponse, *http.Response, error)
	CreateRelocation(
		ctx context.Context,
		req CreateRelocationRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// Register a tariff change for a legacy tariff.
func (c *clientImpl) CreateLegacyTariffChange(
	ctx context.Context,
	req CreateLegacyTariffChangeRequest,
	reqEditors ...func(req *http.Request) error,
) (*CreateLegacyTariffChangeResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest(reqEditors...)
	if err != nil {
		return nil, nil, err
	}

	httpRes, err := c.client.Do(httpReq.WithContext(ctx))
	if err != nil {
		return nil, httpRes, err
	}

	if httpRes.StatusCode >= 400 {
		err := httperr.ErrFromResponse(httpRes)
		return nil, httpRes, err
	}

	var response CreateLegacyTariffChangeResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Relocate an external Project to mittwald.
//
// Give mittwald access to your Provider and let them move your Project to mittwald.
func (c *clientImpl) CreateRelocation(
	ctx context.Context,
	req CreateRelocationRequest,
	reqEditors ...func(req *http.Request) error,
) (*http.Response, error) {
	httpReq, err := req.BuildRequest(reqEditors...)
	if err != nil {
		return nil, err
	}

	httpRes, err := c.client.Do(httpReq.WithContext(ctx))
	if err != nil {
		return httpRes, err
	}

	if httpRes.StatusCode >= 400 {
		err := httperr.ErrFromResponse(httpRes)
		return httpRes, err
	}

	return httpRes, nil
}
