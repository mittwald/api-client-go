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
// oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordUnset"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordSRVComponent"}
// description: RecordSrvSetDeprecatedRequestBody models the JSON body of a 'dns-record-srv-set-deprecated' request

type RecordSrvSetDeprecatedRequestBody struct {
	AlternativeRecordUnset        *dnsv1.RecordUnset
	AlternativeRecordSRVComponent *dnsv1.RecordSRVComponent
}

func (a *RecordSrvSetDeprecatedRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeRecordUnset != nil {
		return json.Marshal(a.AlternativeRecordUnset)
	}
	if a.AlternativeRecordSRVComponent != nil {
		return json.Marshal(a.AlternativeRecordSRVComponent)
	}
	return []byte("null"), nil
}

func (a *RecordSrvSetDeprecatedRequestBody) UnmarshalJSON(input []byte) error {
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
	var alternativeRecordSRVComponent dnsv1.RecordSRVComponent
	if err := dec.Decode(&alternativeRecordSRVComponent); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordSRVComponent.Validate(); vErr == nil {
			a.AlternativeRecordSRVComponent = &alternativeRecordSRVComponent
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *RecordSrvSetDeprecatedRequestBody) Validate() error {
	if a.AlternativeRecordUnset != nil {
		return a.AlternativeRecordUnset.Validate()
	}
	if a.AlternativeRecordSRVComponent != nil {
		return a.AlternativeRecordSRVComponent.Validate()
	}
	return errors.New("no alternative set")
}
