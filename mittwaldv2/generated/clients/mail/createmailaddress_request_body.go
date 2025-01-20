package mail

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/mailv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.mail.CreateForwardAddress"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.mail.CreateMailAddress"}

type CreateMailAddressRequestBody struct {
	AlternativeCreateForwardAddress *mailv1.CreateForwardAddress
	AlternativeCreateMailAddress    *mailv1.CreateMailAddress
}

func (a *CreateMailAddressRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeCreateForwardAddress != nil {
		return json.Marshal(a.AlternativeCreateForwardAddress)
	}
	if a.AlternativeCreateMailAddress != nil {
		return json.Marshal(a.AlternativeCreateMailAddress)
	}
	return []byte("null"), nil
}

func (a *CreateMailAddressRequestBody) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeCreateForwardAddress mailv1.CreateForwardAddress
	if err := dec.Decode(&alternativeCreateForwardAddress); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeCreateForwardAddress.Validate(); vErr == nil {
			a.AlternativeCreateForwardAddress = &alternativeCreateForwardAddress
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeCreateMailAddress mailv1.CreateMailAddress
	if err := dec.Decode(&alternativeCreateMailAddress); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeCreateMailAddress.Validate(); vErr == nil {
			a.AlternativeCreateMailAddress = &alternativeCreateMailAddress
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *CreateMailAddressRequestBody) Validate() error {
	if a.AlternativeCreateForwardAddress != nil {
		return a.AlternativeCreateForwardAddress.Validate()
	}
	if a.AlternativeCreateMailAddress != nil {
		return a.AlternativeCreateMailAddress.Validate()
	}
	return errors.New("no alternative set")
}
