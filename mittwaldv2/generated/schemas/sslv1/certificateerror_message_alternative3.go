package sslv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "certificate_parsing_failed"
//description: "An error occurred while parsing the certificate."

// An error occurred while parsing the certificate.
type CertificateErrorMessageAlternative3 string

const CertificateErrorMessageAlternative3Certificateparsingfailed CertificateErrorMessageAlternative3 = "certificate_parsing_failed"

func (e CertificateErrorMessageAlternative3) Validate() error {
	if e == CertificateErrorMessageAlternative3Certificateparsingfailed {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
