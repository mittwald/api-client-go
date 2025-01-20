package sslv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "unknown_authority"
//description: "The certificate is signed by an unknown authority."

//The certificate is signed by an unknown authority.
type CertificateErrorMessageAlternative10 string

const CertificateErrorMessageAlternative10Unknownauthority CertificateErrorMessageAlternative10 = "unknown_authority"

func (e CertificateErrorMessageAlternative10) Validate() error {
	if e == CertificateErrorMessageAlternative10Unknownauthority {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}