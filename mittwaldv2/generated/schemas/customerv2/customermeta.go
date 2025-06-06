package customerv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "avatarRefId":
//        type: "string"
//    "customerId":
//        type: "string"
//    "name":
//        type: "string"
// required:
//    - "customerId"
//    - "name"

type CustomerMeta struct {
	AvatarRefId *string `json:"avatarRefId,omitempty"`
	CustomerId  string  `json:"customerId"`
	Name        string  `json:"name"`
}

func (o *CustomerMeta) Validate() error {
	return nil
}
