package orderv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "contractChangeContractId":
//        type: "string"
//        format: "uuid"
//    "customerId":
//        type: "string"
//        minLength: 1
//        example: "4317f5c5-1ea8-4084-a9f1-3a8e7e1c94ff"
//    "dueDate":
//        type: "string"
//        format: "date-time"
//    "invoicingPeriod":
//        type: "number"
//        description: "Invoicing period in months"
//        example: 12
//    "items":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.order.OrderItem"}
//    "orderDate":
//        type: "string"
//        format: "date-time"
//    "orderId":
//        type: "string"
//        format: "uuid"
//    "orderNumber":
//        type: "string"
//        minLength: 1
//        example: "5XXXXXX"
//    "profile": {"$ref": "#/components/schemas/de.mittwald.v1.order.Profile"}
//    "status": {"$ref": "#/components/schemas/de.mittwald.v1.order.OrderStatus"}
//    "summary": {"$ref": "#/components/schemas/de.mittwald.v1.order.OrderSummary"}
//    "type": {"$ref": "#/components/schemas/de.mittwald.v1.order.OrderType"}
//required:
//    - "orderId"
//    - "orderNumber"
//    - "summary"
//    - "customerId"
//    - "invoicingPeriod"
//    - "type"
//    - "status"
//    - "items"

//
type CustomerOrder struct {
	ContractChangeContractId *uuid.UUID   `json:"contractChangeContractId,omitempty"`
	CustomerId               string       `json:"customerId"`
	DueDate                  *time.Time   `json:"dueDate,omitempty"`
	InvoicingPeriod          float64      `json:"invoicingPeriod"`
	Items                    []OrderItem  `json:"items"`
	OrderDate                *time.Time   `json:"orderDate,omitempty"`
	OrderId                  uuid.UUID    `json:"orderId"`
	OrderNumber              string       `json:"orderNumber"`
	Profile                  *Profile     `json:"profile,omitempty"`
	Status                   OrderStatus  `json:"status"`
	Summary                  OrderSummary `json:"summary"`
	Type                     OrderType    `json:"type"`
}

func (o *CustomerOrder) Validate() error {
	if o.Items == nil {
		return errors.New("property items is required, but not set")
	}
	if err := func() error {
		for i := range o.Items {
			if err := o.Items[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property items: %w", err)
	}
	if err := func() error {
		if o.Profile == nil {
			return nil
		}
		return o.Profile.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property profile: %w", err)
	}
	if err := o.Status.Validate(); err != nil {
		return fmt.Errorf("invalid property status: %w", err)
	}
	if err := o.Summary.Validate(); err != nil {
		return fmt.Errorf("invalid property summary: %w", err)
	}
	if err := o.Type.Validate(); err != nil {
		return fmt.Errorf("invalid property type: %w", err)
	}
	return nil
}
