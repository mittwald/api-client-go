package appv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "start"
//    - "stop"
//    - "restart"
// description: "An Action is a string that describes a runtime concerning action which can be executed on an AppInstallation or an  App  can be capable of."

// An Action is a string that describes a runtime concerning action which can be executed on an AppInstallation or an  App  can be capable of.
type Action string

const ActionStart Action = "start"
const ActionStop Action = "stop"
const ActionRestart Action = "restart"

func (e Action) Validate() error {
	if e == ActionStart || e == ActionStop || e == ActionRestart {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
