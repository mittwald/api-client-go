package project

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "invitationToken":
//        type: "string"
//        description: "Token contained in the invite for authentication."
// description: AcceptProjectInviteRequestBody models the JSON body of a 'project-accept-project-invite' request

// AcceptProjectInviteRequestBody models the JSON body of a 'project-accept-project-invite' request
type AcceptProjectInviteRequestBody struct {
	InvitationToken *string `json:"invitationToken,omitempty"`
}

func (o *AcceptProjectInviteRequestBody) Validate() error {
	return nil
}
