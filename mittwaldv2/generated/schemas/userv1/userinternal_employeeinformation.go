package userv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "department":
//        type: "string"
//        example: "Kundenservice"
//required:
//    - "department"
//description: "Additional information about mittwald employees."

//Additional information about mittwald employees.
type UserInternalEmployeeInformation struct {
	Department string `json:"department"`
}

func (o *UserInternalEmployeeInformation) Validate() error {
	return nil
}