package conversation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// properties:
//    "messageContent":
//        type: "string"
// description: UpdateMessageRequestBody models the JSON body of a 'conversation-update-message' request

// UpdateMessageRequestBody models the JSON body of a 'conversation-update-message' request
type UpdateMessageRequestBody struct {
	MessageContent *string `json:"messageContent,omitempty"`
}

func (o *UpdateMessageRequestBody) Validate() error {
	return nil
}
