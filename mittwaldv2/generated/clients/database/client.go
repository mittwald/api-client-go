package database

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/databasev1"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

type Client interface {
	ListMysqlDatabases(
		ctx context.Context,
		req ListMysqlDatabasesRequest,
	) (*[]databasev1.MySqlDatabase, *http.Response, error)
	CreateMysqlDatabase(
		ctx context.Context,
		req CreateMysqlDatabaseRequest,
	) (*CreateMysqlDatabaseResponse, *http.Response, error)
	ListMysqlUsers(
		ctx context.Context,
		req ListMysqlUsersRequest,
	) (*[]databasev1.MySqlUser, *http.Response, error)
	CreateMysqlUser(
		ctx context.Context,
		req CreateMysqlUserRequest,
	) (*CreateMysqlUserResponse, *http.Response, error)
	ListRedisDatabases(
		ctx context.Context,
		req ListRedisDatabasesRequest,
	) (*[]databasev1.RedisDatabase, *http.Response, error)
	CreateRedisDatabase(
		ctx context.Context,
		req CreateRedisDatabaseRequest,
	) (*CreateRedisDatabaseResponse, *http.Response, error)
	GetMysqlDatabase(
		ctx context.Context,
		req GetMysqlDatabaseRequest,
	) (*databasev1.MySqlDatabase, *http.Response, error)
	DeleteMysqlDatabase(
		ctx context.Context,
		req DeleteMysqlDatabaseRequest,
	) (*http.Response, error)
	GetMysqlUser(
		ctx context.Context,
		req GetMysqlUserRequest,
	) (*databasev1.MySqlUser, *http.Response, error)
	UpdateMysqlUser(
		ctx context.Context,
		req UpdateMysqlUserRequest,
	) (*http.Response, error)
	DeleteMysqlUser(
		ctx context.Context,
		req DeleteMysqlUserRequest,
	) (*http.Response, error)
	GetRedisDatabase(
		ctx context.Context,
		req GetRedisDatabaseRequest,
	) (*databasev1.RedisDatabase, *http.Response, error)
	DeleteRedisDatabase(
		ctx context.Context,
		req DeleteRedisDatabaseRequest,
	) (*http.Response, error)
	DisableMysqlUser(
		ctx context.Context,
		req DisableMysqlUserRequest,
	) (*http.Response, error)
	EnableMysqlUser(
		ctx context.Context,
		req EnableMysqlUserRequest,
	) (*http.Response, error)
	GetMysqlUserPhpMyAdminURL(
		ctx context.Context,
		req GetMysqlUserPhpMyAdminURLRequest,
	) (*databasev1.PhpMyAdminURL, *http.Response, error)
	ListMysqlCharsets(
		ctx context.Context,
		req ListMysqlCharsetsRequest,
	) (*[]databasev1.MySqlCharacterSettings, *http.Response, error)
	ListMysqlVersions(
		ctx context.Context,
		req ListMysqlVersionsRequest,
	) (*[]databasev1.MySqlVersion, *http.Response, error)
	ListRedisVersions(
		ctx context.Context,
		req ListRedisVersionsRequest,
	) (*[]databasev1.RedisVersion, *http.Response, error)
	UpdateMysqlDatabaseDefaultCharset(
		ctx context.Context,
		req UpdateMysqlDatabaseDefaultCharsetRequest,
	) (*http.Response, error)
	UpdateMysqlDatabaseDescription(
		ctx context.Context,
		req UpdateMysqlDatabaseDescriptionRequest,
	) (*http.Response, error)
	UpdateMysqlUserPassword(
		ctx context.Context,
		req UpdateMysqlUserPasswordRequest,
	) (*http.Response, error)
	UpdateRedisDatabaseConfiguration(
		ctx context.Context,
		req UpdateRedisDatabaseConfigurationRequest,
	) (*http.Response, error)
	UpdateRedisDatabaseDescription(
		ctx context.Context,
		req UpdateRedisDatabaseDescriptionRequest,
	) (*http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// List MySQLDatabases belonging to a Project.
func (c *clientImpl) ListMysqlDatabases(
	ctx context.Context,
	req ListMysqlDatabasesRequest,
) (*[]databasev1.MySqlDatabase, *http.Response, error) {
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

	var response []databasev1.MySqlDatabase
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create a MySQLDatabase with a MySQLUser.
func (c *clientImpl) CreateMysqlDatabase(
	ctx context.Context,
	req CreateMysqlDatabaseRequest,
) (*CreateMysqlDatabaseResponse, *http.Response, error) {
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

	var response CreateMysqlDatabaseResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List MySQLUsers belonging to a Database.
func (c *clientImpl) ListMysqlUsers(
	ctx context.Context,
	req ListMysqlUsersRequest,
) (*[]databasev1.MySqlUser, *http.Response, error) {
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

	var response []databasev1.MySqlUser
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create a MySQLUser.
func (c *clientImpl) CreateMysqlUser(
	ctx context.Context,
	req CreateMysqlUserRequest,
) (*CreateMysqlUserResponse, *http.Response, error) {
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

	var response CreateMysqlUserResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List RedisDatabases belonging to a Project.
func (c *clientImpl) ListRedisDatabases(
	ctx context.Context,
	req ListRedisDatabasesRequest,
) (*[]databasev1.RedisDatabase, *http.Response, error) {
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

	var response []databasev1.RedisDatabase
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create a RedisDatabase.
func (c *clientImpl) CreateRedisDatabase(
	ctx context.Context,
	req CreateRedisDatabaseRequest,
) (*CreateRedisDatabaseResponse, *http.Response, error) {
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

	var response CreateRedisDatabaseResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get a MySQLDatabase.
func (c *clientImpl) GetMysqlDatabase(
	ctx context.Context,
	req GetMysqlDatabaseRequest,
) (*databasev1.MySqlDatabase, *http.Response, error) {
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

	var response databasev1.MySqlDatabase
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a MySQLDatabase.
func (c *clientImpl) DeleteMysqlDatabase(
	ctx context.Context,
	req DeleteMysqlDatabaseRequest,
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

// Get a MySQLUser.
func (c *clientImpl) GetMysqlUser(
	ctx context.Context,
	req GetMysqlUserRequest,
) (*databasev1.MySqlUser, *http.Response, error) {
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

	var response databasev1.MySqlUser
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Update a MySQLUser.
func (c *clientImpl) UpdateMysqlUser(
	ctx context.Context,
	req UpdateMysqlUserRequest,
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

// Delete a MySQLUser.
func (c *clientImpl) DeleteMysqlUser(
	ctx context.Context,
	req DeleteMysqlUserRequest,
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

// Get a RedisDatabase.
func (c *clientImpl) GetRedisDatabase(
	ctx context.Context,
	req GetRedisDatabaseRequest,
) (*databasev1.RedisDatabase, *http.Response, error) {
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

	var response databasev1.RedisDatabase
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Delete a RedisDatabase.
func (c *clientImpl) DeleteRedisDatabase(
	ctx context.Context,
	req DeleteRedisDatabaseRequest,
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

// Disable a MySQLUser.
func (c *clientImpl) DisableMysqlUser(
	ctx context.Context,
	req DisableMysqlUserRequest,
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

// Enable a MySQLUser.
func (c *clientImpl) EnableMysqlUser(
	ctx context.Context,
	req EnableMysqlUserRequest,
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

// Get a MySQLUser's PhpMyAdmin-URL.
func (c *clientImpl) GetMysqlUserPhpMyAdminURL(
	ctx context.Context,
	req GetMysqlUserPhpMyAdminURLRequest,
) (*databasev1.PhpMyAdminURL, *http.Response, error) {
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

	var response databasev1.PhpMyAdminURL
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List available MySQL character sets and collations, optionally filtered by a MySQLVersion.
func (c *clientImpl) ListMysqlCharsets(
	ctx context.Context,
	req ListMysqlCharsetsRequest,
) (*[]databasev1.MySqlCharacterSettings, *http.Response, error) {
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

	var response []databasev1.MySqlCharacterSettings
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List MySQLVersions.
func (c *clientImpl) ListMysqlVersions(
	ctx context.Context,
	req ListMysqlVersionsRequest,
) (*[]databasev1.MySqlVersion, *http.Response, error) {
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

	var response []databasev1.MySqlVersion
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List RedisVersions.
func (c *clientImpl) ListRedisVersions(
	ctx context.Context,
	req ListRedisVersionsRequest,
) (*[]databasev1.RedisVersion, *http.Response, error) {
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

	var response []databasev1.RedisVersion
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Update a MySQLDatabase's default character settings.
func (c *clientImpl) UpdateMysqlDatabaseDefaultCharset(
	ctx context.Context,
	req UpdateMysqlDatabaseDefaultCharsetRequest,
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

// Update a MySQLDatabase's description.
func (c *clientImpl) UpdateMysqlDatabaseDescription(
	ctx context.Context,
	req UpdateMysqlDatabaseDescriptionRequest,
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

// Update a MySQLUser's password.
func (c *clientImpl) UpdateMysqlUserPassword(
	ctx context.Context,
	req UpdateMysqlUserPasswordRequest,
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

// Update a RedisDatabase's configuration.
func (c *clientImpl) UpdateRedisDatabaseConfiguration(
	ctx context.Context,
	req UpdateRedisDatabaseConfigurationRequest,
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

// Update a RedisDatabase's description.
func (c *clientImpl) UpdateRedisDatabaseDescription(
	ctx context.Context,
	req UpdateRedisDatabaseDescriptionRequest,
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
