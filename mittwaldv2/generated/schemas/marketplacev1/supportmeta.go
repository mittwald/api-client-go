package marketplacev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "email":
//        type: "string"
//        example: "a.lovelace@example.com"
//    "phone":
//        type: "string"

//
type SupportMeta struct {
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

func (o *SupportMeta) Validate() error {
	return nil
}
