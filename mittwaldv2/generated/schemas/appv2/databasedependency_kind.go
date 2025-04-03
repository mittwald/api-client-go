package appv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "mysql"

type DatabaseDependencyKind string

const DatabaseDependencyKindMysql DatabaseDependencyKind = "mysql"

func (e DatabaseDependencyKind) Validate() error {
	if e == DatabaseDependencyKindMysql {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
