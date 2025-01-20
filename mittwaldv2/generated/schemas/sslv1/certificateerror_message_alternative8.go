package sslv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "too_many_intermediates"
//description: "There are too many intermediate certificates."

//There are too many intermediate certificates.
type CertificateErrorMessageAlternative8 string

const CertificateErrorMessageAlternative8Toomanyintermediates CertificateErrorMessageAlternative8 = "too_many_intermediates"

func (e CertificateErrorMessageAlternative8) Validate() error {
	if e == CertificateErrorMessageAlternative8Toomanyintermediates {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}