package signupv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "apiTokenId":
//        type: "string"
//        format: "uuid"
//    "createdAt":
//        type: "string"
//        format: "date-time"
//    "description":
//        type: "string"
//        example: "Api Token - read"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//    "roles":
//        type: "array"
//        items:
//            type: "string"
//            enum:
//                - "api_read"
//                - "api_write"
//required:
//    - "apiTokenId"
//    - "roles"
//    - "description"
//    - "createdAt"

//
type ApiToken struct {
	ApiTokenId  uuid.UUID           `json:"apiTokenId"`
	CreatedAt   time.Time           `json:"createdAt"`
	Description string              `json:"description"`
	ExpiresAt   *time.Time          `json:"expiresAt,omitempty"`
	Roles       []ApiTokenRolesItem `json:"roles"`
}

func (o *ApiToken) Validate() error {
	if o.Roles == nil {
		return errors.New("property roles is required, but not set")
	}
	if err := func() error {
		for i := range o.Roles {
			if err := o.Roles[i].Validate(); err != nil {
				return fmt.Errorf("item %d is invalid %w", i, err)
			}
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property roles: %w", err)
	}
	return nil
}
