package conversationclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/conversationv2"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

type Client interface {
	ListConversations(
		ctx context.Context,
		req ListConversationsRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]conversationv2.Conversation, *http.Response, error)
	CreateConversation(
		ctx context.Context,
		req CreateConversationRequest,
		reqEditors ...func(req *http.Request) error,
	) (*CreateConversationResponse, *http.Response, error)
	ListMessagesByConversation(
		ctx context.Context,
		req ListMessagesByConversationRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]ListMessagesByConversationResponseItem, *http.Response, error)
	CreateMessage(
		ctx context.Context,
		req CreateMessageRequest,
		reqEditors ...func(req *http.Request) error,
	) (*CreateMessageResponse, *http.Response, error)
	GetCategory(
		ctx context.Context,
		req GetCategoryRequest,
		reqEditors ...func(req *http.Request) error,
	) (*conversationv2.Category, *http.Response, error)
	GetConversationMembers(
		ctx context.Context,
		req GetConversationMembersRequest,
		reqEditors ...func(req *http.Request) error,
	) (*conversationv2.ConversationMembers, *http.Response, error)
	GetConversationPreferencesOfCustomer(
		ctx context.Context,
		req GetConversationPreferencesOfCustomerRequest,
		reqEditors ...func(req *http.Request) error,
	) (*conversationv2.ConversationPreferences, *http.Response, error)
	GetConversation(
		ctx context.Context,
		req GetConversationRequest,
		reqEditors ...func(req *http.Request) error,
	) (*conversationv2.Conversation, *http.Response, error)
	UpdateConversation(
		ctx context.Context,
		req UpdateConversationRequest,
		reqEditors ...func(req *http.Request) error,
	) (*UpdateConversationResponse, *http.Response, error)
	GetFileAccessToken(
		ctx context.Context,
		req GetFileAccessTokenRequest,
		reqEditors ...func(req *http.Request) error,
	) (*GetFileAccessTokenResponse, *http.Response, error)
	ListCategories(
		ctx context.Context,
		req ListCategoriesRequest,
		reqEditors ...func(req *http.Request) error,
	) (*[]conversationv2.Category, *http.Response, error)
	RequestFileUpload(
		ctx context.Context,
		req RequestFileUploadRequest,
		reqEditors ...func(req *http.Request) error,
	) (*RequestFileUploadResponse, *http.Response, error)
	SetConversationStatus(
		ctx context.Context,
		req SetConversationStatusRequest,
		reqEditors ...func(req *http.Request) error,
	) (*SetConversationStatusResponse, *http.Response, error)
	UpdateMessage(
		ctx context.Context,
		req UpdateMessageRequest,
		reqEditors ...func(req *http.Request) error,
	) (*UpdateMessageResponse, *http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// Get all conversation the authenticated user has created or has access to.
func (c *clientImpl) ListConversations(
	ctx context.Context,
	req ListConversationsRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]conversationv2.Conversation, *http.Response, error) {
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

	var response []conversationv2.Conversation
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create a conversation.
func (c *clientImpl) CreateConversation(
	ctx context.Context,
	req CreateConversationRequest,
	reqEditors ...func(req *http.Request) error,
) (*CreateConversationResponse, *http.Response, error) {
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

	var response CreateConversationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get all message of the conversation.
func (c *clientImpl) ListMessagesByConversation(
	ctx context.Context,
	req ListMessagesByConversationRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]ListMessagesByConversationResponseItem, *http.Response, error) {
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

	var response []ListMessagesByConversationResponseItem
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Send a new message in the conversation.
func (c *clientImpl) CreateMessage(
	ctx context.Context,
	req CreateMessageRequest,
	reqEditors ...func(req *http.Request) error,
) (*CreateMessageResponse, *http.Response, error) {
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

	var response CreateMessageResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a specific conversation category.
func (c *clientImpl) GetCategory(
	ctx context.Context,
	req GetCategoryRequest,
	reqEditors ...func(req *http.Request) error,
) (*conversationv2.Category, *http.Response, error) {
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

	var response conversationv2.Category
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get members of a support conversation.
func (c *clientImpl) GetConversationMembers(
	ctx context.Context,
	req GetConversationMembersRequest,
	reqEditors ...func(req *http.Request) error,
) (*conversationv2.ConversationMembers, *http.Response, error) {
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

	var response conversationv2.ConversationMembers
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get preferences for customer conversations.
func (c *clientImpl) GetConversationPreferencesOfCustomer(
	ctx context.Context,
	req GetConversationPreferencesOfCustomerRequest,
	reqEditors ...func(req *http.Request) error,
) (*conversationv2.ConversationPreferences, *http.Response, error) {
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

	var response conversationv2.ConversationPreferences
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a support conversation.
func (c *clientImpl) GetConversation(
	ctx context.Context,
	req GetConversationRequest,
	reqEditors ...func(req *http.Request) error,
) (*conversationv2.Conversation, *http.Response, error) {
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

	var response conversationv2.Conversation
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Update the basic properties of the conversation.
func (c *clientImpl) UpdateConversation(
	ctx context.Context,
	req UpdateConversationRequest,
	reqEditors ...func(req *http.Request) error,
) (*UpdateConversationResponse, *http.Response, error) {
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

	var response UpdateConversationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Request an access token for the File belonging to the Conversation.
func (c *clientImpl) GetFileAccessToken(
	ctx context.Context,
	req GetFileAccessTokenRequest,
	reqEditors ...func(req *http.Request) error,
) (*GetFileAccessTokenResponse, *http.Response, error) {
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

	var response GetFileAccessTokenResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get all conversation categories.
func (c *clientImpl) ListCategories(
	ctx context.Context,
	req ListCategoriesRequest,
	reqEditors ...func(req *http.Request) error,
) (*[]conversationv2.Category, *http.Response, error) {
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

	var response []conversationv2.Category
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Request a file upload token for the conversation.
func (c *clientImpl) RequestFileUpload(
	ctx context.Context,
	req RequestFileUploadRequest,
	reqEditors ...func(req *http.Request) error,
) (*RequestFileUploadResponse, *http.Response, error) {
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

	var response RequestFileUploadResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Update the status of a conversation.
func (c *clientImpl) SetConversationStatus(
	ctx context.Context,
	req SetConversationStatusRequest,
	reqEditors ...func(req *http.Request) error,
) (*SetConversationStatusResponse, *http.Response, error) {
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

	var response SetConversationStatusResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Update the content of the message
func (c *clientImpl) UpdateMessage(
	ctx context.Context,
	req UpdateMessageRequest,
	reqEditors ...func(req *http.Request) error,
) (*UpdateMessageResponse, *http.Response, error) {
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

	var response UpdateMessageResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}
