package mail

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
//    - type: "object"
//      properties:
//        "blacklist":
//            type: "array"
//            items:
//                type: "string"
//                format: "idn-email"
//      required:
//        - "blacklist"
//    - type: "object"
//      properties:
//        "whitelist":
//            type: "array"
//            items:
//                type: "string"
//                format: "idn-email"
//      required:
//        - "whitelist"

type DeprecatedUpdateProjectMailSettingRequestBody struct {
	AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1 *DeprecatedUpdateProjectMailSettingRequestBodyAlternative1
	AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2 *DeprecatedUpdateProjectMailSettingRequestBodyAlternative2
}

func (a *DeprecatedUpdateProjectMailSettingRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1 != nil {
		return json.Marshal(a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1)
	}
	if a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2 != nil {
		return json.Marshal(a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2)
	}
	return []byte("null"), nil
}

func (a *DeprecatedUpdateProjectMailSettingRequestBody) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1 DeprecatedUpdateProjectMailSettingRequestBodyAlternative1
	if err := dec.Decode(&alternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1); err == nil {
		//subtype: *generator.ObjectType
		if vErr := alternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1.Validate(); vErr == nil {
			a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1 = &alternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2 DeprecatedUpdateProjectMailSettingRequestBodyAlternative2
	if err := dec.Decode(&alternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2); err == nil {
		//subtype: *generator.ObjectType
		if vErr := alternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2.Validate(); vErr == nil {
			a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2 = &alternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *DeprecatedUpdateProjectMailSettingRequestBody) Validate() error {
	if a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1 != nil {
		return a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative1.Validate()
	}
	if a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2 != nil {
		return a.AlternativeDeprecatedUpdateProjectMailSettingRequestBodyAlternative2.Validate()
	}
	return errors.New("no alternative set")
}