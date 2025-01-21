package invoicev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "message":
//        type: "string"
//        example: "Unfortunately, we were unable to debit your account. Please update your account details."
//    "severity":
//        type: "string"
//        enum:
//            - "success"
//            - "info"
//            - "warning"
//            - "error"
//    "type":
//        type: "string"
//        enum:
//            - "returnDebitNote"
//            - "returnDebitNoteWaitingForPayment"
//            - "debtWrittenOff"
//            - "bankrupt"
//required:
//    - "severity"
//    - "type"
//    - "message"

type InvoiceSettingsStatus struct {
	Message  string                        `json:"message"`
	Severity InvoiceSettingsStatusSeverity `json:"severity"`
	Type     InvoiceSettingsStatusType     `json:"type"`
}

func (o *InvoiceSettingsStatus) Validate() error {
	if err := o.Severity.Validate(); err != nil {
		return fmt.Errorf("invalid property severity: %w", err)
	}
	if err := o.Type.Validate(); err != nil {
		return fmt.Errorf("invalid property type: %w", err)
	}
	return nil
}
