package ingressv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "useDefaultPage":
//        type: "boolean"
//required:
//    - "useDefaultPage"

//
type TargetUseDefaultPage struct {
	UseDefaultPage bool `json:"useDefaultPage"`
}

func (o *TargetUseDefaultPage) Validate() error {
	return nil
}