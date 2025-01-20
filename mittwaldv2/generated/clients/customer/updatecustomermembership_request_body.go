package customer

import (
	"fmt"
	"time"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/membershipv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//        description: "Time the CustomerMembership should expire at."
//    "role": {"$ref": "#/components/schemas/de.mittwald.v1.membership.CustomerRoles"}
//required:
//    - "role"

//
type UpdateCustomerMembershipRequestBody struct {
	ExpiresAt *time.Time                 `json:"expiresAt,omitempty"`
	Role      membershipv1.CustomerRoles `json:"role"`
}

func (o *UpdateCustomerMembershipRequestBody) Validate() error {
	if err := o.Role.Validate(); err != nil {
		return fmt.Errorf("invalid property role: %w", err)
	}
	return nil
}