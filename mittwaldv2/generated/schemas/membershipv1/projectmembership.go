package membershipv1

import (
	"fmt"
	"time"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "email":
//        type: "string"
//        format: "email"
//        description: "Email used by the invited user."
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//        description: "Time the ProjectMembership should expire at."
//    "id":
//        type: "string"
//        format: "uuid"
//        description: "ID of the ProjectMembership."
//    "inherited":
//        type: "boolean"
//        description: "Whether the ProjectMembership was inherited from a CustomerMembership."
//    "inviteId":
//        type: "string"
//        format: "uuid"
//        description: "ID of the ProjectInvite the membership was created from."
//    "memberSince":
//        type: "string"
//        format: "date-time"
//        description: "Date the projectMembership was created at."
//    "mfa":
//        type: "boolean"
//        description: "MFA activated by the user."
//    "projectId":
//        type: "string"
//        format: "uuid"
//        description: "ID of the Project the membership is for."
//    "role": {"$ref": "#/components/schemas/de.mittwald.v1.membership.ProjectRoles"}
//    "userId":
//        type: "string"
//        format: "uuid"
//        description: "ID of the user the ProjectMembership is for."
// required:
//    - "id"
//    - "projectId"
//    - "userId"
//    - "role"
//    - "inherited"
//    - "email"
//    - "mfa"

type ProjectMembership struct {
	Email       string       `json:"email"`
	ExpiresAt   *time.Time   `json:"expiresAt,omitempty"`
	Id          string       `json:"id"`
	Inherited   bool         `json:"inherited"`
	InviteId    *string      `json:"inviteId,omitempty"`
	MemberSince *time.Time   `json:"memberSince,omitempty"`
	Mfa         bool         `json:"mfa"`
	ProjectId   string       `json:"projectId"`
	Role        ProjectRoles `json:"role"`
	UserId      string       `json:"userId"`
}

func (o *ProjectMembership) Validate() error {
	if err := o.Role.Validate(); err != nil {
		return fmt.Errorf("invalid property role: %w", err)
	}
	return nil
}
