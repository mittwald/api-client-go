package appv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "lastExitCode":
//        type: "number"
//    "logFileLocation":
//        type: "string"
//    "state":
//        type: "string"
//        enum:
//            - "running"
//            - "stopped"
//            - "exited"
//    "uptimeSeconds":
//        type: "number"
//required:
//    - "state"
//    - "logFileLocation"
//description: "AppInstallationStatus describes the overall runtime status of an AppInstallation."

// AppInstallationStatus describes the overall runtime status of an AppInstallation.
type AppInstallationStatus struct {
	LastExitCode    *float64                   `json:"lastExitCode,omitempty"`
	LogFileLocation string                     `json:"logFileLocation"`
	State           AppInstallationStatusState `json:"state"`
	UptimeSeconds   *float64                   `json:"uptimeSeconds,omitempty"`
}

func (o *AppInstallationStatus) Validate() error {
	if err := o.State.Validate(); err != nil {
		return fmt.Errorf("invalid property state: %w", err)
	}
	return nil
}
