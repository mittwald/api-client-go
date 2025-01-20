package conversation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/conversationv1"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	ListConversations(
		ctx context.Context,
		req ListConversationsRequest,
	) (*[]conversationv1.Conversation, *http.Response, error)
	CreateConversation(
		ctx context.Context,
		req CreateConversationRequest,
	) (*CreateConversationResponse, *http.Response, error)
	ListMessagesByConversation(
		ctx context.Context,
		req ListMessagesByConversationRequest,
	) (*[]ListMessagesByConversationResponseItem, *http.Response, error)
	CreateMessage(
		ctx context.Context,
		req CreateMessageRequest,
	) (*CreateMessageResponse, *http.Response, error)
	GetCategory(
		ctx context.Context,
		req GetCategoryRequest,
	) (*conversationv1.Category, *http.Response, error)
	GetConversationMembers(
		ctx context.Context,
		req GetConversationMembersRequest,
	) (*conversationv1.ConversationMembers, *http.Response, error)
	GetConversationPreferencesOfCustomer(
		ctx context.Context,
		req GetConversationPreferencesOfCustomerRequest,
	) (*conversationv1.ConversationPreferences, *http.Response, error)
	GetConversation(
		ctx context.Context,
		req GetConversationRequest,
	) (*conversationv1.Conversation, *http.Response, error)
	UpdateConversation(
		ctx context.Context,
		req UpdateConversationRequest,
	) (*UpdateConversationResponse, *http.Response, error)
	GetFileAccessToken(
		ctx context.Context,
		req GetFileAccessTokenRequest,
	) (*GetFileAccessTokenResponse, *http.Response, error)
	ListCategories(
		ctx context.Context,
		req ListCategoriesRequest,
	) (*[]conversationv1.Category, *http.Response, error)
	RequestFileUpload(
		ctx context.Context,
		req RequestFileUploadRequest,
	) (*RequestFileUploadResponse, *http.Response, error)
	SetConversationStatus(
		ctx context.Context,
		req SetConversationStatusRequest,
	) (*SetConversationStatusResponse, *http.Response, error)
	UpdateMessage(
		ctx context.Context,
		req UpdateMessageRequest,
	) (*UpdateMessageResponse, *http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

//Get all conversation the authenticated user has created or has access to.
func (c *clientImpl) ListConversations(
	ctx context.Context,
	req ListConversationsRequest,
) (*[]conversationv1.Conversation, *http.Response, error) {
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

	var response []conversationv1.Conversation
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Create a conversation.
func (c *clientImpl) CreateConversation(
	ctx context.Context,
	req CreateConversationRequest,
) (*CreateConversationResponse, *http.Response, error) {
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

	var response CreateConversationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get all message of the conversation.
func (c *clientImpl) ListMessagesByConversation(
	ctx context.Context,
	req ListMessagesByConversationRequest,
) (*[]ListMessagesByConversationResponseItem, *http.Response, error) {
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

	var response []ListMessagesByConversationResponseItem
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Send a new message in the conversation.
func (c *clientImpl) CreateMessage(
	ctx context.Context,
	req CreateMessageRequest,
) (*CreateMessageResponse, *http.Response, error) {
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

	var response CreateMessageResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a specific conversation category.
func (c *clientImpl) GetCategory(
	ctx context.Context,
	req GetCategoryRequest,
) (*conversationv1.Category, *http.Response, error) {
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

	var response conversationv1.Category
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get members of a support conversation.
func (c *clientImpl) GetConversationMembers(
	ctx context.Context,
	req GetConversationMembersRequest,
) (*conversationv1.ConversationMembers, *http.Response, error) {
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

	var response conversationv1.ConversationMembers
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get preferences for customer conversations.
func (c *clientImpl) GetConversationPreferencesOfCustomer(
	ctx context.Context,
	req GetConversationPreferencesOfCustomerRequest,
) (*conversationv1.ConversationPreferences, *http.Response, error) {
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

	var response conversationv1.ConversationPreferences
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get a support conversation.
func (c *clientImpl) GetConversation(
	ctx context.Context,
	req GetConversationRequest,
) (*conversationv1.Conversation, *http.Response, error) {
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

	var response conversationv1.Conversation
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Update the basic properties of the conversation.
func (c *clientImpl) UpdateConversation(
	ctx context.Context,
	req UpdateConversationRequest,
) (*UpdateConversationResponse, *http.Response, error) {
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

	var response UpdateConversationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Request an access token for the File belonging to the Conversation.
func (c *clientImpl) GetFileAccessToken(
	ctx context.Context,
	req GetFileAccessTokenRequest,
) (*GetFileAccessTokenResponse, *http.Response, error) {
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

	var response GetFileAccessTokenResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Get all conversation categories.
func (c *clientImpl) ListCategories(
	ctx context.Context,
	req ListCategoriesRequest,
) (*[]conversationv1.Category, *http.Response, error) {
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

	var response []conversationv1.Category
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Request a file upload token for the conversation.
func (c *clientImpl) RequestFileUpload(
	ctx context.Context,
	req RequestFileUploadRequest,
) (*RequestFileUploadResponse, *http.Response, error) {
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

	var response RequestFileUploadResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Update the status of a conversation.
func (c *clientImpl) SetConversationStatus(
	ctx context.Context,
	req SetConversationStatusRequest,
) (*SetConversationStatusResponse, *http.Response, error) {
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

	var response SetConversationStatusResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

//Update the content of the message
func (c *clientImpl) UpdateMessage(
	ctx context.Context,
	req UpdateMessageRequest,
) (*UpdateMessageResponse, *http.Response, error) {
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

	var response UpdateMessageResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}
