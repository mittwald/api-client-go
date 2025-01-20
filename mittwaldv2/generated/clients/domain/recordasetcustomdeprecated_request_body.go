package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/dnsv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordUnset"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.CombinedACustom"}

type RecordASetCustomDeprecatedRequestBody struct {
	AlternativeRecordUnset     *dnsv1.RecordUnset
	AlternativeCombinedACustom *dnsv1.CombinedACustom
}

func (a *RecordASetCustomDeprecatedRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeRecordUnset != nil {
		return json.Marshal(a.AlternativeRecordUnset)
	}
	if a.AlternativeCombinedACustom != nil {
		return json.Marshal(a.AlternativeCombinedACustom)
	}
	return []byte("null"), nil
}

func (a *RecordASetCustomDeprecatedRequestBody) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeRecordUnset dnsv1.RecordUnset
	if err := dec.Decode(&alternativeRecordUnset); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordUnset.Validate(); vErr == nil {
			a.AlternativeRecordUnset = &alternativeRecordUnset
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeCombinedACustom dnsv1.CombinedACustom
	if err := dec.Decode(&alternativeCombinedACustom); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeCombinedACustom.Validate(); vErr == nil {
			a.AlternativeCombinedACustom = &alternativeCombinedACustom
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *RecordASetCustomDeprecatedRequestBody) Validate() error {
	if a.AlternativeRecordUnset != nil {
		return a.AlternativeRecordUnset.Validate()
	}
	if a.AlternativeCombinedACustom != nil {
		return a.AlternativeCombinedACustom.Validate()
	}
	return errors.New("no alternative set")
}
