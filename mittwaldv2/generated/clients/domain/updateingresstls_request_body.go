package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// oneOf:
//    - type: "object"
//      properties:
//        "acme":
//            type: "boolean"
//        "isCreated":
//            type: "boolean"
//            description: "Was added by mistake. Never did anything."
//            deprecated: true
//        "requestDeadline":
//            type: "string"
//            format: "date-time"
//            description: "Was added by mistake. Never did anything."
//            deprecated: true
//      required:
//        - "acme"
//      additionalProperties: false
//    - type: "object"
//      properties:
//        "certificateId":
//            type: "string"
//            format: "uuid"
//      required:
//        - "certificateId"
//      additionalProperties: false
// description: UpdateIngressTLSRequestBody models the JSON body of a 'ingress-update-ingress-tls' request

type UpdateIngressTLSRequestBody struct {
	AlternativeUpdateIngressTLSRequestBodyAlternative1 *UpdateIngressTLSRequestBodyAlternative1
	AlternativeUpdateIngressTLSRequestBodyAlternative2 *UpdateIngressTLSRequestBodyAlternative2
}

func (a *UpdateIngressTLSRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeUpdateIngressTLSRequestBodyAlternative1 != nil {
		return json.Marshal(a.AlternativeUpdateIngressTLSRequestBodyAlternative1)
	}
	if a.AlternativeUpdateIngressTLSRequestBodyAlternative2 != nil {
		return json.Marshal(a.AlternativeUpdateIngressTLSRequestBodyAlternative2)
	}
	return []byte("null"), nil
}

func (a *UpdateIngressTLSRequestBody) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeUpdateIngressTLSRequestBodyAlternative1 UpdateIngressTLSRequestBodyAlternative1
	if err := dec.Decode(&alternativeUpdateIngressTLSRequestBodyAlternative1); err == nil {
		//subtype: *generator.ObjectType
		if vErr := alternativeUpdateIngressTLSRequestBodyAlternative1.Validate(); vErr == nil {
			a.AlternativeUpdateIngressTLSRequestBodyAlternative1 = &alternativeUpdateIngressTLSRequestBodyAlternative1
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeUpdateIngressTLSRequestBodyAlternative2 UpdateIngressTLSRequestBodyAlternative2
	if err := dec.Decode(&alternativeUpdateIngressTLSRequestBodyAlternative2); err == nil {
		//subtype: *generator.ObjectType
		if vErr := alternativeUpdateIngressTLSRequestBodyAlternative2.Validate(); vErr == nil {
			a.AlternativeUpdateIngressTLSRequestBodyAlternative2 = &alternativeUpdateIngressTLSRequestBodyAlternative2
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *UpdateIngressTLSRequestBody) Validate() error {
	if a.AlternativeUpdateIngressTLSRequestBodyAlternative1 != nil {
		return a.AlternativeUpdateIngressTLSRequestBodyAlternative1.Validate()
	}
	if a.AlternativeUpdateIngressTLSRequestBodyAlternative2 != nil {
		return a.AlternativeUpdateIngressTLSRequestBodyAlternative2.Validate()
	}
	return errors.New("no alternative set")
}
