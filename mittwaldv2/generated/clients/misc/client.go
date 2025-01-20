package misc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	ServicetokenAuthenticateService(
		ctx context.Context,
		req ServicetokenAuthenticateServiceRequest,
	) (*ServicetokenAuthenticateServiceResponse, *http.Response, error)
	VerificationVerifyAddress(
		ctx context.Context,
		req VerificationVerifyAddressRequest,
	) (*VerificationVerifyAddressResponse, *http.Response, error)
	VerificationVerifyCompany(
		ctx context.Context,
		req VerificationVerifyCompanyRequest,
	) (*VerificationVerifyCompanyResponse, *http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// Obtain a service token.
func (c *clientImpl) ServicetokenAuthenticateService(
	ctx context.Context,
	req ServicetokenAuthenticateServiceRequest,
) (*ServicetokenAuthenticateServiceResponse, *http.Response, error) {
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

	var response ServicetokenAuthenticateServiceResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Check if an address exists.
//
// Only the DACH region is currently supported.
func (c *clientImpl) VerificationVerifyAddress(
	ctx context.Context,
	req VerificationVerifyAddressRequest,
) (*VerificationVerifyAddressResponse, *http.Response, error) {
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

	var response VerificationVerifyAddressResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Check if a company exists.
//
// Only companies registered in the german company register are currently supported.
func (c *clientImpl) VerificationVerifyCompany(
	ctx context.Context,
	req VerificationVerifyCompanyRequest,
) (*VerificationVerifyCompanyResponse, *http.Response, error) {
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

	var response VerificationVerifyCompanyResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}
