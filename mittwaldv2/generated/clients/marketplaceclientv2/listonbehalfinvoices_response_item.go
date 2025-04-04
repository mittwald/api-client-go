package marketplaceclientv2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "invoiceDate":
//        type: "string"
//        format: "date-time"
//    "invoiceId":
//        type: "string"
//    "invoiceNumber":
//        type: "string"
//    "pdfLink":
//        type: "string"
//    "totalGross":
//        type: "number"
//    "webLink":
//        type: "string"
// required:
//    - "invoiceId"
//    - "invoiceNumber"
//    - "invoiceDate"
//    - "totalGross"
//    - "webLink"
//    - "pdfLink"

type ListOnbehalfInvoicesResponseItem struct {
	InvoiceDate   time.Time `json:"invoiceDate"`
	InvoiceId     string    `json:"invoiceId"`
	InvoiceNumber string    `json:"invoiceNumber"`
	PdfLink       string    `json:"pdfLink"`
	TotalGross    float64   `json:"totalGross"`
	WebLink       string    `json:"webLink"`
}

func (o *ListOnbehalfInvoicesResponseItem) Validate() error {
	return nil
}
