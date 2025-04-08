package customerv2

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "activeSuspension":
//        type: "object"
//        properties:
//            "createdAt":
//                type: "string"
//                format: "date-time"
//        required:
//            - "createdAt"
//    "avatarRefId":
//        type: "string"
//    "categoryId":
//        type: "string"
//    "creationDate":
//        type: "string"
//        format: "date-time"
//    "customerId":
//        type: "string"
//    "customerNumber":
//        type: "string"
//    "executingUserRoles":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.customer.Role"}
//    "flags":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.customer.CustomerFlag"}
//    "isAllowedToPlaceOrders":
//        type: "boolean"
//    "isBanned":
//        type: "boolean"
//    "isInDefaultOfPayment":
//        type: "boolean"
//    "levelOfUndeliverableDunningNotice":
//        type: "string"
//        enum:
//            - "first"
//            - "second"
//    "memberCount":
//        type: "integer"
//        minimum: 0
//    "name":
//        type: "string"
//    "owner": {"$ref": "#/components/schemas/de.mittwald.v1.customer.Contact"}
//    "projectCount":
//        type: "integer"
//        minimum: 0
//    "vatId":
//        type: "string"
//    "vatIdValidationState":
//        type: "string"
//        enum:
//            - "valid"
//            - "invalid"
//            - "pending"
//            - "unspecified"
// required:
//    - "customerId"
//    - "customerNumber"
//    - "name"
//    - "creationDate"
//    - "memberCount"
//    - "projectCount"

type Customer struct {
	ActiveSuspension                  *CustomerActiveSuspension                  `json:"activeSuspension,omitempty"`
	AvatarRefId                       *string                                    `json:"avatarRefId,omitempty"`
	CategoryId                        *string                                    `json:"categoryId,omitempty"`
	CreationDate                      time.Time                                  `json:"creationDate"`
	CustomerId                        string                                     `json:"customerId"`
	CustomerNumber                    string                                     `json:"customerNumber"`
	ExecutingUserRoles                []Role                                     `json:"executingUserRoles,omitempty"`
	Flags                             []CustomerFlag                             `json:"flags,omitempty"`
	IsAllowedToPlaceOrders            *bool                                      `json:"isAllowedToPlaceOrders,omitempty"`
	IsBanned                          *bool                                      `json:"isBanned,omitempty"`
	IsInDefaultOfPayment              *bool                                      `json:"isInDefaultOfPayment,omitempty"`
	LevelOfUndeliverableDunningNotice *CustomerLevelOfUndeliverableDunningNotice `json:"levelOfUndeliverableDunningNotice,omitempty"`
	MemberCount                       int64                                      `json:"memberCount"`
	Name                              string                                     `json:"name"`
	Owner                             *Contact                                   `json:"owner,omitempty"`
	ProjectCount                      int64                                      `json:"projectCount"`
	VatId                             *string                                    `json:"vatId,omitempty"`
	VatIdValidationState              *CustomerVatIDValidationState              `json:"vatIdValidationState,omitempty"`
}

func (o *Customer) Validate() error {
	if err := func() error {
		if o.ActiveSuspension == nil {
			return nil
		}
		return o.ActiveSuspension.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property activeSuspension: %w", err)
	}
	if err := func() error {
		if o.ExecutingUserRoles == nil {
			return nil
		}
		return func() error {
			for i := range o.ExecutingUserRoles {
				if err := o.ExecutingUserRoles[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property executingUserRoles: %w", err)
	}
	if err := func() error {
		if o.Flags == nil {
			return nil
		}
		return func() error {
			for i := range o.Flags {
				if err := o.Flags[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property flags: %w", err)
	}
	if err := func() error {
		if o.LevelOfUndeliverableDunningNotice == nil {
			return nil
		}
		return o.LevelOfUndeliverableDunningNotice.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property levelOfUndeliverableDunningNotice: %w", err)
	}
	if err := func() error {
		if o.Owner == nil {
			return nil
		}
		return o.Owner.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property owner: %w", err)
	}
	if err := func() error {
		if o.VatIdValidationState == nil {
			return nil
		}
		return o.VatIdValidationState.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property vatIdValidationState: %w", err)
	}
	return nil
}
