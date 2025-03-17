package marketplaceclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "description":
//        type: "string"
//    "humanReadableName":
//        type: "string"
//    "redirectUris":
//        type: "array"
//        items:
//            type: "string"
//    "scopes":
//        type: "array"
//        items:
//            type: "string"
// description: PatchOauthClientRequestBody models the JSON body of a 'contributor-patch-oauth-client' request

// PatchOauthClientRequestBody models the JSON body of a 'contributor-patch-oauth-client' request
type PatchOauthClientRequestBody struct {
	Description       *string  `json:"description,omitempty"`
	HumanReadableName *string  `json:"humanReadableName,omitempty"`
	RedirectUris      []string `json:"redirectUris,omitempty"`
	Scopes            []string `json:"scopes,omitempty"`
}

func (o *PatchOauthClientRequestBody) Validate() error {
	if err := func() error {
		if o.RedirectUris == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property redirectUris: %w", err)
	}
	if err := func() error {
		if o.Scopes == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property scopes: %w", err)
	}
	return nil
}
