package marketplacev1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.UrlFrontendFragment"}

type FrontendFragment struct {
	AlternativeUrlFrontendFragment *UrlFrontendFragment
}

func (a *FrontendFragment) MarshalJSON() ([]byte, error) {
	if a.AlternativeUrlFrontendFragment != nil {
		return json.Marshal(a.AlternativeUrlFrontendFragment)
	}
	return []byte("null"), nil
}

func (a *FrontendFragment) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeUrlFrontendFragment UrlFrontendFragment
	if err := dec.Decode(&alternativeUrlFrontendFragment); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeUrlFrontendFragment.Validate(); vErr == nil {
			a.AlternativeUrlFrontendFragment = &alternativeUrlFrontendFragment
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *FrontendFragment) Validate() error {
	if a.AlternativeUrlFrontendFragment != nil {
		return a.AlternativeUrlFrontendFragment.Validate()
	}
	return errors.New("no alternative set")
}
