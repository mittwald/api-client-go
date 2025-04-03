package appv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "mysql"
//    - "redis"

type LinkedDatabaseKind string

const LinkedDatabaseKindMysql LinkedDatabaseKind = "mysql"
const LinkedDatabaseKindRedis LinkedDatabaseKind = "redis"

func (e LinkedDatabaseKind) Validate() error {
	if e == LinkedDatabaseKindMysql || e == LinkedDatabaseKindRedis {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
