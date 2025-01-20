package conversationv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "conversationId":
//        type: "string"
//        format: "uuid"
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "files":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.conversation.File"}
//    "internal":
//        type: "boolean"
//    "messageContent":
//        type: "string"
//        enum:
//            - "relocation"
//            - "call"
//    "messageId":
//        type: "string"
//        format: "uuid"
//    "meta":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.conversation.ServiceRequestRelocationPayload"}
//    "type":
//        type: "string"
//        enum:
//            - "SERVICE_REQUEST"
//required:
//    - "conversationId"
//    - "messageId"
//    - "messageContent"
//    - "createdAt"
//    - "type"
//    - "internal"
//    - "meta"

//
type ServiceRequest struct {
	ConversationId uuid.UUID                    `json:"conversationId"`
	CreatedAt      time.Time                    `json:"createdAt"`
	Files          []File                       `json:"files,omitempty"`
	Internal       bool                         `json:"internal"`
	MessageContent ServiceRequestMessageContent `json:"messageContent"`
	MessageId      uuid.UUID                    `json:"messageId"`
	Meta           ServiceRequestMeta           `json:"meta"`
	Type           ServiceRequestType           `json:"type"`
}

func (o *ServiceRequest) Validate() error {
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
	if err := o.MessageContent.Validate(); err != nil {
		return fmt.Errorf("invalid property messageContent: %w", err)
	}
	if err := o.Meta.Validate(); err != nil {
		return fmt.Errorf("invalid property meta: %w", err)
	}
	if err := o.Type.Validate(); err != nil {
		return fmt.Errorf("invalid property type: %w", err)
	}
	return nil
}