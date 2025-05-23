package extensionv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "notStarted"
//    - "pending"
//    - "active"
//    - "terminationPending"

type SubscriptionBasedContractStatus string

const SubscriptionBasedContractStatusNotStarted SubscriptionBasedContractStatus = "notStarted"
const SubscriptionBasedContractStatusPending SubscriptionBasedContractStatus = "pending"
const SubscriptionBasedContractStatusActive SubscriptionBasedContractStatus = "active"
const SubscriptionBasedContractStatusTerminationPending SubscriptionBasedContractStatus = "terminationPending"

func (e SubscriptionBasedContractStatus) Validate() error {
	if e == SubscriptionBasedContractStatusNotStarted || e == SubscriptionBasedContractStatusPending || e == SubscriptionBasedContractStatusActive || e == SubscriptionBasedContractStatusTerminationPending {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
