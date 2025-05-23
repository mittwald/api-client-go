package conversationv2

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "category": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.Category"}
//    "conversationId":
//        type: "string"
//        format: "uuid"
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "createdBy": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.User"}
//    "lastMessage":
//        type: "object"
//        properties:
//            "createdAt":
//                type: "string"
//                format: "date-time"
//            "createdBy": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.User"}
//        required:
//            - "createdAt"
//    "lastMessageAt":
//        type: "string"
//        format: "date-time"
//    "lastMessageBy": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.User"}
//    "mainUser": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.User"}
//    "notificationRoles":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.conversation.NotificationRole"}
//    "relatedTo": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.RelatedAggregateReference"}
//    "relations":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.conversation.AggregateReference"}
//    "sharedWith": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.ShareableAggregateReference"}
//    "shortId":
//        type: "string"
//    "status": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.Status"}
//    "title":
//        type: "string"
//    "visibility":
//        type: "string"
//        enum:
//            - "shared"
//            - "private"
// required:
//    - "conversationId"
//    - "shortId"
//    - "title"
//    - "createdAt"
//    - "status"
//    - "visibility"
//    - "mainUser"

type Conversation struct {
	Category          *Category                    `json:"category,omitempty"`
	ConversationId    string                       `json:"conversationId"`
	CreatedAt         time.Time                    `json:"createdAt"`
	CreatedBy         *User                        `json:"createdBy,omitempty"`
	LastMessage       *ConversationLastMessage     `json:"lastMessage,omitempty"`
	LastMessageAt     *time.Time                   `json:"lastMessageAt,omitempty"`
	LastMessageBy     *User                        `json:"lastMessageBy,omitempty"`
	MainUser          User                         `json:"mainUser"`
	NotificationRoles []NotificationRole           `json:"notificationRoles,omitempty"`
	RelatedTo         *RelatedAggregateReference   `json:"relatedTo,omitempty"`
	Relations         []AggregateReference         `json:"relations,omitempty"`
	SharedWith        *ShareableAggregateReference `json:"sharedWith,omitempty"`
	ShortId           string                       `json:"shortId"`
	Status            Status                       `json:"status"`
	Title             string                       `json:"title"`
	Visibility        ConversationVisibility       `json:"visibility"`
}

func (o *Conversation) Validate() error {
	if err := func() error {
		if o.Category == nil {
			return nil
		}
		return o.Category.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property category: %w", err)
	}
	if err := func() error {
		if o.CreatedBy == nil {
			return nil
		}
		return o.CreatedBy.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property createdBy: %w", err)
	}
	if err := func() error {
		if o.LastMessage == nil {
			return nil
		}
		return o.LastMessage.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property lastMessage: %w", err)
	}
	if err := func() error {
		if o.LastMessageBy == nil {
			return nil
		}
		return o.LastMessageBy.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property lastMessageBy: %w", err)
	}
	if err := o.MainUser.Validate(); err != nil {
		return fmt.Errorf("invalid property mainUser: %w", err)
	}
	if err := func() error {
		if o.NotificationRoles == nil {
			return nil
		}
		return func() error {
			for i := range o.NotificationRoles {
				if err := o.NotificationRoles[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property notificationRoles: %w", err)
	}
	if err := func() error {
		if o.RelatedTo == nil {
			return nil
		}
		return o.RelatedTo.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property relatedTo: %w", err)
	}
	if err := func() error {
		if o.Relations == nil {
			return nil
		}
		return func() error {
			for i := range o.Relations {
				if err := o.Relations[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property relations: %w", err)
	}
	if err := func() error {
		if o.SharedWith == nil {
			return nil
		}
		return o.SharedWith.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property sharedWith: %w", err)
	}
	if err := o.Status.Validate(); err != nil {
		return fmt.Errorf("invalid property status: %w", err)
	}
	if err := o.Visibility.Validate(); err != nil {
		return fmt.Errorf("invalid property visibility: %w", err)
	}
	return nil
}
