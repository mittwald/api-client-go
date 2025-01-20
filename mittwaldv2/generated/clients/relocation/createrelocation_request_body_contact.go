package relocation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "email":
//        type: "string"
//        minLength: 1
//        format: "email"
//    "firstName":
//        type: "string"
//        minLength: 1
//    "lastName":
//        type: "string"
//        minLength: 1
//    "phoneNumber":
//        type: "string"
//        pattern: "|^\\+([0-9]{2,3}|1)-[0-9]{2,5}-[0-9]+$"
//required:
//    - "firstName"
//    - "lastName"
//    - "email"

//
type CreateRelocationRequestBodyContact struct {
	Email       string  `json:"email"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

func (o *CreateRelocationRequestBodyContact) Validate() error {
	return nil
}
