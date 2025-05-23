package invoicev2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "cancellationId":
//        type: "string"
//        format: "uuid"
//    "cancelledAt":
//        type: "string"
//        format: "date-time"
//    "correctionNumber":
//        type: "string"
//        example: "RG1234567"
//    "pdfId":
//        type: "string"
//        format: "uuid"
//    "reason":
//        type: "string"
//        example: "Kulanz"
// required:
//    - "pdfId"
//    - "correctionNumber"
//    - "cancelledAt"
//    - "cancellationId"

type Cancellation struct {
	CancellationId   string    `json:"cancellationId"`
	CancelledAt      time.Time `json:"cancelledAt"`
	CorrectionNumber string    `json:"correctionNumber"`
	PdfId            string    `json:"pdfId"`
	Reason           *string   `json:"reason,omitempty"`
}

func (o *Cancellation) Validate() error {
	return nil
}
