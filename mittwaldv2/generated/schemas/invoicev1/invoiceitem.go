package invoicev1

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "additionalDescription":
//        type: "string"
//        example: "Short-ID: \"s-123456\""
//    "contractItemId":
//        type: "string"
//        format: "uuid"
//    "description":
//        type: "string"
//        example: "Space-Server: \"Mein Space-Server\""
//    "itemCancelledOrCorrectedBy":
//        type: "array"
//        items:
//            type: "object"
//            properties:
//                "sourceInvoiceId":
//                    type: "string"
//                    format: "uuid"
//                "sourceInvoiceItemId":
//                    type: "string"
//                    format: "uuid"
//    "itemId":
//        type: "string"
//        format: "uuid"
//    "price": {"$ref": "#/components/schemas/de.mittwald.v1.invoice.Price"}
//    "reference":
//        type: "object"
//        properties:
//            "sourceInvoiceId":
//                type: "string"
//                format: "uuid"
//            "sourceInvoiceItemId":
//                type: "string"
//                format: "uuid"
//        required:
//            - "sourceInvoiceId"
//            - "sourceInvoiceItemId"
//    "serviceDate":
//        type: "string"
//        format: "date-time"
//    "servicePeriod": {"$ref": "#/components/schemas/de.mittwald.v1.invoice.DatePeriod"}
//    "vatRate":
//        type: "number"
//        example: 19
// required:
//    - "itemId"
//    - "price"
//    - "vatRate"
//    - "contractItemId"
//    - "description"

type InvoiceItem struct {
	AdditionalDescription      *string                                     `json:"additionalDescription,omitempty"`
	ContractItemId             string                                      `json:"contractItemId"`
	Description                string                                      `json:"description"`
	ItemCancelledOrCorrectedBy []InvoiceItemItemCancelledOrCorrectedByItem `json:"itemCancelledOrCorrectedBy,omitempty"`
	ItemId                     string                                      `json:"itemId"`
	Price                      Price                                       `json:"price"`
	Reference                  *InvoiceItemReference                       `json:"reference,omitempty"`
	ServiceDate                *time.Time                                  `json:"serviceDate,omitempty"`
	ServicePeriod              *DatePeriod                                 `json:"servicePeriod,omitempty"`
	VatRate                    float64                                     `json:"vatRate"`
}

func (o *InvoiceItem) Validate() error {
	if err := func() error {
		if o.ItemCancelledOrCorrectedBy == nil {
			return nil
		}
		return func() error {
			for i := range o.ItemCancelledOrCorrectedBy {
				if err := o.ItemCancelledOrCorrectedBy[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property itemCancelledOrCorrectedBy: %w", err)
	}
	if err := o.Price.Validate(); err != nil {
		return fmt.Errorf("invalid property price: %w", err)
	}
	if err := func() error {
		if o.Reference == nil {
			return nil
		}
		return o.Reference.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property reference: %w", err)
	}
	if err := func() error {
		if o.ServicePeriod == nil {
			return nil
		}
		return o.ServicePeriod.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property servicePeriod: %w", err)
	}
	return nil
}
