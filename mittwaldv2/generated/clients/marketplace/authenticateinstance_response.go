package marketplace

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "expiry":
//        type: "string"
//        format: "date-time"
//    "publicToken":
//        type: "string"
//        description: "Set this in the 'x-access-token' header when performing domain actions."
//required:
//    - "publicToken"
//    - "expiry"

type AuthenticateInstanceResponse struct {
	Expiry      time.Time `json:"expiry"`
	PublicToken string    `json:"publicToken"`
}

func (o *AuthenticateInstanceResponse) Validate() error {
	return nil
}
