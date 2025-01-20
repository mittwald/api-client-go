package invoicev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "sourceInvoiceId":
//        type: "string"
//        format: "uuid"
//    "sourceInvoiceItemId":
//        type: "string"
//        format: "uuid"
//required:
//    - "sourceInvoiceId"
//    - "sourceInvoiceItemId"

//
type InvoiceItemReference struct {
	SourceInvoiceId     uuid.UUID `json:"sourceInvoiceId"`
	SourceInvoiceItemId uuid.UUID `json:"sourceInvoiceItemId"`
}

func (o *InvoiceItemReference) Validate() error {
	return nil
}
