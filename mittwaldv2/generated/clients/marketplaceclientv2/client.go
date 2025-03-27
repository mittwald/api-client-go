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
		reqEditors ...func(req *http.Request) error,
	) (*GetCustomerBillingPortalLinkResponse, *http.Response, error)
	GetLoginLink(
		ctx context.Context,
		req GetLoginLinkRequest,
		reqEditors ...func(req *http.Request) error,
	) (*GetLoginLinkResponse, *http.Response, error)
	RotateSecretForExtensionInstance(
		ctx context.Context,
		req RotateSecretForExtensionInstanceRequest,
		reqEditors ...func(req *http.Request) error,
	) (*RotateSecretForExtensionInstanceResponse, *http.Response, error)
	AuthenticateInstance(
		ctx context.Context,
		req AuthenticateInstanceRequest,
		reqEditors ...func(req *http.Request) error,
	) (*AuthenticateInstanceResponse, *http.Response, error)
	AuthenticateWithSessionToken(
		ctx context.Context,
		req AuthenticateWithSessionTokenRequest,
		reqEditors ...func(req *http.Request) error,
	) (*AuthenticateWithSessionTokenResponse, *http.Response, error)
	ConsentToExtensionScopes(
		ctx context.Context,
		req ConsentToExtensionScopesRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	CreateContributorOnboardingProcess(
		ctx context.Context,
		req CreateContributorOnboardingProcessRequest,
		reqEditors ...func(req *http.Request) error,
	) (*CreateContributorOnboardingProcessResponse, *http.Response, error)
	ListExtensionInstances(
		ctx context.Context,
		req ListExtensionInstancesRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]marketplacev2.ExtensionInstance, *http.Response, error)
	CreateExtensionInstance(
		ctx context.Context,
		req CreateExtensionInstanceRequest,
		reqEditors ...func(req *http.Request) error,
	) (*CreateExtensionInstanceResponse, *http.Response, error)
	CreateRetrievalKey(
		ctx context.Context,
		req CreateRetrievalKeyRequest,
		reqEditors ...func(req *http.Request) error,
	) (*CreateRetrievalKeyResponse, *http.Response, error)
	GetExtensionInstance(
		ctx context.Context,
		req GetExtensionInstanceRequest,
		reqEditors ...func(req *http.Request) error,
	) (*marketplacev2.ExtensionInstance, *http.Response, error)
	DeleteExtensionInstance(
		ctx context.Context,
		req DeleteExtensionInstanceRequest,
		reqEditors ...func(req *http.Request) error,
	) (*any, *http.Response, error)
	GetOwnExtension(
		ctx context.Context,
		req GetOwnExtensionRequest,
		reqEditors ...func(req *http.Request) error,
	) (*marketplacev2.OwnExtension, *http.Response, error)
	DeleteExtension(
		ctx context.Context,
		req DeleteExtensionRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	PatchExtension(
		ctx context.Context,
		req PatchExtensionRequest,
		reqEditors ...func(req *http.Request) error,
	) (*marketplacev2.OwnExtension, *http.Response, error)
	DisableExtensionInstance(
		ctx context.Context,
		req DisableExtensionInstanceRequest,
		reqEditors ...func(req *http.Request) error,
	) (*any, *http.Response, error)
	DryRunWebhook(
		ctx context.Context,
		req DryRunWebhookRequest,
		reqEditors ...func(req *http.Request) error,
	) (*DryRunWebhookResponse, *http.Response, error)
	EnableExtensionInstance(
		ctx context.Context,
		req EnableExtensionInstanceRequest,
		reqEditors ...func(req *http.Request) error,
	) (*any, *http.Response, error)
	GenerateSessionToken(
		ctx context.Context,
		req GenerateSessionTokenRequest,
		reqEditors ...func(req *http.Request) error,
	) (*GenerateSessionTokenResponse, *http.Response, error)
	GetContributor(
		ctx context.Context,
		req GetContributorRequest,
		reqEditors ...func(req *http.Request) error,
	) (*marketplacev2.Contributor, *http.Response, error)
	GetExtensionInstanceForCustomer(
		ctx context.Context,
		req GetExtensionInstanceForCustomerRequest,
		reqEditors ...func(req *http.Request) error,
	) (*marketplacev2.ExtensionInstance, *http.Response, error)
	GetExtensionInstanceForProject(
		ctx context.Context,
		req GetExtensionInstanceForProjectRequest,
		reqEditors ...func(req *http.Request) error,
	) (*marketplacev2.ExtensionInstance, *http.Response, error)
	GetExtension(
		ctx context.Context,
		req GetExtensionRequest,
		reqEditors ...func(req *http.Request) error,
	) (*marketplacev2.Extension, *http.Response, error)
	GetPublicKey(
		ctx context.Context,
		req GetPublicKeyRequest,
		reqEditors ...func(req *http.Request) error,
	) (*marketplacev2.PublicKey, *http.Response, error)
	ListContributors(
		ctx context.Context,
		req ListContributorsRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]marketplacev2.Contributor, *http.Response, error)
	ListExtensions(
		ctx context.Context,
		req ListExtensionsRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]marketplacev2.Extension, *http.Response, error)
	ListOwnExtensions(
		ctx context.Context,
		req ListOwnExtensionsRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]marketplacev2.OwnExtension, *http.Response, error)
	RegisterExtension(
		ctx context.Context,
		req RegisterExtensionRequest,
		reqEditors ...func(req *http.Request) error,
	) (*RegisterExtensionResponse, *http.Response, error)
	RemoveAsset(
		ctx context.Context,
		req RemoveAssetRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	RequestAssetUpload(
		ctx context.Context,
		req RequestAssetUploadRequest,
		reqEditors ...func(req *http.Request) error,
	) (*RequestAssetUploadResponse, *http.Response, error)
	RequestExtensionVerification(
		ctx context.Context,
		req RequestExtensionVerificationRequest,
		reqEditors ...func(req *http.Request) error,
	) (*RequestExtensionVerificationResponse, *http.Response, error)
	RequestLogoUpload(
		ctx context.Context,
		req RequestLogoUploadRequest,
		reqEditors ...func(req *http.Request) error,
	) (*RequestLogoUploadResponse, *http.Response, error)
	SetExtensionPublishedState(
		ctx context.Context,
		req SetExtensionPublishedStateRequest,
		reqEditors ...func(req *http.Request) error,
	) (*SetExtensionPublishedStateResponse, *http.Response, error)
	UpdateExtensionInstanceContract(
		ctx context.Context,
		req UpdateExtensionInstanceContractRequest,
		reqEditors ...func(req *http.Request) error,
	) (*UpdateExtensionInstanceContractResponse, *http.Response, error)
	UpdateExtensionPricing(
		ctx context.Context,
		req UpdateExtensionPricingRequest,
		reqEditors ...func(req *http.Request) error,
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
	reqEditors ...func(req *http.Request) error,
) (*GetCustomerBillingPortalLinkResponse, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*GetLoginLinkResponse, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*RotateSecretForExtensionInstanceResponse, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*AuthenticateInstanceResponse, *http.Response, error) {
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

	var response AuthenticateInstanceResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Authenticate your external application using a session token and an extension secret
func (c *clientImpl) AuthenticateWithSessionToken(
	ctx context.Context,
	req AuthenticateWithSessionTokenRequest,
	reqEditors ...func(req *http.Request) error,
) (*AuthenticateWithSessionTokenResponse, *http.Response, error) {
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

	var response AuthenticateWithSessionTokenResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Consent to extension scopes.
func (c *clientImpl) ConsentToExtensionScopes(
	ctx context.Context,
	req ConsentToExtensionScopesRequest,
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

// Create the OnboardingProcess of a Contributor.
//
// The OnboardingProcess is needed to publish paid extensions.
func (c *clientImpl) CreateContributorOnboardingProcess(
	ctx context.Context,
	req CreateContributorOnboardingProcessRequest,
	reqEditors ...func(req *http.Request) error,
) (*CreateContributorOnboardingProcessResponse, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*[]marketplacev2.ExtensionInstance, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*CreateExtensionInstanceResponse, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*CreateRetrievalKeyResponse, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*marketplacev2.ExtensionInstance, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*any, *http.Response, error) {
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

	var response any
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get Extension of own contributor.
func (c *clientImpl) GetOwnExtension(
	ctx context.Context,
	req GetOwnExtensionRequest,
	reqEditors ...func(req *http.Request) error,
) (*marketplacev2.OwnExtension, *http.Response, error) {
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

	var response marketplacev2.OwnExtension
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete an extension.
//
// This action deletes all ExtensionInstances and afterwards the Extension itself.
func (c *clientImpl) DeleteExtension(
	ctx context.Context,
	req DeleteExtensionRequest,
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

// Patch Extension.
func (c *clientImpl) PatchExtension(
	ctx context.Context,
	req PatchExtensionRequest,
	reqEditors ...func(req *http.Request) error,
) (*marketplacev2.OwnExtension, *http.Response, error) {
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

	var response marketplacev2.OwnExtension
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Disable an ExtensionInstance.
func (c *clientImpl) DisableExtensionInstance(
	ctx context.Context,
	req DisableExtensionInstanceRequest,
	reqEditors ...func(req *http.Request) error,
) (*any, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*DryRunWebhookResponse, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*any, *http.Response, error) {
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

	var response any
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Generate a session token to transmit it to the extensions frontend fragment.
func (c *clientImpl) GenerateSessionToken(
	ctx context.Context,
	req GenerateSessionTokenRequest,
	reqEditors ...func(req *http.Request) error,
) (*GenerateSessionTokenResponse, *http.Response, error) {
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

	var response GenerateSessionTokenResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a Contributor.
func (c *clientImpl) GetContributor(
	ctx context.Context,
	req GetContributorRequest,
	reqEditors ...func(req *http.Request) error,
) (*marketplacev2.Contributor, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*marketplacev2.ExtensionInstance, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*marketplacev2.ExtensionInstance, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*marketplacev2.Extension, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*marketplacev2.PublicKey, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*[]marketplacev2.Contributor, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*[]marketplacev2.Extension, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*[]marketplacev2.OwnExtension, *http.Response, error) {
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

	var response []marketplacev2.OwnExtension
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Register an Extension.
func (c *clientImpl) RegisterExtension(
	ctx context.Context,
	req RegisterExtensionRequest,
	reqEditors ...func(req *http.Request) error,
) (*RegisterExtensionResponse, *http.Response, error) {
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

	var response RegisterExtensionResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Remove an asset of an extension.
func (c *clientImpl) RemoveAsset(
	ctx context.Context,
	req RemoveAssetRequest,
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

// Add an asset to an extension.
func (c *clientImpl) RequestAssetUpload(
	ctx context.Context,
	req RequestAssetUploadRequest,
	reqEditors ...func(req *http.Request) error,
) (*RequestAssetUploadResponse, *http.Response, error) {
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

	var response RequestAssetUploadResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Start the verification process of an Extension.
func (c *clientImpl) RequestExtensionVerification(
	ctx context.Context,
	req RequestExtensionVerificationRequest,
	reqEditors ...func(req *http.Request) error,
) (*RequestExtensionVerificationResponse, *http.Response, error) {
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

	var response RequestExtensionVerificationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Add a logo to an extension.
func (c *clientImpl) RequestLogoUpload(
	ctx context.Context,
	req RequestLogoUploadRequest,
	reqEditors ...func(req *http.Request) error,
) (*RequestLogoUploadResponse, *http.Response, error) {
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

	var response RequestLogoUploadResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Publish or withdraw an Extension.
func (c *clientImpl) SetExtensionPublishedState(
	ctx context.Context,
	req SetExtensionPublishedStateRequest,
	reqEditors ...func(req *http.Request) error,
) (*SetExtensionPublishedStateResponse, *http.Response, error) {
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

	var response SetExtensionPublishedStateResponse
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
	reqEditors ...func(req *http.Request) error,
) (*UpdateExtensionInstanceContractResponse, *http.Response, error) {
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
	reqEditors ...func(req *http.Request) error,
) (*UpdateExtensionPricingResponse, *http.Response, error) {
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

	var response UpdateExtensionPricingResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}
