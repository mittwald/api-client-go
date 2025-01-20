package sshuserv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "ssh"
//    - "sftp"

type EntityTypes string

const EntityTypesSSH EntityTypes = "ssh"
const EntityTypesSFTP EntityTypes = "sftp"

func (e EntityTypes) Validate() error {
	if e == EntityTypesSSH || e == EntityTypesSFTP {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
