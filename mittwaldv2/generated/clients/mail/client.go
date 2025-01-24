package mail

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/mailmigrationv1"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/mailv1"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	DeprecatedUpdateDescription(
		ctx context.Context,
		req DeprecatedUpdateDescriptionRequest,
	) (*http.Response, error)
	DeprecatedUpdatePassword(
		ctx context.Context,
		req DeprecatedUpdatePasswordRequest,
	) (*http.Response, error)
	DeprecatedMailaddressUpdateAddress(
		ctx context.Context,
		req DeprecatedMailaddressUpdateAddressRequest,
	) (*http.Response, error)
	DeprecatedProjectsettingUpdateBlacklist(
		ctx context.Context,
		req DeprecatedProjectsettingUpdateBlacklistRequest,
	) (*http.Response, error)
	DeprecatedProjectsettingUpdateWhitelist(
		ctx context.Context,
		req DeprecatedProjectsettingUpdateWhitelistRequest,
	) (*http.Response, error)
	DeprecatedUpdateMailAddressAutoresponder(
		ctx context.Context,
		req DeprecatedUpdateMailAddressAutoresponderRequest,
	) (*http.Response, error)
	UpdateMailAddressAutoresponder(
		ctx context.Context,
		req UpdateMailAddressAutoresponderRequest,
	) (*http.Response, error)
	DeprecatedUpdateMailAddressCatchall(
		ctx context.Context,
		req DeprecatedUpdateMailAddressCatchallRequest,
	) (*http.Response, error)
	DeprecatedUpdateMailAddressForwardAddresses(
		ctx context.Context,
		req DeprecatedUpdateMailAddressForwardAddressesRequest,
	) (*http.Response, error)
	UpdateMailAddressForwardAddresses(
		ctx context.Context,
		req UpdateMailAddressForwardAddressesRequest,
	) (*http.Response, error)
	DeprecatedUpdateMailAddressPassword(
		ctx context.Context,
		req DeprecatedUpdateMailAddressPasswordRequest,
	) (*http.Response, error)
	UpdateMailAddressPassword(
		ctx context.Context,
		req UpdateMailAddressPasswordRequest,
	) (*http.Response, error)
	DeprecatedUpdateMailAddressQuota(
		ctx context.Context,
		req DeprecatedUpdateMailAddressQuotaRequest,
	) (*http.Response, error)
	UpdateMailAddressQuota(
		ctx context.Context,
		req UpdateMailAddressQuotaRequest,
	) (*http.Response, error)
	DeprecatedUpdateMailAddressSpamProtection(
		ctx context.Context,
		req DeprecatedUpdateMailAddressSpamProtectionRequest,
	) (*http.Response, error)
	UpdateMailAddressSpamProtection(
		ctx context.Context,
		req UpdateMailAddressSpamProtectionRequest,
	) (*http.Response, error)
	DeprecatedUpdateProjectMailSetting(
		ctx context.Context,
		req DeprecatedUpdateProjectMailSettingRequest,
	) (*http.Response, error)
	ListDeliveryBoxes(
		ctx context.Context,
		req ListDeliveryBoxesRequest,
	) (*[]mailv1.Deliverybox, *http.Response, error)
	CreateDeliverybox(
		ctx context.Context,
		req CreateDeliveryboxRequest,
	) (*CreateDeliveryboxResponse, *http.Response, error)
	ListMailAddresses(
		ctx context.Context,
		req ListMailAddressesRequest,
	) (*[]mailv1.MailAddress, *http.Response, error)
	CreateMailAddress(
		ctx context.Context,
		req CreateMailAddressRequest,
	) (*CreateMailAddressResponse, *http.Response, error)
	GetDeliveryBox(
		ctx context.Context,
		req GetDeliveryBoxRequest,
	) (*mailv1.Deliverybox, *http.Response, error)
	DeleteDeliveryBox(
		ctx context.Context,
		req DeleteDeliveryBoxRequest,
	) (*http.Response, error)
	GetMailAddress(
		ctx context.Context,
		req GetMailAddressRequest,
	) (*mailv1.MailAddress, *http.Response, error)
	DeleteMailAddress(
		ctx context.Context,
		req DeleteMailAddressRequest,
	) (*http.Response, error)
	ListProjectMailSettings(
		ctx context.Context,
		req ListProjectMailSettingsRequest,
	) (*ListProjectMailSettingsResponse, *http.Response, error)
	MigrationCheckMigrationIsPossible(
		ctx context.Context,
		req MigrationCheckMigrationIsPossibleRequest,
	) (*mailmigrationv1.CheckMigrationIsPossibleErrorResponse, *http.Response, error)
	MigrationGetMigration(
		ctx context.Context,
		req MigrationGetMigrationRequest,
	) (*mailmigrationv1.Migration, *http.Response, error)
	MigrationListMigrations(
		ctx context.Context,
		req MigrationListMigrationsRequest,
	) (*[]mailmigrationv1.Migration, *http.Response, error)
	MigrationRequestMailMigration(
		ctx context.Context,
		req MigrationRequestMailMigrationRequest,
	) (*http.Response, error)
	UpdateDeliveryBoxDescription(
		ctx context.Context,
		req UpdateDeliveryBoxDescriptionRequest,
	) (*http.Response, error)
	UpdateDeliveryBoxPassword(
		ctx context.Context,
		req UpdateDeliveryBoxPasswordRequest,
	) (*http.Response, error)
	UpdateMailAddressAddress(
		ctx context.Context,
		req UpdateMailAddressAddressRequest,
	) (*http.Response, error)
	UpdateMailAddressCatchAll(
		ctx context.Context,
		req UpdateMailAddressCatchAllRequest,
	) (*http.Response, error)
	UpdateProjectMailSetting(
		ctx context.Context,
		req UpdateProjectMailSettingRequest,
	) (*http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// Update the description of an deliverybox
//
// This operation is deprecated. Use the PATCH v2/delivery-boxes/{deliveryBoxId}/description endpoint instead.
func (c *clientImpl) DeprecatedUpdateDescription(
	ctx context.Context,
	req DeprecatedUpdateDescriptionRequest,
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

// Update the password for a specific deliverybox
//
// This operation is deprecated. Use the PATCH v2/delivery-boxes/{deliveryBoxId}/password endpoint instead.
func (c *clientImpl) DeprecatedUpdatePassword(
	ctx context.Context,
	req DeprecatedUpdatePasswordRequest,
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

// Update mail-address
//
// This operation is deprecated. Use the PATCH v2/mail-addresses/{mailAddressId}/address endpoint instead.
func (c *clientImpl) DeprecatedMailaddressUpdateAddress(
	ctx context.Context,
	req DeprecatedMailaddressUpdateAddressRequest,
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

// Update blacklist for a given project ID
//
// This operation is deprecated. Use the PATCH v2/{projectId}/mail-settings/{mailSetting} endpoint instead.
func (c *clientImpl) DeprecatedProjectsettingUpdateBlacklist(
	ctx context.Context,
	req DeprecatedProjectsettingUpdateBlacklistRequest,
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

// Update whitelist for a given project ID
//
// This operation is deprecated. Use the PATCH v2/{projectId}/mail-settings/{mailSetting} endpoint instead.
func (c *clientImpl) DeprecatedProjectsettingUpdateWhitelist(
	ctx context.Context,
	req DeprecatedProjectsettingUpdateWhitelistRequest,
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

// Update the autoresponder of a MailAddress.
//
// This operation is deprecated. Use the PATCH v2/mail-addresses/{mailAddressId}/autoresponder endpoint instead.
func (c *clientImpl) DeprecatedUpdateMailAddressAutoresponder(
	ctx context.Context,
	req DeprecatedUpdateMailAddressAutoresponderRequest,
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

// Update the autoresponder of a MailAddress.
func (c *clientImpl) UpdateMailAddressAutoresponder(
	ctx context.Context,
	req UpdateMailAddressAutoresponderRequest,
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

// Update the catchall of a MailAddress.
//
// This operation is deprecated. Use the PATCH v2/mail-addresses/{mailAddressId}/catch-all endpoint instead.
func (c *clientImpl) DeprecatedUpdateMailAddressCatchall(
	ctx context.Context,
	req DeprecatedUpdateMailAddressCatchallRequest,
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

// Update the forward addresses of a MailAddresses.
//
// This operation is deprecated. Use the PATCH v2/mail-addresses/{mailAddressId}/forward-addresses endpoint instead.
func (c *clientImpl) DeprecatedUpdateMailAddressForwardAddresses(
	ctx context.Context,
	req DeprecatedUpdateMailAddressForwardAddressesRequest,
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

// Update the forward addresses of a MailAddresses.
func (c *clientImpl) UpdateMailAddressForwardAddresses(
	ctx context.Context,
	req UpdateMailAddressForwardAddressesRequest,
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

// Update the password for a MailAddress.
//
// This operation is deprecated. Use the PATCH v2/mail-addresses/{mailAddressId}/password endpoint instead.
func (c *clientImpl) DeprecatedUpdateMailAddressPassword(
	ctx context.Context,
	req DeprecatedUpdateMailAddressPasswordRequest,
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

// Update the password for a MailAddress.
func (c *clientImpl) UpdateMailAddressPassword(
	ctx context.Context,
	req UpdateMailAddressPasswordRequest,
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

// Update the quota of a MailAddress.
//
// This operation is deprecated. Use the PATCH v2/mail-addresses/{mailAddressId}/quota endpoint instead.
func (c *clientImpl) DeprecatedUpdateMailAddressQuota(
	ctx context.Context,
	req DeprecatedUpdateMailAddressQuotaRequest,
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

// Update the quota of a MailAddress.
func (c *clientImpl) UpdateMailAddressQuota(
	ctx context.Context,
	req UpdateMailAddressQuotaRequest,
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

// Update the spam protection of a MailAddress.
//
// This operation is deprecated. Use the PATCH v2/mail-addresses/{mailAddressId}/spam-protection endpoint instead.
func (c *clientImpl) DeprecatedUpdateMailAddressSpamProtection(
	ctx context.Context,
	req DeprecatedUpdateMailAddressSpamProtectionRequest,
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

// Update the spam protection of a MailAddress.
func (c *clientImpl) UpdateMailAddressSpamProtection(
	ctx context.Context,
	req UpdateMailAddressSpamProtectionRequest,
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

// Update a mail setting of a Project.
//
// This operation is deprecated. Use the PATCH v2/{projectId}/mail-settings/{mailSetting} endpoint instead.
func (c *clientImpl) DeprecatedUpdateProjectMailSetting(
	ctx context.Context,
	req DeprecatedUpdateProjectMailSettingRequest,
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

// List DeliveryBoxes belonging to a Project.
func (c *clientImpl) ListDeliveryBoxes(
	ctx context.Context,
	req ListDeliveryBoxesRequest,
) (*[]mailv1.Deliverybox, *http.Response, error) {
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

	var response []mailv1.Deliverybox
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create a DeliveryBox.
func (c *clientImpl) CreateDeliverybox(
	ctx context.Context,
	req CreateDeliveryboxRequest,
) (*CreateDeliveryboxResponse, *http.Response, error) {
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

	var response CreateDeliveryboxResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List MailAddresses belonging to a Project.
func (c *clientImpl) ListMailAddresses(
	ctx context.Context,
	req ListMailAddressesRequest,
) (*[]mailv1.MailAddress, *http.Response, error) {
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

	var response []mailv1.MailAddress
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create a MailAddress.
func (c *clientImpl) CreateMailAddress(
	ctx context.Context,
	req CreateMailAddressRequest,
) (*CreateMailAddressResponse, *http.Response, error) {
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

	var response CreateMailAddressResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a DeliveryBox.
func (c *clientImpl) GetDeliveryBox(
	ctx context.Context,
	req GetDeliveryBoxRequest,
) (*mailv1.Deliverybox, *http.Response, error) {
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

	var response mailv1.Deliverybox
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a DeliveryBox.
func (c *clientImpl) DeleteDeliveryBox(
	ctx context.Context,
	req DeleteDeliveryBoxRequest,
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

// Get a MailAddress.
func (c *clientImpl) GetMailAddress(
	ctx context.Context,
	req GetMailAddressRequest,
) (*mailv1.MailAddress, *http.Response, error) {
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

	var response mailv1.MailAddress
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a MailAddress.
func (c *clientImpl) DeleteMailAddress(
	ctx context.Context,
	req DeleteMailAddressRequest,
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

// List mail settings of a Project.
func (c *clientImpl) ListProjectMailSettings(
	ctx context.Context,
	req ListProjectMailSettingsRequest,
) (*ListProjectMailSettingsResponse, *http.Response, error) {
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

	var response ListProjectMailSettingsResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Check if a Migration between two projects is possible.
func (c *clientImpl) MigrationCheckMigrationIsPossible(
	ctx context.Context,
	req MigrationCheckMigrationIsPossibleRequest,
) (*mailmigrationv1.CheckMigrationIsPossibleErrorResponse, *http.Response, error) {
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

	var response mailmigrationv1.CheckMigrationIsPossibleErrorResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a Migration.
func (c *clientImpl) MigrationGetMigration(
	ctx context.Context,
	req MigrationGetMigrationRequest,
) (*mailmigrationv1.Migration, *http.Response, error) {
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

	var response mailmigrationv1.Migration
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List Migrations belonging to a Project in customer center or mStudio.
//
// In case you want to list the Migrations for a p-Account you will have to use `commons.LegacyBearerAuthentication` and `commons.AccessToken` for authentication.
// If you want to list the Migrations for a mStudio-project you will have to use `commons.AccessToken` for authentication.
func (c *clientImpl) MigrationListMigrations(
	ctx context.Context,
	req MigrationListMigrationsRequest,
) (*[]mailmigrationv1.Migration, *http.Response, error) {
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

	var response []mailmigrationv1.Migration
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Request a Mail Migration between two projects.
func (c *clientImpl) MigrationRequestMailMigration(
	ctx context.Context,
	req MigrationRequestMailMigrationRequest,
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

// Update the description of a DeliveryBox.
func (c *clientImpl) UpdateDeliveryBoxDescription(
	ctx context.Context,
	req UpdateDeliveryBoxDescriptionRequest,
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

// Update the password of a DeliveryBox.
func (c *clientImpl) UpdateDeliveryBoxPassword(
	ctx context.Context,
	req UpdateDeliveryBoxPasswordRequest,
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

// Update a MailAddress.
func (c *clientImpl) UpdateMailAddressAddress(
	ctx context.Context,
	req UpdateMailAddressAddressRequest,
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

// Update the catchall of a MailAddress.
func (c *clientImpl) UpdateMailAddressCatchAll(
	ctx context.Context,
	req UpdateMailAddressCatchAllRequest,
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

// Update a mail setting of a Project.
func (c *clientImpl) UpdateProjectMailSetting(
	ctx context.Context,
	req UpdateProjectMailSettingRequest,
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
