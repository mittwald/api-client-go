package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "settingsString":
//        type: "string"
//required:
//    - "settingsString"

//
type DeprecatedServicePersonalizedSettingsUpdateRequestBody struct {
	SettingsString string `json:"settingsString"`
}

func (o *DeprecatedServicePersonalizedSettingsUpdateRequestBody) Validate() error {
	return nil
}