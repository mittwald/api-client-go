package signupv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "browser":
//        type: "string"
//        example: "Safari"
//    "model":
//        type: "string"
//        example: "Apple Macintosh"
//    "os":
//        type: "string"
//        example: "Mac OS"
//    "type":
//        type: "string"
//        example: "Macbook"

//
type DeviceInfo struct {
	Browser *string `json:"browser,omitempty"`
	Model   *string `json:"model,omitempty"`
	Os      *string `json:"os,omitempty"`
	Type    *string `json:"type,omitempty"`
}

func (o *DeviceInfo) Validate() error {
	return nil
}