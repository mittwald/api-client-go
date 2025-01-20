package sslv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "private_key_encrypted"
//description: "The private key is encrypted."

//The private key is encrypted.
type CertificateErrorMessageAlternative14 string

const CertificateErrorMessageAlternative14Privatekeyencrypted CertificateErrorMessageAlternative14 = "private_key_encrypted"

func (e CertificateErrorMessageAlternative14) Validate() error {
	if e == CertificateErrorMessageAlternative14Privatekeyencrypted {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
