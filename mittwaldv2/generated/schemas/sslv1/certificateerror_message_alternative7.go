package sslv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "ca_not_authorized_for_this_name"
//description: "The Certificate Authority (CA) is not authorized for this name."

//The Certificate Authority (CA) is not authorized for this name.
type CertificateErrorMessageAlternative7 string

const CertificateErrorMessageAlternative7Canotauthorizedforthisname CertificateErrorMessageAlternative7 = "ca_not_authorized_for_this_name"

func (e CertificateErrorMessageAlternative7) Validate() error {
	if e == CertificateErrorMessageAlternative7Canotauthorizedforthisname {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
