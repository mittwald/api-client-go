package misc

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "exists":
//        type: "boolean"
//required:
//    - "exists"

type VerificationVerifyCompanyResponse struct {
	Exists bool `json:"exists"`
}

func (o *VerificationVerifyCompanyResponse) Validate() error {
	return nil
}
