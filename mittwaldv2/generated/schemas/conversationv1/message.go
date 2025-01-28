package conversationv1

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "conversationId":
//        type: "string"
//        format: "uuid"
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "createdBy": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.User"}
//    "files":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.conversation.File"}
//    "internal":
//        type: "boolean"
//    "messageContent":
//        type: "string"
//    "messageId":
//        type: "string"
//        format: "uuid"
//    "type":
//        type: "string"
//        enum:
//            - "MESSAGE"
// required:
//    - "messageId"
//    - "conversationId"
//    - "createdAt"
//    - "type"

type Message struct {
	ConversationId string      `json:"conversationId"`
	CreatedAt      time.Time   `json:"createdAt"`
	CreatedBy      *User       `json:"createdBy,omitempty"`
	Files          []File      `json:"files,omitempty"`
	Internal       *bool       `json:"internal,omitempty"`
	MessageContent *string     `json:"messageContent,omitempty"`
	MessageId      string      `json:"messageId"`
	Type           MessageType `json:"type"`
}

func (o *Message) Validate() error {
	if err := func() error {
		if o.CreatedBy == nil {
			return nil
		}
		return o.CreatedBy.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property createdBy: %w", err)
	}
	if err := func() error {
		if o.Files == nil {
			return nil
		}
		return func() error {
			for i := range o.Files {
				if err := o.Files[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property files: %w", err)
	}
	if err := o.Type.Validate(); err != nil {
		return fmt.Errorf("invalid property type: %w", err)
	}
	return nil
}
