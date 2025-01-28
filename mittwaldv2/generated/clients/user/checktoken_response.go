package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "id":
//        type: "string"
//        format: "uuid"
//    "publicToken":
//        type: "string"
// required:
//    - "id"
//    - "publicToken"

type CheckTokenResponse struct {
	Id          string `json:"id"`
	PublicToken string `json:"publicToken"`
}

func (o *CheckTokenResponse) Validate() error {
	return nil
}
