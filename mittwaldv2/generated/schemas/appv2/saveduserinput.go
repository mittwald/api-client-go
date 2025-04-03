package appv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "name":
//        type: "string"
//    "value":
//        type: "string"
// required:
//    - "name"
//    - "value"
// description: "A SavedUserInput is an entered value for a desired UserInput of an AppVersion or SystemSoftwareVersion."

// A SavedUserInput is an entered value for a desired UserInput of an AppVersion or SystemSoftwareVersion.
type SavedUserInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *SavedUserInput) Validate() error {
	return nil
}
