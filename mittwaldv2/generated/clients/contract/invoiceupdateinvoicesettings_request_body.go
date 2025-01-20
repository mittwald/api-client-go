package contract

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/invoicev1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "additionalEmailRecipients":
//        type: "array"
//        items:
//            type: "string"
//            format: "email"
//        example: ["hallo@test.de"]
//    "invoicePeriod":
//        type: "integer"
//        minimum: 0
//        example: 1
//    "paymentSettings": {"$ref": "#/components/schemas/de.mittwald.v1.invoice.PaymentSettings"}
//    "printedInvoices":
//        type: "boolean"
//    "recipient": {"$ref": "#/components/schemas/de.mittwald.v1.invoice.Recipient"}
//    "recipientSameAsOwner":
//        type: "boolean"
//    "targetDay":
//        type: "integer"
//        maximum: 28
//        minimum: 0
//        example: 15
//required:
//    - "invoicePeriod"
//    - "paymentSettings"

//
type InvoiceUpdateInvoiceSettingsRequestBody struct {
	AdditionalEmailRecipients []string                  `json:"additionalEmailRecipients,omitempty"`
	InvoicePeriod             int64                     `json:"invoicePeriod"`
	PaymentSettings           invoicev1.PaymentSettings `json:"paymentSettings"`
	PrintedInvoices           *bool                     `json:"printedInvoices,omitempty"`
	Recipient                 *invoicev1.Recipient      `json:"recipient,omitempty"`
	RecipientSameAsOwner      *bool                     `json:"recipientSameAsOwner,omitempty"`
	TargetDay                 *int64                    `json:"targetDay,omitempty"`
}

func (o *InvoiceUpdateInvoiceSettingsRequestBody) Validate() error {
	if err := func() error {
		if o.AdditionalEmailRecipients == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property additionalEmailRecipients: %w", err)
	}
	if err := o.PaymentSettings.Validate(); err != nil {
		return fmt.Errorf("invalid property paymentSettings: %w", err)
	}
	if err := func() error {
		if o.Recipient == nil {
			return nil
		}
		return o.Recipient.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property recipient: %w", err)
	}
	return nil
}
