package membershipv1

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
//    "avatarRefId":
//        type: "string"
//        format: "uuid"
//        description: "Reference to the Project's avatar."
//    "id":
//        type: "string"
//        format: "uuid"
//        description: "ID of the ProjectInvite."
//    "information": {"$ref": "#/components/schemas/de.mittwald.v1.membership.InviteInformation"}
//    "mailAddress":
//        type: "string"
//        format: "email"
//        description: "Mail-address of the user the ProjectInvite is for."
//    "membershipExpiresAt":
//        type: "string"
//        format: "date-time"
//        description: "Time the ProjectMembership should expire at."
//    "message":
//        type: "string"
//        description: "Message contained in the ProjectInvite."
//    "projectDescription":
//        type: "string"
//        description: "Description of the Project the invite is created for."
//    "projectId":
//        type: "string"
//        format: "uuid"
//        description: "ID of the Project the invitation is for."
//    "role": {"$ref": "#/components/schemas/de.mittwald.v1.membership.ProjectRoles"}
//required:
//    - "id"
//    - "projectId"
//    - "mailAddress"
//    - "role"
//    - "accepted"
//    - "information"
//    - "projectDescription"

//
type ProjectInvite struct {
	AvatarRefId         *uuid.UUID        `json:"avatarRefId,omitempty"`
	Id                  uuid.UUID         `json:"id"`
	Information         InviteInformation `json:"information"`
	MailAddress         string            `json:"mailAddress"`
	MembershipExpiresAt *time.Time        `json:"membershipExpiresAt,omitempty"`
	Message             *string           `json:"message,omitempty"`
	ProjectDescription  string            `json:"projectDescription"`
	ProjectId           uuid.UUID         `json:"projectId"`
	Role                ProjectRoles      `json:"role"`
}

func (o *ProjectInvite) Validate() error {
	if err := o.Information.Validate(); err != nil {
		return fmt.Errorf("invalid property information: %w", err)
	}
	if err := o.Role.Validate(); err != nil {
		return fmt.Errorf("invalid property role: %w", err)
	}
	return nil
}
