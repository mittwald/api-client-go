package invoicev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "returnDebitNote"
//    - "returnDebitNoteWaitingForPayment"
//    - "debtWrittenOff"
//    - "bankrupt"

type InvoiceSettingsStatusType string

const InvoiceSettingsStatusTypeReturnDebitNote InvoiceSettingsStatusType = "returnDebitNote"
const InvoiceSettingsStatusTypeReturnDebitNoteWaitingForPayment InvoiceSettingsStatusType = "returnDebitNoteWaitingForPayment"
const InvoiceSettingsStatusTypeDebtWrittenOff InvoiceSettingsStatusType = "debtWrittenOff"
const InvoiceSettingsStatusTypeBankrupt InvoiceSettingsStatusType = "bankrupt"

func (e InvoiceSettingsStatusType) Validate() error {
	if e == InvoiceSettingsStatusTypeReturnDebitNote || e == InvoiceSettingsStatusTypeReturnDebitNoteWaitingForPayment || e == InvoiceSettingsStatusTypeDebtWrittenOff || e == InvoiceSettingsStatusTypeBankrupt {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
