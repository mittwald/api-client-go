package projectfilesystem

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/projectv1"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	ProjectFileSystemGetDirectories(
		ctx context.Context,
		req ProjectFileSystemGetDirectoriesRequest,
	) (*projectv1.FilesystemDirectoryListing, *http.Response, error)
	ProjectFileSystemGetDiskUsage(
		ctx context.Context,
		req ProjectFileSystemGetDiskUsageRequest,
	) (*projectv1.FilesystemUsagesDisk, *http.Response, error)
	ProjectFileSystemGetFileContent(
		ctx context.Context,
		req ProjectFileSystemGetFileContentRequest,
	) (*http.Response, error)
	ProjectFileSystemGetJwt(
		ctx context.Context,
		req ProjectFileSystemGetJwtRequest,
	) (*projectv1.FsApiJwt, *http.Response, error)
	ProjectFileSystemListFiles(
		ctx context.Context,
		req ProjectFileSystemListFilesRequest,
	) (*projectv1.FilesystemDirectoryListing, *http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// List directories belonging to a Project.
func (c *clientImpl) ProjectFileSystemGetDirectories(
	ctx context.Context,
	req ProjectFileSystemGetDirectoriesRequest,
) (*projectv1.FilesystemDirectoryListing, *http.Response, error) {
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

	var response projectv1.FilesystemDirectoryListing
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a Project directory filesystem usage.
func (c *clientImpl) ProjectFileSystemGetDiskUsage(
	ctx context.Context,
	req ProjectFileSystemGetDiskUsageRequest,
) (*projectv1.FilesystemUsagesDisk, *http.Response, error) {
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

	var response projectv1.FilesystemUsagesDisk
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a Project file's content.
func (c *clientImpl) ProjectFileSystemGetFileContent(
	ctx context.Context,
	req ProjectFileSystemGetFileContentRequest,
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

// Get a Project's file/filesystem authorization token.
func (c *clientImpl) ProjectFileSystemGetJwt(
	ctx context.Context,
	req ProjectFileSystemGetJwtRequest,
) (*projectv1.FsApiJwt, *http.Response, error) {
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

	var response projectv1.FsApiJwt
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a Project file's information.
func (c *clientImpl) ProjectFileSystemListFiles(
	ctx context.Context,
	req ProjectFileSystemListFilesRequest,
) (*projectv1.FilesystemDirectoryListing, *http.Response, error) {
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

	var response projectv1.FilesystemDirectoryListing
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}
