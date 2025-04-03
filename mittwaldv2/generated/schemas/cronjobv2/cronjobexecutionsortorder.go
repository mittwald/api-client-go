package cronjobv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "oldestFirst"
//    - "newestFirst"
//    - "slowestFirst"
//    - "fastestFirst"

type CronjobExecutionSortOrder string

const CronjobExecutionSortOrderOldestFirst CronjobExecutionSortOrder = "oldestFirst"
const CronjobExecutionSortOrderNewestFirst CronjobExecutionSortOrder = "newestFirst"
const CronjobExecutionSortOrderSlowestFirst CronjobExecutionSortOrder = "slowestFirst"
const CronjobExecutionSortOrderFastestFirst CronjobExecutionSortOrder = "fastestFirst"

func (e CronjobExecutionSortOrder) Validate() error {
	if e == CronjobExecutionSortOrderOldestFirst || e == CronjobExecutionSortOrderNewestFirst || e == CronjobExecutionSortOrderSlowestFirst || e == CronjobExecutionSortOrderFastestFirst {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
