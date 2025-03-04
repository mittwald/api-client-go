package marketplaceclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/marketplacev2"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

type Client interface {
	GetCustomerBillingPortalLink(
		ctx context.Context,
		req GetCustomerBillingPortalLinkRequest,
	) (*GetCustomerBillingPortalLinkResponse, *http.Response, error)
	GetLoginLink(
		ctx context.Context,
		req GetLoginLinkRequest,
	) (*GetLoginLinkResponse, *http.Response, error)
	RotateSecretForExtensionInstance(
		ctx context.Context,
		req RotateSecretForExtensionInstanceRequest,
	) (*RotateSecretForExtensionInstanceResponse, *http.Response, error)
	AuthenticateInstance(
		ctx context.Context,
		req AuthenticateInstanceRequest,
	) (*AuthenticateInstanceResponse, *http.Response, error)
	ConsentToExtensionScopes(
		ctx context.Context,
		req ConsentToExtensionScopesRequest,
	) (*http.Response, error)
	CreateContributorOnboardingProcess(
		ctx context.Context,
		req CreateContributorOnboardingProcessRequest,
	) (*CreateContributorOnboardingProcessResponse, *http.Response, error)
	ListExtensionInstances(
		ctx context.Context,
		req ListExtensionInstancesRequest,
	) (*[]marketplacev2.ExtensionInstance, *http.Response, error)
	CreateExtensionInstance(
		ctx context.Context,
		req CreateExtensionInstanceRequest,
	) (*CreateExtensionInstanceResponse, *http.Response, error)
	CreateRetrievalKey(
		ctx context.Context,
		req CreateRetrievalKeyRequest,
	) (*CreateRetrievalKeyResponse, *http.Response, error)
	GetExtensionInstance(
		ctx context.Context,
		req GetExtensionInstanceRequest,
	) (*marketplacev2.ExtensionInstance, *http.Response, error)
	DeleteExtensionInstance(
		ctx context.Context,
		req DeleteExtensionInstanceRequest,
	) (*any, *http.Response, error)
	DisableExtensionInstance(
		ctx context.Context,
		req DisableExtensionInstanceRequest,
	) (*any, *http.Response, error)
	DryRunWebhook(
		ctx context.Context,
		req DryRunWebhookRequest,
	) (*DryRunWebhookResponse, *http.Response, error)
	EnableExtensionInstance(
		ctx context.Context,
		req EnableExtensionInstanceRequest,
	) (*any, *http.Response, error)
	GetContributor(
		ctx context.Context,
		req GetContributorRequest,
	) (*marketplacev2.Contributor, *http.Response, error)
	GetExtensionInstanceForCustomer(
		ctx context.Context,
		req GetExtensionInstanceForCustomerRequest,
	) (*marketplacev2.ExtensionInstance, *http.Response, error)
	GetExtensionInstanceForProject(
		ctx context.Context,
		req GetExtensionInstanceForProjectRequest,
	) (*marketplacev2.ExtensionInstance, *http.Response, error)
	GetExtension(
		ctx context.Context,
		req GetExtensionRequest,
	) (*marketplacev2.Extension, *http.Response, error)
	GetPublicKey(
		ctx context.Context,
		req GetPublicKeyRequest,
	) (*marketplacev2.PublicKey, *http.Response, error)
	ListContributors(
		ctx context.Context,
		req ListContributorsRequest,
	) (*[]marketplacev2.Contributor, *http.Response, error)
	ListExtensions(
		ctx context.Context,
		req ListExtensionsRequest,
	) (*[]marketplacev2.Extension, *http.Response, error)
	ListOwnExtensions(
		ctx context.Context,
		req ListOwnExtensionsRequest,
	) (*[]marketplacev2.OwnExtension, *http.Response, error)
	UpdateExtensionInstanceContract(
		ctx context.Context,
		req UpdateExtensionInstanceContractRequest,
	) (*UpdateExtensionInstanceContractResponse, *http.Response, error)
	UpdateExtensionPricing(
		ctx context.Context,
		req UpdateExtensionPricingRequest,
	) (*UpdateExtensionPricingResponse, *http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// Get the Stripe Billing Portal Link for a Customer
//
// Get the Stripe Billing Portal Link for a Customer.
func (c *clientImpl) GetCustomerBillingPortalLink(
	ctx context.Context,
	req GetCustomerBillingPortalLinkRequest,
) (*GetCustomerBillingPortalLinkResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response GetCustomerBillingPortalLinkResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get the Stripe Dashboard Link for a Contributor.
//
// Get the Stripe Dashboard Link for a Contributor.
func (c *clientImpl) GetLoginLink(
	ctx context.Context,
	req GetLoginLinkRequest,
) (*GetLoginLinkResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response GetLoginLinkResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Rotate the secret for an extension instance.
func (c *clientImpl) RotateSecretForExtensionInstance(
	ctx context.Context,
	req RotateSecretForExtensionInstanceRequest,
) (*RotateSecretForExtensionInstanceResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response RotateSecretForExtensionInstanceResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Authenticate your external application using the extensionInstanceSecret.
func (c *clientImpl) AuthenticateInstance(
	ctx context.Context,
	req AuthenticateInstanceRequest,
) (*AuthenticateInstanceResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response AuthenticateInstanceResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Consent to extension scopes.
func (c *clientImpl) ConsentToExtensionScopes(
	ctx context.Context,
	req ConsentToExtensionScopesRequest,
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
		err := httperr.ErrFromResponse(httpRes)
		return httpRes, err
	}

	return httpRes, nil
}

// Create the OnboardingProcess of a Contributor.
//
// The OnboardingProcess is needed to publish paid extensions.
func (c *clientImpl) CreateContributorOnboardingProcess(
	ctx context.Context,
	req CreateContributorOnboardingProcessRequest,
) (*CreateContributorOnboardingProcessResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response CreateContributorOnboardingProcessResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List ExtensionInstances.
func (c *clientImpl) ListExtensionInstances(
	ctx context.Context,
	req ListExtensionInstancesRequest,
) (*[]marketplacev2.ExtensionInstance, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []marketplacev2.ExtensionInstance
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create an ExtensionInstance.
func (c *clientImpl) CreateExtensionInstance(
	ctx context.Context,
	req CreateExtensionInstanceRequest,
) (*CreateExtensionInstanceResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response CreateExtensionInstanceResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create an access token retrieval key for an extension instance.
func (c *clientImpl) CreateRetrievalKey(
	ctx context.Context,
	req CreateRetrievalKeyRequest,
) (*CreateRetrievalKeyResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response CreateRetrievalKeyResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get an ExtensionInstance.
func (c *clientImpl) GetExtensionInstance(
	ctx context.Context,
	req GetExtensionInstanceRequest,
) (*marketplacev2.ExtensionInstance, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response marketplacev2.ExtensionInstance
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete an ExtensionInstance.
func (c *clientImpl) DeleteExtensionInstance(
	ctx context.Context,
	req DeleteExtensionInstanceRequest,
) (*any, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response any
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Disable an ExtensionInstance.
func (c *clientImpl) DisableExtensionInstance(
	ctx context.Context,
	req DisableExtensionInstanceRequest,
) (*any, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response any
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Dry run a webhook with random or given values.
func (c *clientImpl) DryRunWebhook(
	ctx context.Context,
	req DryRunWebhookRequest,
) (*DryRunWebhookResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response DryRunWebhookResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Enable an ExtensionInstance.
func (c *clientImpl) EnableExtensionInstance(
	ctx context.Context,
	req EnableExtensionInstanceRequest,
) (*any, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response any
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a Contributor.
func (c *clientImpl) GetContributor(
	ctx context.Context,
	req GetContributorRequest,
) (*marketplacev2.Contributor, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response marketplacev2.Contributor
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get the ExtensionInstance of a specific customer and extension, if existing.
func (c *clientImpl) GetExtensionInstanceForCustomer(
	ctx context.Context,
	req GetExtensionInstanceForCustomerRequest,
) (*marketplacev2.ExtensionInstance, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response marketplacev2.ExtensionInstance
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get the ExtensionInstance of a specific project and extension, if existing.
func (c *clientImpl) GetExtensionInstanceForProject(
	ctx context.Context,
	req GetExtensionInstanceForProjectRequest,
) (*marketplacev2.ExtensionInstance, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response marketplacev2.ExtensionInstance
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get an Extension.
func (c *clientImpl) GetExtension(
	ctx context.Context,
	req GetExtensionRequest,
) (*marketplacev2.Extension, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response marketplacev2.Extension
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get the public key to verify the webhook signature.
func (c *clientImpl) GetPublicKey(
	ctx context.Context,
	req GetPublicKeyRequest,
) (*marketplacev2.PublicKey, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response marketplacev2.PublicKey
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List Contributors.
func (c *clientImpl) ListContributors(
	ctx context.Context,
	req ListContributorsRequest,
) (*[]marketplacev2.Contributor, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []marketplacev2.Contributor
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List Extensions.
func (c *clientImpl) ListExtensions(
	ctx context.Context,
	req ListExtensionsRequest,
) (*[]marketplacev2.Extension, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []marketplacev2.Extension
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List Extensions of own contributor.
func (c *clientImpl) ListOwnExtensions(
	ctx context.Context,
	req ListOwnExtensionsRequest,
) (*[]marketplacev2.OwnExtension, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []marketplacev2.OwnExtension
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Update or Create Contract for existing Extension Instances.
//
// Call to update Contract for existing Extension Instances. For example to accept a new Pricing.
func (c *clientImpl) UpdateExtensionInstanceContract(
	ctx context.Context,
	req UpdateExtensionInstanceContractRequest,
) (*UpdateExtensionInstanceContractResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response UpdateExtensionInstanceContractResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Creates or Updates Pricing for an Extension.
//
// The Pricing is needed to publish paid extensions.
func (c *clientImpl) UpdateExtensionPricing(
	ctx context.Context,
	req UpdateExtensionPricingRequest,
) (*UpdateExtensionPricingResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response UpdateExtensionPricingResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}
