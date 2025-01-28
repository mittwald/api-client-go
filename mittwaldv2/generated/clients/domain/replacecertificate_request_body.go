package domain

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "certificate":
//        type: "string"
//    "privateKey":
//        type: "string"
// required:
//    - "certificate"
// description: ReplaceCertificateRequestBody models the JSON body of a 'ssl-replace-certificate' request

// ReplaceCertificateRequestBody models the JSON body of a 'ssl-replace-certificate' request
type ReplaceCertificateRequestBody struct {
	Certificate string  `json:"certificate"`
	PrivateKey  *string `json:"privateKey,omitempty"`
}

func (o *ReplaceCertificateRequestBody) Validate() error {
	return nil
}
