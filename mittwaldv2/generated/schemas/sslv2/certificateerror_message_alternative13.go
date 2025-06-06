package sslv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "private_key_parse_failed"
// description: "An error occurred while parsing the private key."

// An error occurred while parsing the private key.
type CertificateErrorMessageAlternative13 string

const CertificateErrorMessageAlternative13PrivateKeyParseFailed CertificateErrorMessageAlternative13 = "private_key_parse_failed"

func (e CertificateErrorMessageAlternative13) Validate() error {
	if e == CertificateErrorMessageAlternative13PrivateKeyParseFailed {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
