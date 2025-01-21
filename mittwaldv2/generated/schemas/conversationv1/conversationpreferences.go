package conversationv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "customerId":
//        type: "string"
//        format: "uuid"
//    "preferredUsers":
//        type: "object"
//        properties:
//            "categoryUserPairs":
//                type: "object"
//                additionalProperties:
//                    type: "string"
//                    format: "uuid"
//            "fallback":
//                type: "string"
//                format: "uuid"
//required:
//    - "customerId"
//    - "preferredUsers"

type ConversationPreferences struct {
	CustomerId     uuid.UUID                             `json:"customerId"`
	PreferredUsers ConversationPreferencesPreferredUsers `json:"preferredUsers"`
}

func (o *ConversationPreferences) Validate() error {
	if err := o.PreferredUsers.Validate(); err != nil {
		return fmt.Errorf("invalid property preferredUsers: %w", err)
	}
	return nil
}
