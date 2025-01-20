package sshuserv1

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
//        "password":
//            type: "string"
//            maxLength: 72
//      required:
//        - "password"
//    - type: "object"
//      properties:
//        "publicKeys":
//            type: "array"
//            items: {"$ref": "#/components/schemas/de.mittwald.v1.sshuser.PublicKey"}
//      required:
//        - "publicKeys"
//description: "Method of authentication for an SFTPUser or SSHUser. Can be password or public-keys."

type Authentication struct {
	AlternativeAuthenticationAlternative1 *AuthenticationAlternative1
	AlternativeAuthenticationAlternative2 *AuthenticationAlternative2
}

func (a *Authentication) MarshalJSON() ([]byte, error) {
	if a.AlternativeAuthenticationAlternative1 != nil {
		return json.Marshal(a.AlternativeAuthenticationAlternative1)
	}
	if a.AlternativeAuthenticationAlternative2 != nil {
		return json.Marshal(a.AlternativeAuthenticationAlternative2)
	}
	return []byte("null"), nil
}

func (a *Authentication) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeAuthenticationAlternative1 AuthenticationAlternative1
	if err := dec.Decode(&alternativeAuthenticationAlternative1); err == nil {
		//subtype: *generator.ObjectType
		if vErr := alternativeAuthenticationAlternative1.Validate(); vErr == nil {
			a.AlternativeAuthenticationAlternative1 = &alternativeAuthenticationAlternative1
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeAuthenticationAlternative2 AuthenticationAlternative2
	if err := dec.Decode(&alternativeAuthenticationAlternative2); err == nil {
		//subtype: *generator.ObjectType
		if vErr := alternativeAuthenticationAlternative2.Validate(); vErr == nil {
			a.AlternativeAuthenticationAlternative2 = &alternativeAuthenticationAlternative2
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *Authentication) Validate() error {
	if a.AlternativeAuthenticationAlternative1 != nil {
		return a.AlternativeAuthenticationAlternative1.Validate()
	}
	if a.AlternativeAuthenticationAlternative2 != nil {
		return a.AlternativeAuthenticationAlternative2.Validate()
	}
	return errors.New("no alternative set")
}