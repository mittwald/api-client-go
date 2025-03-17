package marketplaceclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "allowedRedirectUris":
//        type: "array"
//        items:
//            type: "string"
//    "allowedScopes":
//        type: "array"
//        items:
//            type: "string"
//    "description":
//        type: "string"
//    "humanReadableName":
//        type: "string"
//    "id":
//        type: "string"
// required:
//    - "id"
//    - "humanReadableName"

type PatchOauthClientResponse struct {
	AllowedRedirectUris []string `json:"allowedRedirectUris,omitempty"`
	AllowedScopes       []string `json:"allowedScopes,omitempty"`
	Description         *string  `json:"description,omitempty"`
	HumanReadableName   string   `json:"humanReadableName"`
	Id                  string   `json:"id"`
}

func (o *PatchOauthClientResponse) Validate() error {
	if err := func() error {
		if o.AllowedRedirectUris == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property allowedRedirectUris: %w", err)
	}
	if err := func() error {
		if o.AllowedScopes == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property allowedScopes: %w", err)
	}
	return nil
}
