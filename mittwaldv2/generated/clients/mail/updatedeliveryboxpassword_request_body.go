package mail

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//properties:
//    "password":
//        type: "string"
//required:
//    - "password"

type UpdateDeliveryBoxPasswordRequestBody struct {
	Password string `json:"password"`
}

func (o *UpdateDeliveryBoxPasswordRequestBody) Validate() error {
	return nil
}
