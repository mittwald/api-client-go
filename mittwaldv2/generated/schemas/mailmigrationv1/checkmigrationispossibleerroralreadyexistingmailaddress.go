package mailmigrationv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "address":
//        type: "string"
//required:
//    - "address"

type CheckMigrationIsPossibleErrorAlreadyExistingMailAddress struct {
	Address string `json:"address"`
}

func (o *CheckMigrationIsPossibleErrorAlreadyExistingMailAddress) Validate() error {
	return nil
}
