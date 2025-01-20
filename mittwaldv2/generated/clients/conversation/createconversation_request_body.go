package conversation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/conversationv1"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "categoryId":
//        type: "string"
//    "mainUserId":
//        type: "string"
//        format: "uuid"
//    "notificationRoles":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.conversation.NotificationRole"}
//    "relatedTo": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.RelatedAggregateReference"}
//    "sharedWith": {"$ref": "#/components/schemas/de.mittwald.v1.conversation.ShareableAggregateReference"}
//    "title":
//        type: "string"

//
type CreateConversationRequestBody struct {
	CategoryId        *string                                     `json:"categoryId,omitempty"`
	MainUserId        *uuid.UUID                                  `json:"mainUserId,omitempty"`
	NotificationRoles []conversationv1.NotificationRole           `json:"notificationRoles,omitempty"`
	RelatedTo         *conversationv1.RelatedAggregateReference   `json:"relatedTo,omitempty"`
	SharedWith        *conversationv1.ShareableAggregateReference `json:"sharedWith,omitempty"`
	Title             *string                                     `json:"title,omitempty"`
}

func (o *CreateConversationRequestBody) Validate() error {
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
		if o.SharedWith == nil {
			return nil
		}
		return o.SharedWith.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property sharedWith: %w", err)
	}
	return nil
}