package sslv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "certificate_decode_failed"
// description: "Failed to decode the certificate."

// Failed to decode the certificate.
type CertificateErrorMessageAlternative2 string

const CertificateErrorMessageAlternative2Certificatedecodefailed CertificateErrorMessageAlternative2 = "certificate_decode_failed"

func (e CertificateErrorMessageAlternative2) Validate() error {
	if e == CertificateErrorMessageAlternative2Certificatedecodefailed {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
