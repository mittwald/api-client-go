package sslv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "private_key_decode_failed"
// description: "Failed to decode the private key."

// Failed to decode the private key.
type CertificateErrorMessageAlternative12 string

const CertificateErrorMessageAlternative12Privatekeydecodefailed CertificateErrorMessageAlternative12 = "private_key_decode_failed"

func (e CertificateErrorMessageAlternative12) Validate() error {
	if e == CertificateErrorMessageAlternative12Privatekeydecodefailed {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
