package cronjobv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "Complete"
//    - "Failed"
//    - "AbortedBySystem"
//    - "Pending"
//    - "Running"
//    - "AbortedByUser"
//    - "TimedOut"

type CronjobExecutionStatus string

const CronjobExecutionStatusComplete CronjobExecutionStatus = "Complete"
const CronjobExecutionStatusFailed CronjobExecutionStatus = "Failed"
const CronjobExecutionStatusAbortedBySystem CronjobExecutionStatus = "AbortedBySystem"
const CronjobExecutionStatusPending CronjobExecutionStatus = "Pending"
const CronjobExecutionStatusRunning CronjobExecutionStatus = "Running"
const CronjobExecutionStatusAbortedByUser CronjobExecutionStatus = "AbortedByUser"
const CronjobExecutionStatusTimedOut CronjobExecutionStatus = "TimedOut"

func (e CronjobExecutionStatus) Validate() error {
	if e == CronjobExecutionStatusComplete || e == CronjobExecutionStatusFailed || e == CronjobExecutionStatusAbortedBySystem || e == CronjobExecutionStatusPending || e == CronjobExecutionStatusRunning || e == CronjobExecutionStatusAbortedByUser || e == CronjobExecutionStatusTimedOut {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
