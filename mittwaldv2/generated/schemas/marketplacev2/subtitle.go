package marketplacev2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "de":
//        type: "string"
//        maxLength: 40
//        example: "Ping deine App an"
//    "en":
//        type: "string"
//        maxLength: 40
//        example: "Ping your app"
// required:
//    - "de"
// description: "A few words to promote your Extension."

// A few words to promote your Extension.
type SubTitle struct {
	De string  `json:"de"`
	En *string `json:"en,omitempty"`
}

func (o *SubTitle) Validate() error {
	return nil
}
