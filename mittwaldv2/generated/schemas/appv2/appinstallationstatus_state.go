package appv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "running"
//    - "stopped"
//    - "exited"

type AppInstallationStatusState string

const AppInstallationStatusStateRunning AppInstallationStatusState = "running"
const AppInstallationStatusStateStopped AppInstallationStatusState = "stopped"
const AppInstallationStatusStateExited AppInstallationStatusState = "exited"

func (e AppInstallationStatusState) Validate() error {
	if e == AppInstallationStatusStateRunning || e == AppInstallationStatusStateStopped || e == AppInstallationStatusStateExited {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
