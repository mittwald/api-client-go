package projectclientv2

import (
	"fmt"
	"time"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/membershipv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//        description: "Time the ProjectMembership should expire at."
//    "role": {"$ref": "#/components/schemas/de.mittwald.v1.membership.ProjectRoles"}
// required:
//    - "role"
// description: UpdateProjectMembershipRequestBody models the JSON body of a 'project-update-project-membership' request

// UpdateProjectMembershipRequestBody models the JSON body of a 'project-update-project-membership' request
type UpdateProjectMembershipRequestBody struct {
	ExpiresAt *time.Time                `json:"expiresAt,omitempty"`
	Role      membershipv2.ProjectRoles `json:"role"`
}

func (o *UpdateProjectMembershipRequestBody) Validate() error {
	if err := o.Role.Validate(); err != nil {
		return fmt.Errorf("invalid property role: %w", err)
	}
	return nil
}
