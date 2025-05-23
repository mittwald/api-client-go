package conversationv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "categoryUserPairs":
//        type: "object"
//        additionalProperties:
//            type: "string"
//            format: "uuid"
//    "fallback":
//        type: "string"
//        format: "uuid"

type ConversationPreferencesPreferredUsers struct {
	CategoryUserPairs map[string]string `json:"categoryUserPairs,omitempty"`
	Fallback          *string           `json:"fallback,omitempty"`
}

func (o *ConversationPreferencesPreferredUsers) Validate() error {
	return nil
}
