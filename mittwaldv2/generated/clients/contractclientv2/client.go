package contractclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/contractv2"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/invoicev2"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv2"
	"github.com/mittwald/api-client-go/pkg/httpclient"
	"github.com/mittwald/api-client-go/pkg/httperr"
)

type Client interface {
	TerminateContractItem(
		ctx context.Context,
		req TerminateContractItemRequest,
	) (*TerminateContractItemResponse, *http.Response, error)
	CancelContractItemTermination(
		ctx context.Context,
		req CancelContractItemTerminationRequest,
	) (*CancelContractItemTerminationResponse, *http.Response, error)
	CancelContractTariffChange(
		ctx context.Context,
		req CancelContractTariffChangeRequest,
	) (*CancelContractTariffChangeResponse, *http.Response, error)
	TerminateContract(
		ctx context.Context,
		req TerminateContractRequest,
	) (*TerminateContractResponse, *http.Response, error)
	CancelContractTermination(
		ctx context.Context,
		req CancelContractTerminationRequest,
	) (*CancelContractTerminationResponse, *http.Response, error)
	GetBaseItemOfContract(
		ctx context.Context,
		req GetBaseItemOfContractRequest,
	) (*contractv2.ContractItem, *http.Response, error)
	GetDetailOfContractByCertificate(
		ctx context.Context,
		req GetDetailOfContractByCertificateRequest,
	) (*contractv2.Contract, *http.Response, error)
	GetDetailOfContractByDomain(
		ctx context.Context,
		req GetDetailOfContractByDomainRequest,
	) (*contractv2.Contract, *http.Response, error)
	GetDetailOfContractByProject(
		ctx context.Context,
		req GetDetailOfContractByProjectRequest,
	) (*contractv2.Contract, *http.Response, error)
	GetDetailOfContractByServer(
		ctx context.Context,
		req GetDetailOfContractByServerRequest,
	) (*contractv2.Contract, *http.Response, error)
	GetDetailOfContractItem(
		ctx context.Context,
		req GetDetailOfContractItemRequest,
	) (*contractv2.ContractItem, *http.Response, error)
	GetDetailOfContract(
		ctx context.Context,
		req GetDetailOfContractRequest,
	) (*contractv2.Contract, *http.Response, error)
	GetNextTerminationDateForItem(
		ctx context.Context,
		req GetNextTerminationDateForItemRequest,
	) (*GetNextTerminationDateForItemResponse, *http.Response, error)
	ListContracts(
		ctx context.Context,
		req ListContractsRequest,
	) (*[]contractv2.Contract, *http.Response, error)
	DeprecatedInvoiceDetailOfInvoice(
		ctx context.Context,
		req DeprecatedInvoiceDetailOfInvoiceRequest,
	) (*invoicev2.Invoice, *http.Response, error)
	InvoiceDetail(
		ctx context.Context,
		req InvoiceDetailRequest,
	) (*invoicev2.Invoice, *http.Response, error)
	InvoiceGetDetailOfInvoiceSettings(
		ctx context.Context,
		req InvoiceGetDetailOfInvoiceSettingsRequest,
	) (*invoicev2.InvoiceSettings, *http.Response, error)
	InvoiceUpdateInvoiceSettings(
		ctx context.Context,
		req InvoiceUpdateInvoiceSettingsRequest,
	) (*invoicev2.InvoiceSettings, *http.Response, error)
	InvoiceGetFileAccessToken(
		ctx context.Context,
		req InvoiceGetFileAccessTokenRequest,
	) (*InvoiceGetFileAccessTokenResponse, *http.Response, error)
	InvoiceListCustomerInvoices(
		ctx context.Context,
		req InvoiceListCustomerInvoicesRequest,
	) (*[]invoicev2.Invoice, *http.Response, error)
	ListOrders(
		ctx context.Context,
		req ListOrdersRequest,
	) (*[]orderv2.CustomerOrder, *http.Response, error)
	CreateOrder(
		ctx context.Context,
		req CreateOrderRequest,
	) (*CreateOrderResponse, *http.Response, error)
	CreateTariffChange(
		ctx context.Context,
		req CreateTariffChangeRequest,
	) (*CreateTariffChangeResponse, *http.Response, error)
	GetOrder(
		ctx context.Context,
		req GetOrderRequest,
	) (*orderv2.CustomerOrder, *http.Response, error)
	ListCustomerOrders(
		ctx context.Context,
		req ListCustomerOrdersRequest,
	) (*[]orderv2.CustomerOrder, *http.Response, error)
	ListProjectOrders(
		ctx context.Context,
		req ListProjectOrdersRequest,
	) (*[]orderv2.CustomerOrder, *http.Response, error)
	PreviewOrder(
		ctx context.Context,
		req PreviewOrderRequest,
	) (*PreviewOrderResponse, *http.Response, error)
	PreviewTariffChange(
		ctx context.Context,
		req PreviewTariffChangeRequest,
	) (*PreviewTariffChangeResponse, *http.Response, error)
}
type clientImpl struct {
	client httpclient.RequestRunner
}

func NewClient(client httpclient.RequestRunner) Client {
	return &clientImpl{client: client}
}

// Schedule the Termination of a ContractItem.
func (c *clientImpl) TerminateContractItem(
	ctx context.Context,
	req TerminateContractItemRequest,
) (*TerminateContractItemResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response TerminateContractItemResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Cancel the Termination for the referred ContractItem.
func (c *clientImpl) CancelContractItemTermination(
	ctx context.Context,
	req CancelContractItemTerminationRequest,
) (*CancelContractItemTerminationResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response CancelContractItemTerminationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Cancel the TariffChange for the referred ContractItem.
func (c *clientImpl) CancelContractTariffChange(
	ctx context.Context,
	req CancelContractTariffChangeRequest,
) (*CancelContractTariffChangeResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response CancelContractTariffChangeResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Schedule the Termination of a Contract.
func (c *clientImpl) TerminateContract(
	ctx context.Context,
	req TerminateContractRequest,
) (*TerminateContractResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response TerminateContractResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Cancel the Termination for the referred Contract.
func (c *clientImpl) CancelContractTermination(
	ctx context.Context,
	req CancelContractTerminationRequest,
) (*CancelContractTerminationResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response CancelContractTerminationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Return the BaseItem of the Contract with the given ID.
func (c *clientImpl) GetBaseItemOfContract(
	ctx context.Context,
	req GetBaseItemOfContractRequest,
) (*contractv2.ContractItem, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response contractv2.ContractItem
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Return the Contract for the given Certificate.
func (c *clientImpl) GetDetailOfContractByCertificate(
	ctx context.Context,
	req GetDetailOfContractByCertificateRequest,
) (*contractv2.Contract, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response contractv2.Contract
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Return the Contract for the given Domain.
func (c *clientImpl) GetDetailOfContractByDomain(
	ctx context.Context,
	req GetDetailOfContractByDomainRequest,
) (*contractv2.Contract, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response contractv2.Contract
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Return the Contract for the given Project.
func (c *clientImpl) GetDetailOfContractByProject(
	ctx context.Context,
	req GetDetailOfContractByProjectRequest,
) (*contractv2.Contract, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response contractv2.Contract
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Return the Contract for the given Server.
func (c *clientImpl) GetDetailOfContractByServer(
	ctx context.Context,
	req GetDetailOfContractByServerRequest,
) (*contractv2.Contract, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response contractv2.Contract
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get the ContractItem with the given ID.
func (c *clientImpl) GetDetailOfContractItem(
	ctx context.Context,
	req GetDetailOfContractItemRequest,
) (*contractv2.ContractItem, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response contractv2.ContractItem
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Returns the Contract with the given ID.
func (c *clientImpl) GetDetailOfContract(
	ctx context.Context,
	req GetDetailOfContractRequest,
) (*contractv2.Contract, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response contractv2.Contract
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Return the next TerminationDate for the ContractItem with the given ID.
func (c *clientImpl) GetNextTerminationDateForItem(
	ctx context.Context,
	req GetNextTerminationDateForItemRequest,
) (*GetNextTerminationDateForItemResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response GetNextTerminationDateForItemResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Return a list of Contracts for the given Customer.
func (c *clientImpl) ListContracts(
	ctx context.Context,
	req ListContractsRequest,
) (*[]contractv2.Contract, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []contractv2.Contract
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get details of an Invoice.
//
// This route is deprecated. Use /v2/invoices/{invoiceId} instead.
func (c *clientImpl) DeprecatedInvoiceDetailOfInvoice(
	ctx context.Context,
	req DeprecatedInvoiceDetailOfInvoiceRequest,
) (*invoicev2.Invoice, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response invoicev2.Invoice
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get details of an Invoice.
func (c *clientImpl) InvoiceDetail(
	ctx context.Context,
	req InvoiceDetailRequest,
) (*invoicev2.Invoice, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response invoicev2.Invoice
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get InvoiceSettings of a Customer.
func (c *clientImpl) InvoiceGetDetailOfInvoiceSettings(
	ctx context.Context,
	req InvoiceGetDetailOfInvoiceSettingsRequest,
) (*invoicev2.InvoiceSettings, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response invoicev2.InvoiceSettings
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Update InvoiceSettings of a Customer.
func (c *clientImpl) InvoiceUpdateInvoiceSettings(
	ctx context.Context,
	req InvoiceUpdateInvoiceSettingsRequest,
) (*invoicev2.InvoiceSettings, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response invoicev2.InvoiceSettings
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Request an Access Token for the Invoice file.
func (c *clientImpl) InvoiceGetFileAccessToken(
	ctx context.Context,
	req InvoiceGetFileAccessTokenRequest,
) (*InvoiceGetFileAccessTokenResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response InvoiceGetFileAccessTokenResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// List Invoices of a Customer.
func (c *clientImpl) InvoiceListCustomerInvoices(
	ctx context.Context,
	req InvoiceListCustomerInvoicesRequest,
) (*[]invoicev2.Invoice, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []invoicev2.Invoice
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get list of Orders.
//
// The list of Orders the User has access to.
func (c *clientImpl) ListOrders(
	ctx context.Context,
	req ListOrdersRequest,
) (*[]orderv2.CustomerOrder, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []orderv2.CustomerOrder
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create an Order.
func (c *clientImpl) CreateOrder(
	ctx context.Context,
	req CreateOrderRequest,
) (*CreateOrderResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response CreateOrderResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Create TariffChange Order.
func (c *clientImpl) CreateTariffChange(
	ctx context.Context,
	req CreateTariffChangeRequest,
) (*CreateTariffChangeResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response CreateTariffChangeResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get Order for Customer.
//
// Get details of a single Order, User must have access to the Order/Customer.
func (c *clientImpl) GetOrder(
	ctx context.Context,
	req GetOrderRequest,
) (*orderv2.CustomerOrder, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response orderv2.CustomerOrder
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get list of Orders of a Customer.
//
// The list of Orders of a Customer the User has access to, can be filtered by orderStatus, articleTemplate and count.
func (c *clientImpl) ListCustomerOrders(
	ctx context.Context,
	req ListCustomerOrdersRequest,
) (*[]orderv2.CustomerOrder, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []orderv2.CustomerOrder
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Get list of Orders of a Project.
//
// The list of Order of a Project the User has access to, can be filtered by orderStatus, articleTemplate and count.
func (c *clientImpl) ListProjectOrders(
	ctx context.Context,
	req ListProjectOrdersRequest,
) (*[]orderv2.CustomerOrder, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response []orderv2.CustomerOrder
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Preview Order.
func (c *clientImpl) PreviewOrder(
	ctx context.Context,
	req PreviewOrderRequest,
) (*PreviewOrderResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response PreviewOrderResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}

// Preview TariffChange.
func (c *clientImpl) PreviewTariffChange(
	ctx context.Context,
	req PreviewTariffChangeRequest,
) (*PreviewTariffChangeResponse, *http.Response, error) {
	httpReq, err := req.BuildRequest()
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

	var response PreviewTariffChangeResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&response); err != nil {
		return nil, httpRes, err
	}
	return &response, httpRes, nil
}
