package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "settingsString":
//        type: "string"

//
type DeprecatedUserServicePersonalizedSettingsGetResponse struct {
	SettingsString *string `json:"settingsString,omitempty"`
}

func (o *DeprecatedUserServicePersonalizedSettingsGetResponse) Validate() error {
	return nil
}