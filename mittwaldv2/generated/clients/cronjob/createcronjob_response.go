package cronjob

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "id":
//        type: "string"
//        format: "uuid"
// required:
//    - "id"

type CreateCronjobResponse struct {
	Id string `json:"id"`
}

func (o *CreateCronjobResponse) Validate() error {
	return nil
}
