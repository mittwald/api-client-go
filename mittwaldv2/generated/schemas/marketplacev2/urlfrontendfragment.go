package marketplacev2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "url":
//        type: "string"
// required:
//    - "url"

type UrlFrontendFragment struct {
	Url string `json:"url"`
}

func (o *UrlFrontendFragment) Validate() error {
	return nil
}
