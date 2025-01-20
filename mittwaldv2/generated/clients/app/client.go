package app

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/appv1"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	ExecuteAction(
		ctx context.Context,
		req ExecuteActionRequest,
	) (*http.Response, error)
	GetApp(
		ctx context.Context,
		req GetAppRequest,
	) (*appv1.App, *http.Response, error)
	GetAppinstallation(
		ctx context.Context,
		req GetAppinstallationRequest,
	) (*appv1.AppInstallation, *http.Response, error)
	UninstallAppinstallation(
		ctx context.Context,
		req UninstallAppinstallationRequest,
	) (*http.Response, error)
	PatchAppinstallation(
		ctx context.Context,
		req PatchAppinstallationRequest,
	) (*http.Response, error)
	GetAppversion(
		ctx context.Context,
		req GetAppversionRequest,
	) (*appv1.AppVersion, *http.Response, error)
	GetInstalledSystemsoftwareForAppinstallation(
		ctx context.Context,
		req GetInstalledSystemsoftwareForAppinstallationRequest,
	) (*[]appv1.SystemSoftware, *http.Response, error)
	GetMissingDependenciesForAppinstallation(
		ctx context.Context,
		req GetMissingDependenciesForAppinstallationRequest,
	) (*GetMissingDependenciesForAppinstallationResponse, *http.Response, error)
	GetSystemsoftware(
		ctx context.Context,
		req GetSystemsoftwareRequest,
	) (*appv1.SystemSoftware, *http.Response, error)
	GetSystemsoftwareversion(
		ctx context.Context,
		req GetSystemsoftwareversionRequest,
	) (*appv1.SystemSoftwareVersion, *http.Response, error)
	LinkDatabase(
		ctx context.Context,
		req LinkDatabaseRequest,
	) (*http.Response, error)
	ListAppinstallationsForUser(
		ctx context.Context,
		req ListAppinstallationsForUserRequest,
	) (*[]appv1.AppInstallation, *http.Response, error)
	ListAppinstallations(
		ctx context.Context,
		req ListAppinstallationsRequest,
	) (*[]appv1.AppInstallation, *http.Response, error)
	RequestAppinstallation(
		ctx context.Context,
		req RequestAppinstallationRequest,
	) (*RequestAppinstallationResponse, *http.Response, error)
	ListApps(
		ctx context.Context,
		req ListAppsRequest,
	) (*[]appv1.App, *http.Response, error)
	ListAppversions(
		ctx context.Context,
		req ListAppversionsRequest,
	) (*[]appv1.AppVersion, *http.Response, error)
	ListSystemsoftwares(
		ctx context.Context,
		req ListSystemsoftwaresRequest,
	) (*[]appv1.SystemSoftware, *http.Response, error)
	ListSystemsoftwareversions(
		ctx context.Context,
		req ListSystemsoftwareversionsRequest,
	) (*[]appv1.SystemSoftwareVersion, *http.Response, error)
	ListUpdateCandidatesForAppversion(
		ctx context.Context,
		req ListUpdateCandidatesForAppversionRequest,
	) (*[]appv1.AppVersion, *http.Response, error)
	ReplaceDatabase(
		ctx context.Context,
		req ReplaceDatabaseRequest,
	) (*http.Response, error)
	RequestAppinstallationCopy(
		ctx context.Context,
		req RequestAppinstallationCopyRequest,
	) (*RequestAppinstallationCopyResponse, *http.Response, error)
	RetrieveStatus(
		ctx context.Context,
		req RetrieveStatusRequest,
	) (*appv1.AppInstallationStatus, *http.Response, error)
	UnlinkDatabase(
		ctx context.Context,
		req UnlinkDatabaseRequest,
	) (*http.Response, error)
	SetDatabaseUsers(
		ctx context.Context,
		req SetDatabaseUsersRequest,
	) (*http.Response, error)
	DeprecatedAppLinkDatabase(
		ctx context.Context,
		req DeprecatedAppLinkDatabaseRequest,
	) (*http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

//Trigger a runtime action belonging to an AppInstallation.
func (c *clientImpl) ExecuteAction(
	ctx context.Context,
	req ExecuteActionRequest,
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

//Get an App.
func (c *clientImpl) GetApp(
	ctx context.Context,
	req GetAppRequest,
) (*appv1.App, *http.Response, error) {
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

	var response appv1.App
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get an AppInstallation.
func (c *clientImpl) GetAppinstallation(
	ctx context.Context,
	req GetAppinstallationRequest,
) (*appv1.AppInstallation, *http.Response, error) {
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

	var response appv1.AppInstallation
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Trigger an uninstallation process for an AppInstallation.
func (c *clientImpl) UninstallAppinstallation(
	ctx context.Context,
	req UninstallAppinstallationRequest,
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

//Update properties belonging to an AppInstallation.
func (c *clientImpl) PatchAppinstallation(
	ctx context.Context,
	req PatchAppinstallationRequest,
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

//Get an AppVersion.
func (c *clientImpl) GetAppversion(
	ctx context.Context,
	req GetAppversionRequest,
) (*appv1.AppVersion, *http.Response, error) {
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

	var response appv1.AppVersion
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get the installed `SystemSoftware' for a specific `AppInstallation`.
func (c *clientImpl) GetInstalledSystemsoftwareForAppinstallation(
	ctx context.Context,
	req GetInstalledSystemsoftwareForAppinstallationRequest,
) (*[]appv1.SystemSoftware, *http.Response, error) {
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

	var response []appv1.SystemSoftware
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get the missing requirements of an appInstallation for a specific target AppVersion.
func (c *clientImpl) GetMissingDependenciesForAppinstallation(
	ctx context.Context,
	req GetMissingDependenciesForAppinstallationRequest,
) (*GetMissingDependenciesForAppinstallationResponse, *http.Response, error) {
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

	var response GetMissingDependenciesForAppinstallationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a SystemSoftware.
func (c *clientImpl) GetSystemsoftware(
	ctx context.Context,
	req GetSystemsoftwareRequest,
) (*appv1.SystemSoftware, *http.Response, error) {
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

	var response appv1.SystemSoftware
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a SystemSoftwareVersion.
func (c *clientImpl) GetSystemsoftwareversion(
	ctx context.Context,
	req GetSystemsoftwareversionRequest,
) (*appv1.SystemSoftwareVersion, *http.Response, error) {
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

	var response appv1.SystemSoftwareVersion
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Create linkage between an AppInstallation and a MySQLDatabase.
func (c *clientImpl) LinkDatabase(
	ctx context.Context,
	req LinkDatabaseRequest,
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

//List AppInstallations that a user has access to.
func (c *clientImpl) ListAppinstallationsForUser(
	ctx context.Context,
	req ListAppinstallationsForUserRequest,
) (*[]appv1.AppInstallation, *http.Response, error) {
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

	var response []appv1.AppInstallation
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//List AppInstallations belonging to a Project.
func (c *clientImpl) ListAppinstallations(
	ctx context.Context,
	req ListAppinstallationsRequest,
) (*[]appv1.AppInstallation, *http.Response, error) {
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

	var response []appv1.AppInstallation
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Request an AppInstallation.
func (c *clientImpl) RequestAppinstallation(
	ctx context.Context,
	req RequestAppinstallationRequest,
) (*RequestAppinstallationResponse, *http.Response, error) {
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

	var response RequestAppinstallationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//List Apps.
func (c *clientImpl) ListApps(
	ctx context.Context,
	req ListAppsRequest,
) (*[]appv1.App, *http.Response, error) {
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

	var response []appv1.App
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//List AppVersions belonging to an App.
func (c *clientImpl) ListAppversions(
	ctx context.Context,
	req ListAppversionsRequest,
) (*[]appv1.AppVersion, *http.Response, error) {
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

	var response []appv1.AppVersion
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//List SystemSoftwares.
func (c *clientImpl) ListSystemsoftwares(
	ctx context.Context,
	req ListSystemsoftwaresRequest,
) (*[]appv1.SystemSoftware, *http.Response, error) {
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

	var response []appv1.SystemSoftware
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//List SystemSoftwareVersions belonging to a SystemSoftware.
func (c *clientImpl) ListSystemsoftwareversions(
	ctx context.Context,
	req ListSystemsoftwareversionsRequest,
) (*[]appv1.SystemSoftwareVersion, *http.Response, error) {
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

	var response []appv1.SystemSoftwareVersion
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//List update candidates belonging to an AppVersion.
func (c *clientImpl) ListUpdateCandidatesForAppversion(
	ctx context.Context,
	req ListUpdateCandidatesForAppversionRequest,
) (*[]appv1.AppVersion, *http.Response, error) {
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

	var response []appv1.AppVersion
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Replace a MySQL Database with another MySQL Database.
func (c *clientImpl) ReplaceDatabase(
	ctx context.Context,
	req ReplaceDatabaseRequest,
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

//Request a copy of an AppInstallation.
func (c *clientImpl) RequestAppinstallationCopy(
	ctx context.Context,
	req RequestAppinstallationCopyRequest,
) (*RequestAppinstallationCopyResponse, *http.Response, error) {
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

	var response RequestAppinstallationCopyResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get runtime status belonging to an AppInstallation.
func (c *clientImpl) RetrieveStatus(
	ctx context.Context,
	req RetrieveStatusRequest,
) (*appv1.AppInstallationStatus, *http.Response, error) {
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

	var response appv1.AppInstallationStatus
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Remove linkage between an AppInstallation and a Database.
func (c *clientImpl) UnlinkDatabase(
	ctx context.Context,
	req UnlinkDatabaseRequest,
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

//Create linkage between an AppInstallation and DatabaseUsers.
func (c *clientImpl) SetDatabaseUsers(
	ctx context.Context,
	req SetDatabaseUsersRequest,
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

//Create linkage between an AppInstallation and a MySql-Database.
//
//This route is deprecated. Use PATCH /v2/app-installations/{appInstallationId}/database instead.
func (c *clientImpl) DeprecatedAppLinkDatabase(
	ctx context.Context,
	req DeprecatedAppLinkDatabaseRequest,
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
