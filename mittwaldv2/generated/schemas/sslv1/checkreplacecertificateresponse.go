package sslv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "changes": {"$ref": "#/components/schemas/de.mittwald.v1.ssl.CheckReplaceChanges"}
//    "errors":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.ssl.CertificateError"}
//    "isReplaceable":
//        type: "boolean"
//required:
//    - "isReplaceable"

//
type CheckReplaceCertificateResponse struct {
	Changes       *CheckReplaceChanges `json:"changes,omitempty"`
	Errors        []CertificateError   `json:"errors,omitempty"`
	IsReplaceable bool                 `json:"isReplaceable"`
}

func (o *CheckReplaceCertificateResponse) Validate() error {
	if err := func() error {
		if o.Changes == nil {
			return nil
		}
		return o.Changes.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property changes: %w", err)
	}
	if err := func() error {
		if o.Errors == nil {
			return nil
		}
		return func() error {
			for i := range o.Errors {
				if err := o.Errors[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property errors: %w", err)
	}
	return nil
}
