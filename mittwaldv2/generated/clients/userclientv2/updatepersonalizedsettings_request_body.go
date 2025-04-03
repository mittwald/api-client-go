package userclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "settingsString":
//        type: "string"
//        format: "base64"
//        example: "eyJvbmJvYXJkaW5nIjp7ImNvbXBsZXRlZCI6W119LCJyZWNlbnRWaXNpdHMiOnsicmVjZW50VmlzaXRzIjp7fSwicmVjZW50VmlzaXRlZFBhdGhzIjp7fX0sImNoYW5nZWxvZ3MiOnsicmVhZElkcyI6WzIsMSwzLDQsNSw2LDgsNyw5XX19"
// required:
//    - "settingsString"
// description: UpdatePersonalizedSettingsRequestBody models the JSON body of a 'user-update-personalized-settings' request

// UpdatePersonalizedSettingsRequestBody models the JSON body of a 'user-update-personalized-settings' request
type UpdatePersonalizedSettingsRequestBody struct {
	SettingsString string `json:"settingsString"`
}

func (o *UpdatePersonalizedSettingsRequestBody) Validate() error {
	return nil
}
