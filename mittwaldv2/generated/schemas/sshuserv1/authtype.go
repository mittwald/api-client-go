package sshuserv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "password"
//    - "publicKey"
//description: "Method of authentication that a given SFTPUser or SSHuser supports."

//Method of authentication that a given SFTPUser or SSHuser supports.
type AuthType string

const AuthTypePassword AuthType = "password"
const AuthTypePublicKey AuthType = "publicKey"

func (e AuthType) Validate() error {
	if e == AuthTypePassword || e == AuthTypePublicKey {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}