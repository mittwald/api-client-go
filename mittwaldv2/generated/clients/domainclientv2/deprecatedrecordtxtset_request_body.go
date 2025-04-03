package domainclientv2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/dnsv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordUnset"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordTXTComponent"}
// description: DeprecatedRecordTxtSetRequestBody models the JSON body of a 'deprecated-dns-record-txt-set' request

type DeprecatedRecordTxtSetRequestBody struct {
	AlternativeRecordUnset        *dnsv2.RecordUnset
	AlternativeRecordTXTComponent *dnsv2.RecordTXTComponent
}

func (a *DeprecatedRecordTxtSetRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeRecordUnset != nil {
		return json.Marshal(a.AlternativeRecordUnset)
	}
	if a.AlternativeRecordTXTComponent != nil {
		return json.Marshal(a.AlternativeRecordTXTComponent)
	}
	return []byte("null"), nil
}

func (a *DeprecatedRecordTxtSetRequestBody) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeRecordUnset dnsv2.RecordUnset
	if err := dec.Decode(&alternativeRecordUnset); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordUnset.Validate(); vErr == nil {
			a.AlternativeRecordUnset = &alternativeRecordUnset
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeRecordTXTComponent dnsv2.RecordTXTComponent
	if err := dec.Decode(&alternativeRecordTXTComponent); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordTXTComponent.Validate(); vErr == nil {
			a.AlternativeRecordTXTComponent = &alternativeRecordTXTComponent
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *DeprecatedRecordTxtSetRequestBody) Validate() error {
	if a.AlternativeRecordUnset != nil {
		return a.AlternativeRecordUnset.Validate()
	}
	if a.AlternativeRecordTXTComponent != nil {
		return a.AlternativeRecordTXTComponent.Validate()
	}
	return errors.New("no alternative set")
}
