package varnishv1

import "time"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "accessToken":
//        type: "string"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "accessToken"
//    - "expiresAt"

//
type SoftwareConfigFileAccessToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

func (o *SoftwareConfigFileAccessToken) Validate() error {
	return nil
}