package domain

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "authCode":
//        type: "string"
// required:
//    - "authCode"

type CreateDomainAuthCodeResponse struct {
	AuthCode string `json:"authCode"`
}

func (o *CreateDomainAuthCodeResponse) Validate() error {
	return nil
}
