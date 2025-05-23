package invoicev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "NEW"
//    - "CONFIRMED"
//    - "DENIED"
//    - "PAID"
//    - "PARTIALLY_PAID"
//    - "OVERPAID"

type InvoiceStatus string

const InvoiceStatusNEW InvoiceStatus = "NEW"
const InvoiceStatusCONFIRMED InvoiceStatus = "CONFIRMED"
const InvoiceStatusDENIED InvoiceStatus = "DENIED"
const InvoiceStatusPAID InvoiceStatus = "PAID"
const InvoiceStatusPARTIALLYPAID InvoiceStatus = "PARTIALLY_PAID"
const InvoiceStatusOVERPAID InvoiceStatus = "OVERPAID"

func (e InvoiceStatus) Validate() error {
	if e == InvoiceStatusNEW || e == InvoiceStatusCONFIRMED || e == InvoiceStatusDENIED || e == InvoiceStatusPAID || e == InvoiceStatusPARTIALLYPAID || e == InvoiceStatusOVERPAID {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
