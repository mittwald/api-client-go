package databaseclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "id":
//        type: "string"
//        format: "uuid"
//    "userId":
//        type: "string"
//        format: "uuid"
// required:
//    - "id"
//    - "userId"

type CreateMysqlDatabaseResponse struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
}

func (o *CreateMysqlDatabaseResponse) Validate() error {
	return nil
}
