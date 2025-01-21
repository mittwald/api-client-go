package customer

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "invitationToken":
//        type: "string"
//        description: "Token contained in the invite for authentication."

type AcceptCustomerInviteRequestBody struct {
	InvitationToken *string `json:"invitationToken,omitempty"`
}

func (o *AcceptCustomerInviteRequestBody) Validate() error {
	return nil
}
