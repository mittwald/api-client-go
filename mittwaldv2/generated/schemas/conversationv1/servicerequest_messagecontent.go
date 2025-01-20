package conversationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "relocation"
//    - "call"

type ServiceRequestMessageContent string

const ServiceRequestMessageContentRelocation ServiceRequestMessageContent = "relocation"
const ServiceRequestMessageContentCall ServiceRequestMessageContent = "call"

func (e ServiceRequestMessageContent) Validate() error {
	if e == ServiceRequestMessageContentRelocation || e == ServiceRequestMessageContentCall {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
