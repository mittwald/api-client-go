package marketplacev2

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "deprecatedAt":
//        type: "string"
//        format: "date-time"
//    "note":
//        type: "string"
//        example: "This extension is no longer actively maintained. Please Use the successor extension instead."
//    "successorId":
//        type: "string"
//        format: "uuid"
//        description: "The ID of the successor extension."
// required:
//    - "deprecatedAt"
// description: "The Extension is deprecated by the contributor and will expire at the given date."

// The Extension is deprecated by the contributor and will expire at the given date.
type ExtensionDeprecation struct {
	DeprecatedAt time.Time `json:"deprecatedAt"`
	Note         *string   `json:"note,omitempty"`
	SuccessorId  *string   `json:"successorId,omitempty"`
}

func (o *ExtensionDeprecation) Validate() error {
	return nil
}
