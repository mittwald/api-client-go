package userclientv2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// null

type ChangePasswordResponse struct {
	AlternativeChangePasswordOKResponse       *ChangePasswordOKResponse
	AlternativeChangePasswordAcceptedResponse *any
}

func (a *ChangePasswordResponse) MarshalJSON() ([]byte, error) {
	if a.AlternativeChangePasswordOKResponse != nil {
		return json.Marshal(a.AlternativeChangePasswordOKResponse)
	}
	if a.AlternativeChangePasswordAcceptedResponse != nil {
		return json.Marshal(a.AlternativeChangePasswordAcceptedResponse)
	}
	return []byte("null"), nil
}

func (a *ChangePasswordResponse) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeChangePasswordOKResponse ChangePasswordOKResponse
	if err := dec.Decode(&alternativeChangePasswordOKResponse); err == nil {
		//subtype: *generator.ObjectType
		if vErr := alternativeChangePasswordOKResponse.Validate(); vErr == nil {
			a.AlternativeChangePasswordOKResponse = &alternativeChangePasswordOKResponse
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeChangePasswordAcceptedResponse any
	if err := dec.Decode(&alternativeChangePasswordAcceptedResponse); err == nil {
		a.AlternativeChangePasswordAcceptedResponse = &alternativeChangePasswordAcceptedResponse
		decodedAtLeastOnce = true
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *ChangePasswordResponse) Validate() error {
	if a.AlternativeChangePasswordOKResponse != nil {
		return a.AlternativeChangePasswordOKResponse.Validate()
	}
	// The AlternativeChangePasswordAcceptedResponse subtype does not implement validation, so we consider being non-nil as valid
	if a.AlternativeChangePasswordAcceptedResponse != nil {
		return nil
	}
	return errors.New("no alternative set")
}
