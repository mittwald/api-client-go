package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "name":
//        type: "string"
//        enum:
//            - "SecondFactorRequired"

//
type AuthenticateAcceptedResponse struct {
	Name *AuthenticateAcceptedResponseName `json:"name,omitempty"`
}

func (o *AuthenticateAcceptedResponse) Validate() error {
	if err := func() error {
		if o.Name == nil {
			return nil
		}
		return o.Name.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property name: %w", err)
	}
	return nil
}