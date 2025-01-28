package invoicev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "contractItemId":
//        type: "string"
//        format: "uuid"
//    "creditPeriod": {"$ref": "#/components/schemas/de.mittwald.v1.invoice.DatePeriod"}
// required:
//    - "contractItemId"
//    - "creditPeriod"

type CorrectionDefinition struct {
	ContractItemId string     `json:"contractItemId"`
	CreditPeriod   DatePeriod `json:"creditPeriod"`
}

func (o *CorrectionDefinition) Validate() error {
	if err := o.CreditPeriod.Validate(); err != nil {
		return fmt.Errorf("invalid property creditPeriod: %w", err)
	}
	return nil
}
