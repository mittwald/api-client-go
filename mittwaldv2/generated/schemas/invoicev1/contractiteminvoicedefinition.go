package invoicev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "contractItemId":
//        type: "string"
//        format: "uuid"
//    "isDue":
//        type: "boolean"
//    "serviceDate":
//        type: "string"
//        format: "date-time"
//    "servicePeriod": {"$ref": "#/components/schemas/de.mittwald.v1.invoice.DatePeriod"}
//    "vatRate":
//        type: "integer"
//        example: 19
//required:
//    - "contractItemId"
//    - "vatRate"
//    - "servicePeriod"

type ContractItemInvoiceDefinition struct {
	ContractItemId uuid.UUID  `json:"contractItemId"`
	IsDue          *bool      `json:"isDue,omitempty"`
	ServiceDate    *time.Time `json:"serviceDate,omitempty"`
	ServicePeriod  DatePeriod `json:"servicePeriod"`
	VatRate        int64      `json:"vatRate"`
}

func (o *ContractItemInvoiceDefinition) Validate() error {
	if err := o.ServicePeriod.Validate(); err != nil {
		return fmt.Errorf("invalid property servicePeriod: %w", err)
	}
	return nil
}
