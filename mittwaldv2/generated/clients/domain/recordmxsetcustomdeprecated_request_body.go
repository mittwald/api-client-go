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
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordMXCustom"}

type RecordMxSetCustomDeprecatedRequestBody struct {
	AlternativeRecordUnset    *dnsv1.RecordUnset
	AlternativeRecordMXCustom *dnsv1.RecordMXCustom
}

func (a *RecordMxSetCustomDeprecatedRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeRecordUnset != nil {
		return json.Marshal(a.AlternativeRecordUnset)
	}
	if a.AlternativeRecordMXCustom != nil {
		return json.Marshal(a.AlternativeRecordMXCustom)
	}
	return []byte("null"), nil
}

func (a *RecordMxSetCustomDeprecatedRequestBody) UnmarshalJSON(input []byte) error {
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
	var alternativeRecordMXCustom dnsv1.RecordMXCustom
	if err := dec.Decode(&alternativeRecordMXCustom); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordMXCustom.Validate(); vErr == nil {
			a.AlternativeRecordMXCustom = &alternativeRecordMXCustom
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *RecordMxSetCustomDeprecatedRequestBody) Validate() error {
	if a.AlternativeRecordUnset != nil {
		return a.AlternativeRecordUnset.Validate()
	}
	if a.AlternativeRecordMXCustom != nil {
		return a.AlternativeRecordMXCustom.Validate()
	}
	return errors.New("no alternative set")
}