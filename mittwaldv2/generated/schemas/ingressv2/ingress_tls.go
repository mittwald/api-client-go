package ingressv2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.ingress.TlsAcme"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.ingress.TlsCertificate"}

type IngressTLS struct {
	AlternativeTlsAcme        *TlsAcme
	AlternativeTlsCertificate *TlsCertificate
}

func (a *IngressTLS) MarshalJSON() ([]byte, error) {
	if a.AlternativeTlsAcme != nil {
		return json.Marshal(a.AlternativeTlsAcme)
	}
	if a.AlternativeTlsCertificate != nil {
		return json.Marshal(a.AlternativeTlsCertificate)
	}
	return []byte("null"), nil
}

func (a *IngressTLS) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeTlsAcme TlsAcme
	if err := dec.Decode(&alternativeTlsAcme); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeTlsAcme.Validate(); vErr == nil {
			a.AlternativeTlsAcme = &alternativeTlsAcme
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeTlsCertificate TlsCertificate
	if err := dec.Decode(&alternativeTlsCertificate); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeTlsCertificate.Validate(); vErr == nil {
			a.AlternativeTlsCertificate = &alternativeTlsCertificate
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *IngressTLS) Validate() error {
	if a.AlternativeTlsAcme != nil {
		return a.AlternativeTlsAcme.Validate()
	}
	if a.AlternativeTlsCertificate != nil {
		return a.AlternativeTlsCertificate.Validate()
	}
	return errors.New("no alternative set")
}
