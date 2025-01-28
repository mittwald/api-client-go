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
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.CombinedACustom"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordMXCustom"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordTXTComponent"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordSRVComponent"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordCNAMEComponent"}
// description: UpdateRecordSetRequestBody models the JSON body of a 'dns-update-record-set' request

type UpdateRecordSetRequestBody struct {
	AlternativeRecordUnset          *dnsv1.RecordUnset
	AlternativeCombinedACustom      *dnsv1.CombinedACustom
	AlternativeRecordMXCustom       *dnsv1.RecordMXCustom
	AlternativeRecordTXTComponent   *dnsv1.RecordTXTComponent
	AlternativeRecordSRVComponent   *dnsv1.RecordSRVComponent
	AlternativeRecordCNAMEComponent *dnsv1.RecordCNAMEComponent
}

func (a *UpdateRecordSetRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeRecordUnset != nil {
		return json.Marshal(a.AlternativeRecordUnset)
	}
	if a.AlternativeCombinedACustom != nil {
		return json.Marshal(a.AlternativeCombinedACustom)
	}
	if a.AlternativeRecordMXCustom != nil {
		return json.Marshal(a.AlternativeRecordMXCustom)
	}
	if a.AlternativeRecordTXTComponent != nil {
		return json.Marshal(a.AlternativeRecordTXTComponent)
	}
	if a.AlternativeRecordSRVComponent != nil {
		return json.Marshal(a.AlternativeRecordSRVComponent)
	}
	if a.AlternativeRecordCNAMEComponent != nil {
		return json.Marshal(a.AlternativeRecordCNAMEComponent)
	}
	return []byte("null"), nil
}

func (a *UpdateRecordSetRequestBody) UnmarshalJSON(input []byte) error {
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

	reader.Reset(input)
	var alternativeRecordMXCustom dnsv1.RecordMXCustom
	if err := dec.Decode(&alternativeRecordMXCustom); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordMXCustom.Validate(); vErr == nil {
			a.AlternativeRecordMXCustom = &alternativeRecordMXCustom
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeRecordTXTComponent dnsv1.RecordTXTComponent
	if err := dec.Decode(&alternativeRecordTXTComponent); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordTXTComponent.Validate(); vErr == nil {
			a.AlternativeRecordTXTComponent = &alternativeRecordTXTComponent
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

	reader.Reset(input)
	var alternativeRecordCNAMEComponent dnsv1.RecordCNAMEComponent
	if err := dec.Decode(&alternativeRecordCNAMEComponent); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordCNAMEComponent.Validate(); vErr == nil {
			a.AlternativeRecordCNAMEComponent = &alternativeRecordCNAMEComponent
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *UpdateRecordSetRequestBody) Validate() error {
	if a.AlternativeRecordUnset != nil {
		return a.AlternativeRecordUnset.Validate()
	}
	if a.AlternativeCombinedACustom != nil {
		return a.AlternativeCombinedACustom.Validate()
	}
	if a.AlternativeRecordMXCustom != nil {
		return a.AlternativeRecordMXCustom.Validate()
	}
	if a.AlternativeRecordTXTComponent != nil {
		return a.AlternativeRecordTXTComponent.Validate()
	}
	if a.AlternativeRecordSRVComponent != nil {
		return a.AlternativeRecordSRVComponent.Validate()
	}
	if a.AlternativeRecordCNAMEComponent != nil {
		return a.AlternativeRecordCNAMEComponent.Validate()
	}
	return errors.New("no alternative set")
}
