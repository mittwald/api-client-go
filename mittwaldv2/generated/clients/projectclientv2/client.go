package projectclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/membershipv2"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/projectv2"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/storagespacev2"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

type Client interface {
	DeprecatedLeaveProject(
		ctx context.Context,
		req DeprecatedLeaveProjectRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	AcceptProjectInvite(
		ctx context.Context,
		req AcceptProjectInviteRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	ListInvitesForProject(
		ctx context.Context,
		req ListInvitesForProjectRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]membershipv2.ProjectInvite, *http.Response, error)
	CreateProjectInvite(
		ctx context.Context,
		req CreateProjectInviteRequest,
		reqEditors ...func(req *http.Request) error,
	) (*membershipv2.ProjectInvite, *http.Response, error)
	CreateProject(
		ctx context.Context,
		req CreateProjectRequest,
		reqEditors ...func(req *http.Request) error,
	) (*CreateProjectResponse, *http.Response, error)
	DeclineProjectInvite(
		ctx context.Context,
		req DeclineProjectInviteRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	RequestProjectAvatarUpload(
		ctx context.Context,
		req RequestProjectAvatarUploadRequest,
		reqEditors ...func(req *http.Request) error,
	) (*RequestProjectAvatarUploadResponse, *http.Response, error)
	DeleteProjectAvatar(
		ctx context.Context,
		req DeleteProjectAvatarRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	GetProjectInvite(
		ctx context.Context,
		req GetProjectInviteRequest,
		reqEditors ...func(req *http.Request) error,
	) (*membershipv2.ProjectInvite, *http.Response, error)
	DeleteProjectInvite(
		ctx context.Context,
		req DeleteProjectInviteRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	GetProjectMembership(
		ctx context.Context,
		req GetProjectMembershipRequest,
		reqEditors ...func(req *http.Request) error,
	) (*membershipv2.ProjectMembership, *http.Response, error)
	DeleteProjectMembership(
		ctx context.Context,
		req DeleteProjectMembershipRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	UpdateProjectMembership(
		ctx context.Context,
		req UpdateProjectMembershipRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	GetProject(
		ctx context.Context,
		req GetProjectRequest,
		reqEditors ...func(req *http.Request) error,
	) (*projectv2.Project, *http.Response, error)
	DeleteProject(
		ctx context.Context,
		req DeleteProjectRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	RequestServerAvatarUpload(
		ctx context.Context,
		req RequestServerAvatarUploadRequest,
		reqEditors ...func(req *http.Request) error,
	) (*RequestServerAvatarUploadResponse, *http.Response, error)
	DeleteServerAvatar(
		ctx context.Context,
		req DeleteServerAvatarRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	GetProjectTokenInvite(
		ctx context.Context,
		req GetProjectTokenInviteRequest,
		reqEditors ...func(req *http.Request) error,
	) (*membershipv2.ProjectInvite, *http.Response, error)
	GetSelfMembershipForProject(
		ctx context.Context,
		req GetSelfMembershipForProjectRequest,
		reqEditors ...func(req *http.Request) error,
	) (*membershipv2.ProjectMembership, *http.Response, error)
	GetServer(
		ctx context.Context,
		req GetServerRequest,
		reqEditors ...func(req *http.Request) error,
	) (*projectv2.Server, *http.Response, error)
	ListMembershipsForProject(
		ctx context.Context,
		req ListMembershipsForProjectRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]membershipv2.ProjectMembership, *http.Response, error)
	ListProjectInvites(
		ctx context.Context,
		req ListProjectInvitesRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]membershipv2.ProjectInvite, *http.Response, error)
	ListProjectMemberships(
		ctx context.Context,
		req ListProjectMembershipsRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]membershipv2.ProjectMembership, *http.Response, error)
	ListProjects(
		ctx context.Context,
		req ListProjectsRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]ListProjectsResponseItem, *http.Response, error)
	ListServers(
		ctx context.Context,
		req ListServersRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]projectv2.Server, *http.Response, error)
	ResendProjectInviteMail(
		ctx context.Context,
		req ResendProjectInviteMailRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	UpdateProjectDescription(
		ctx context.Context,
		req UpdateProjectDescriptionRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	UpdateServerDescription(
		ctx context.Context,
		req UpdateServerDescriptionRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	StoragespaceGetProjectStatistics(
		ctx context.Context,
		req StoragespaceGetProjectStatisticsRequest,
		reqEditors ...func(req *http.Request) error,
	) (*storagespacev2.Statistics, *http.Response, error)
	StoragespaceGetServerStatistics(
		ctx context.Context,
		req StoragespaceGetServerStatisticsRequest,
		reqEditors ...func(req *http.Request) error,
	) (*storagespacev2.Statistics, *http.Response, error)
	StoragespaceReplaceProjectNotificationThreshold(
		ctx context.Context,
		req StoragespaceReplaceProjectNotificationThresholdRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
	StoragespaceReplaceServerNotificationThreshold(
		ctx context.Context,
		req StoragespaceReplaceServerNotificationThresholdRequest,
		reqEditors ...func(req *http.Request) error,
	) (*http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// Leave a Project.
//
// Deprecated by `DELETE /v2/project-memberships/{projectMembershipId}`.
func (c *clientImpl) DeprecatedLeaveProject(
	ctx context.Context,
	req DeprecatedLeaveProjectRequest,
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

// Accept a ProjectInvite.
func (c *clientImpl) AcceptProjectInvite(
	ctx context.Context,
	req AcceptProjectInviteRequest,
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

// List Invites belonging to a Project.
func (c *clientImpl) ListInvitesForProject(
	ctx context.Context,
	req ListInvitesForProjectRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]membershipv2.ProjectInvite, *http.Response, error) {
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

	var response []membershipv2.ProjectInvite
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create a ProjectInvite.
func (c *clientImpl) CreateProjectInvite(
	ctx context.Context,
	req CreateProjectInviteRequest,
	reqEditors ...func(req *http.Request) error,
) (*membershipv2.ProjectInvite, *http.Response, error) {
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

	var response membershipv2.ProjectInvite
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create a Project belonging to a Server.
func (c *clientImpl) CreateProject(
	ctx context.Context,
	req CreateProjectRequest,
	reqEditors ...func(req *http.Request) error,
) (*CreateProjectResponse, *http.Response, error) {
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

	var response CreateProjectResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Decline a ProjectInvite.
func (c *clientImpl) DeclineProjectInvite(
	ctx context.Context,
	req DeclineProjectInviteRequest,
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

// Request a Project avatar upload.
func (c *clientImpl) RequestProjectAvatarUpload(
	ctx context.Context,
	req RequestProjectAvatarUploadRequest,
	reqEditors ...func(req *http.Request) error,
) (*RequestProjectAvatarUploadResponse, *http.Response, error) {
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

	var response RequestProjectAvatarUploadResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a Project's avatar.
func (c *clientImpl) DeleteProjectAvatar(
	ctx context.Context,
	req DeleteProjectAvatarRequest,
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

// Get a ProjectInvite.
func (c *clientImpl) GetProjectInvite(
	ctx context.Context,
	req GetProjectInviteRequest,
	reqEditors ...func(req *http.Request) error,
) (*membershipv2.ProjectInvite, *http.Response, error) {
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

	var response membershipv2.ProjectInvite
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a ProjectInvite.
func (c *clientImpl) DeleteProjectInvite(
	ctx context.Context,
	req DeleteProjectInviteRequest,
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

// Get a ProjectMembership
func (c *clientImpl) GetProjectMembership(
	ctx context.Context,
	req GetProjectMembershipRequest,
	reqEditors ...func(req *http.Request) error,
) (*membershipv2.ProjectMembership, *http.Response, error) {
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

	var response membershipv2.ProjectMembership
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a ProjectMembership.
func (c *clientImpl) DeleteProjectMembership(
	ctx context.Context,
	req DeleteProjectMembershipRequest,
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

// Update a ProjectMembership.
func (c *clientImpl) UpdateProjectMembership(
	ctx context.Context,
	req UpdateProjectMembershipRequest,
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

// Get a Project.
func (c *clientImpl) GetProject(
	ctx context.Context,
	req GetProjectRequest,
	reqEditors ...func(req *http.Request) error,
) (*projectv2.Project, *http.Response, error) {
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

	var response projectv2.Project
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a Project.
func (c *clientImpl) DeleteProject(
	ctx context.Context,
	req DeleteProjectRequest,
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

// Request a Server avatar upload.
func (c *clientImpl) RequestServerAvatarUpload(
	ctx context.Context,
	req RequestServerAvatarUploadRequest,
	reqEditors ...func(req *http.Request) error,
) (*RequestServerAvatarUploadResponse, *http.Response, error) {
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

	var response RequestServerAvatarUploadResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a Server's avatar.
func (c *clientImpl) DeleteServerAvatar(
	ctx context.Context,
	req DeleteServerAvatarRequest,
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

// Get a ProjectInvite by token.
func (c *clientImpl) GetProjectTokenInvite(
	ctx context.Context,
	req GetProjectTokenInviteRequest,
	reqEditors ...func(req *http.Request) error,
) (*membershipv2.ProjectInvite, *http.Response, error) {
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

	var response membershipv2.ProjectInvite
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get the executing user's membership in a Project.
func (c *clientImpl) GetSelfMembershipForProject(
	ctx context.Context,
	req GetSelfMembershipForProjectRequest,
	reqEditors ...func(req *http.Request) error,
) (*membershipv2.ProjectMembership, *http.Response, error) {
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

	var response membershipv2.ProjectMembership
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a Server.
func (c *clientImpl) GetServer(
	ctx context.Context,
	req GetServerRequest,
	reqEditors ...func(req *http.Request) error,
) (*projectv2.Server, *http.Response, error) {
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

	var response projectv2.Server
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List Memberships belonging to a Project.
func (c *clientImpl) ListMembershipsForProject(
	ctx context.Context,
	req ListMembershipsForProjectRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]membershipv2.ProjectMembership, *http.Response, error) {
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

	var response []membershipv2.ProjectMembership
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List ProjectInvites belonging to the executing user.
func (c *clientImpl) ListProjectInvites(
	ctx context.Context,
	req ListProjectInvitesRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]membershipv2.ProjectInvite, *http.Response, error) {
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

	var response []membershipv2.ProjectInvite
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List ProjectMemberships belonging to the executing user.
func (c *clientImpl) ListProjectMemberships(
	ctx context.Context,
	req ListProjectMembershipsRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]membershipv2.ProjectMembership, *http.Response, error) {
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

	var response []membershipv2.ProjectMembership
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List Projects belonging to the executing user.
func (c *clientImpl) ListProjects(
	ctx context.Context,
	req ListProjectsRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]ListProjectsResponseItem, *http.Response, error) {
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

	var response []ListProjectsResponseItem
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List Servers belonging to the executing user.
func (c *clientImpl) ListServers(
	ctx context.Context,
	req ListServersRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]projectv2.Server, *http.Response, error) {
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

	var response []projectv2.Server
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Resend the mail for a ProjectInvite.
func (c *clientImpl) ResendProjectInviteMail(
	ctx context.Context,
	req ResendProjectInviteMailRequest,
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

// Update a Project's description.
func (c *clientImpl) UpdateProjectDescription(
	ctx context.Context,
	req UpdateProjectDescriptionRequest,
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

// Update a Servers's description.
func (c *clientImpl) UpdateServerDescription(
	ctx context.Context,
	req UpdateServerDescriptionRequest,
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

// Get storage space Statistics belonging to a Project.
func (c *clientImpl) StoragespaceGetProjectStatistics(
	ctx context.Context,
	req StoragespaceGetProjectStatisticsRequest,
	reqEditors ...func(req *http.Request) error,
) (*storagespacev2.Statistics, *http.Response, error) {
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

	var response storagespacev2.Statistics
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get storage space Statistics belonging to a Server.
func (c *clientImpl) StoragespaceGetServerStatistics(
	ctx context.Context,
	req StoragespaceGetServerStatisticsRequest,
	reqEditors ...func(req *http.Request) error,
) (*storagespacev2.Statistics, *http.Response, error) {
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

	var response storagespacev2.Statistics
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Update a Project's storage space notification threshold.
func (c *clientImpl) StoragespaceReplaceProjectNotificationThreshold(
	ctx context.Context,
	req StoragespaceReplaceProjectNotificationThresholdRequest,
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

// Update a Server's storage space notification threshold.
func (c *clientImpl) StoragespaceReplaceServerNotificationThreshold(
	ctx context.Context,
	req StoragespaceReplaceServerNotificationThresholdRequest,
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
