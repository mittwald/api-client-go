package invoicev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "method":
//        type: "string"
//        enum:
//            - "invoice"
// required:
//    - "method"

type PaymentSettingsInvoice struct {
	Method PaymentSettingsInvoiceMethod `json:"method"`
}

func (o *PaymentSettingsInvoice) Validate() error {
	if err := o.Method.Validate(); err != nil {
		return fmt.Errorf("invalid property method: %w", err)
	}
	return nil
}
