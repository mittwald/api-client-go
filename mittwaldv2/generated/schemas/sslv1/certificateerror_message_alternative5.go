package sslv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "certificate_not_authorized_to_sign"
//description: "The certificate is not authorized to sign."

//The certificate is not authorized to sign.
type CertificateErrorMessageAlternative5 string

const CertificateErrorMessageAlternative5Certificatenotauthorizedtosign CertificateErrorMessageAlternative5 = "certificate_not_authorized_to_sign"

func (e CertificateErrorMessageAlternative5) Validate() error {
	if e == CertificateErrorMessageAlternative5Certificatenotauthorizedtosign {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}