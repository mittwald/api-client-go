package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//        description: "Expiration of the support code"
//    "supportCode":
//        type: "string"
//        description: "support code to authenticate yourself against the mittwald support via telephone"
//        example: "123456"
//required:
//    - "supportCode"
//    - "expiresAt"

//
type SupportCodeRequestResponse struct {
	ExpiresAt   string `json:"expiresAt"`
	SupportCode string `json:"supportCode"`
}

func (o *SupportCodeRequestResponse) Validate() error {
	return nil
}
