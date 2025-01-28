package appv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "faqLink":
//        type: "string"
//        format: "uri"
// required:
//    - "faqLink"
// description: "A BreakingNote is a hint that something serious has changed in the AppVersion containing it, so an automatic update is not possible."

// A BreakingNote is a hint that something serious has changed in the AppVersion containing it, so an automatic update is not possible.
type BreakingNote struct {
	FaqLink string `json:"faqLink"`
}

func (o *BreakingNote) Validate() error {
	return nil
}
