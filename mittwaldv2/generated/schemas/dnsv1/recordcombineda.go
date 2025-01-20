package dnsv1

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
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordUnset"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.CombinedACustom"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.dns.CombinedAManaged"}

type RecordCombinedA struct {
	AlternativeRecordUnset      *RecordUnset
	AlternativeCombinedACustom  *CombinedACustom
	AlternativeCombinedAManaged *CombinedAManaged
}

func (a *RecordCombinedA) MarshalJSON() ([]byte, error) {
	if a.AlternativeRecordUnset != nil {
		return json.Marshal(a.AlternativeRecordUnset)
	}
	if a.AlternativeCombinedACustom != nil {
		return json.Marshal(a.AlternativeCombinedACustom)
	}
	if a.AlternativeCombinedAManaged != nil {
		return json.Marshal(a.AlternativeCombinedAManaged)
	}
	return []byte("null"), nil
}

func (a *RecordCombinedA) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeRecordUnset RecordUnset
	if err := dec.Decode(&alternativeRecordUnset); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRecordUnset.Validate(); vErr == nil {
			a.AlternativeRecordUnset = &alternativeRecordUnset
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeCombinedACustom CombinedACustom
	if err := dec.Decode(&alternativeCombinedACustom); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeCombinedACustom.Validate(); vErr == nil {
			a.AlternativeCombinedACustom = &alternativeCombinedACustom
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeCombinedAManaged CombinedAManaged
	if err := dec.Decode(&alternativeCombinedAManaged); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeCombinedAManaged.Validate(); vErr == nil {
			a.AlternativeCombinedAManaged = &alternativeCombinedAManaged
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *RecordCombinedA) Validate() error {
	if a.AlternativeRecordUnset != nil {
		return a.AlternativeRecordUnset.Validate()
	}
	if a.AlternativeCombinedACustom != nil {
		return a.AlternativeCombinedACustom.Validate()
	}
	if a.AlternativeCombinedAManaged != nil {
		return a.AlternativeCombinedAManaged.Validate()
	}
	return errors.New("no alternative set")
}
