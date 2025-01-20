package file

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/filev1"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	DeprecatedFileGetFileTokenRules(
		ctx context.Context,
		req DeprecatedFileGetFileTokenRulesRequest,
	) (*filev1.FileUploadRules, *http.Response, error)
	DeprecatedFileGetFileTypeRules(
		ctx context.Context,
		req DeprecatedFileGetFileTypeRulesRequest,
	) (*filev1.FileUploadRules, *http.Response, error)
	CreateFile(
		ctx context.Context,
		req CreateFileRequest,
	) (*filev1.FileMeta, *http.Response, error)
	GetFileMeta(
		ctx context.Context,
		req GetFileMetaRequest,
	) (*filev1.FileMeta, *http.Response, error)
	GetFileUploadTokenRules(
		ctx context.Context,
		req GetFileUploadTokenRulesRequest,
	) (*filev1.FileUploadRules, *http.Response, error)
	GetFileUploadTypeRules(
		ctx context.Context,
		req GetFileUploadTypeRulesRequest,
	) (*filev1.FileUploadRules, *http.Response, error)
	GetFile(
		ctx context.Context,
		req GetFileRequest,
	) (*http.Response, error)
	GetFileWithName(
		ctx context.Context,
		req GetFileWithNameRequest,
	) (*http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

//Get a Token's upload rules.
//
//Deprecated by `GET /v2/file-upload-tokens/{fileUploadToken}/rules`.
func (c *clientImpl) DeprecatedFileGetFileTokenRules(
	ctx context.Context,
	req DeprecatedFileGetFileTokenRulesRequest,
) (*filev1.FileUploadRules, *http.Response, error) {
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

	var response filev1.FileUploadRules
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a Type's upload rules.
//
//Deprecated by `GET /v2/file-upload-types/{fileUploadType}/rules`.
func (c *clientImpl) DeprecatedFileGetFileTypeRules(
	ctx context.Context,
	req DeprecatedFileGetFileTypeRulesRequest,
) (*filev1.FileUploadRules, *http.Response, error) {
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

	var response filev1.FileUploadRules
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Create a File.
func (c *clientImpl) CreateFile(
	ctx context.Context,
	req CreateFileRequest,
) (*filev1.FileMeta, *http.Response, error) {
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

	var response filev1.FileMeta
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a File's meta.
func (c *clientImpl) GetFileMeta(
	ctx context.Context,
	req GetFileMetaRequest,
) (*filev1.FileMeta, *http.Response, error) {
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

	var response filev1.FileMeta
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a FileUploadToken's rules.
func (c *clientImpl) GetFileUploadTokenRules(
	ctx context.Context,
	req GetFileUploadTokenRulesRequest,
) (*filev1.FileUploadRules, *http.Response, error) {
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

	var response filev1.FileUploadRules
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a FileUploadType's rules.
func (c *clientImpl) GetFileUploadTypeRules(
	ctx context.Context,
	req GetFileUploadTypeRulesRequest,
) (*filev1.FileUploadRules, *http.Response, error) {
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

	var response filev1.FileUploadRules
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a File.
func (c *clientImpl) GetFile(
	ctx context.Context,
	req GetFileRequest,
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

//Get a File with user-friendly url.
func (c *clientImpl) GetFileWithName(
	ctx context.Context,
	req GetFileWithNameRequest,
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