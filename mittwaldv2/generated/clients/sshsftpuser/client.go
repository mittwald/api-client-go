package sshsftpuser

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/sshuserv1"
	"github.com/mittwald/api-client-go/mittwaldv2/httpclient"
	"github.com/mittwald/api-client-go/mittwaldv2/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	ListSFTPUsers(
		ctx context.Context,
		req ListSFTPUsersRequest,
	) (*[]sshuserv1.SftpUser, *http.Response, error)
	CreateSFTPUser(
		ctx context.Context,
		req CreateSFTPUserRequest,
	) (*sshuserv1.SftpUser, *http.Response, error)
	GetSFTPUser(
		ctx context.Context,
		req GetSFTPUserRequest,
	) (*sshuserv1.SftpUser, *http.Response, error)
	DeleteSFTPUser(
		ctx context.Context,
		req DeleteSFTPUserRequest,
	) (*http.Response, error)
	UpdateSFTPUser(
		ctx context.Context,
		req UpdateSFTPUserRequest,
	) (*http.Response, error)
	ListSSHUsers(
		ctx context.Context,
		req ListSSHUsersRequest,
	) (*[]sshuserv1.SshUser, *http.Response, error)
	CreateSSHUser(
		ctx context.Context,
		req CreateSSHUserRequest,
	) (*sshuserv1.SshUser, *http.Response, error)
	GetSSHUser(
		ctx context.Context,
		req GetSSHUserRequest,
	) (*sshuserv1.SshUser, *http.Response, error)
	DeleteSSHUser(
		ctx context.Context,
		req DeleteSSHUserRequest,
	) (*http.Response, error)
	UpdateSSHUser(
		ctx context.Context,
		req UpdateSSHUserRequest,
	) (*http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

//Get all SFTPUsers for a Project.
func (c *clientImpl) ListSFTPUsers(
	ctx context.Context,
	req ListSFTPUsersRequest,
) (*[]sshuserv1.SftpUser, *http.Response, error) {
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

	var response []sshuserv1.SftpUser
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Create an SFTPUser for a Project.
func (c *clientImpl) CreateSFTPUser(
	ctx context.Context,
	req CreateSFTPUserRequest,
) (*sshuserv1.SftpUser, *http.Response, error) {
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

	var response sshuserv1.SftpUser
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get an SFTPUser.
func (c *clientImpl) GetSFTPUser(
	ctx context.Context,
	req GetSFTPUserRequest,
) (*sshuserv1.SftpUser, *http.Response, error) {
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

	var response sshuserv1.SftpUser
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Delete an SFTPUser.
func (c *clientImpl) DeleteSFTPUser(
	ctx context.Context,
	req DeleteSFTPUserRequest,
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

//Update an SFTPUser.
func (c *clientImpl) UpdateSFTPUser(
	ctx context.Context,
	req UpdateSFTPUserRequest,
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

//Get all SSHUsers for a Project.
func (c *clientImpl) ListSSHUsers(
	ctx context.Context,
	req ListSSHUsersRequest,
) (*[]sshuserv1.SshUser, *http.Response, error) {
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

	var response []sshuserv1.SshUser
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Create an SSHUser for a Project.
func (c *clientImpl) CreateSSHUser(
	ctx context.Context,
	req CreateSSHUserRequest,
) (*sshuserv1.SshUser, *http.Response, error) {
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

	var response sshuserv1.SshUser
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get an SSHUser.
func (c *clientImpl) GetSSHUser(
	ctx context.Context,
	req GetSSHUserRequest,
) (*sshuserv1.SshUser, *http.Response, error) {
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

	var response sshuserv1.SshUser
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Delete an SSHUser.
func (c *clientImpl) DeleteSSHUser(
	ctx context.Context,
	req DeleteSSHUserRequest,
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

//Update an SSHUser.
func (c *clientImpl) UpdateSSHUser(
	ctx context.Context,
	req UpdateSSHUserRequest,
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
