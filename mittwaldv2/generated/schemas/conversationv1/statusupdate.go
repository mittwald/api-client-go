package conversationv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "conversationId":
//        type: "string"
//    "createdAt":
//        type: "string"
//    "internal":
//        type: "boolean"
//    "messageContent":
//        type: "string"
//    "meta":
//        type: "object"
//        properties:
//            "user": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.User"}
//    "type":
//        type: "string"
//        enum:
//            - "STATUS_UPDATE"
// required:
//    - "conversationId"
//    - "messageContent"
//    - "createdAt"
//    - "type"

type StatusUpdate struct {
	ConversationId string            `json:"conversationId"`
	CreatedAt      string            `json:"createdAt"`
	Internal       *bool             `json:"internal,omitempty"`
	MessageContent string            `json:"messageContent"`
	Meta           *StatusUpdateMeta `json:"meta,omitempty"`
	Type           StatusUpdateType  `json:"type"`
}

func (o *StatusUpdate) Validate() error {
	if err := func() error {
		if o.Meta == nil {
			return nil
		}
		return o.Meta.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property meta: %w", err)
	}
	if err := o.Type.Validate(); err != nil {
		return fmt.Errorf("invalid property type: %w", err)
	}
	return nil
}
