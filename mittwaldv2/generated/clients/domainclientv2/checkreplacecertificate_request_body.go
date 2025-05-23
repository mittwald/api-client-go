package domainclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "certificate":
//        type: "string"
//    "privateKey":
//        type: "string"
// required:
//    - "certificate"
// description: CheckReplaceCertificateRequestBody models the JSON body of a 'ssl-check-replace-certificate' request

// CheckReplaceCertificateRequestBody models the JSON body of a 'ssl-check-replace-certificate' request
type CheckReplaceCertificateRequestBody struct {
	Certificate string  `json:"certificate"`
	PrivateKey  *string `json:"privateKey,omitempty"`
}

func (o *CheckReplaceCertificateRequestBody) Validate() error {
	return nil
}
