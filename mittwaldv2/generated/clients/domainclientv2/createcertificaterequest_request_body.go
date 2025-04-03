package domainclientv2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/sslv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.ssl.CertificateRequestCreateRequest"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.ssl.CertificateRequestCreateWithCSRRequest"}
// description: CreateCertificateRequestRequestBody models the JSON body of a 'ssl-create-certificate-request' request

type CreateCertificateRequestRequestBody struct {
	AlternativeCertificateRequestCreateRequest        *sslv2.CertificateRequestCreateRequest
	AlternativeCertificateRequestCreateWithCSRRequest *sslv2.CertificateRequestCreateWithCSRRequest
}

func (a *CreateCertificateRequestRequestBody) MarshalJSON() ([]byte, error) {
	if a.AlternativeCertificateRequestCreateRequest != nil {
		return json.Marshal(a.AlternativeCertificateRequestCreateRequest)
	}
	if a.AlternativeCertificateRequestCreateWithCSRRequest != nil {
		return json.Marshal(a.AlternativeCertificateRequestCreateWithCSRRequest)
	}
	return []byte("null"), nil
}

func (a *CreateCertificateRequestRequestBody) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeCertificateRequestCreateRequest sslv2.CertificateRequestCreateRequest
	if err := dec.Decode(&alternativeCertificateRequestCreateRequest); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeCertificateRequestCreateRequest.Validate(); vErr == nil {
			a.AlternativeCertificateRequestCreateRequest = &alternativeCertificateRequestCreateRequest
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeCertificateRequestCreateWithCSRRequest sslv2.CertificateRequestCreateWithCSRRequest
	if err := dec.Decode(&alternativeCertificateRequestCreateWithCSRRequest); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeCertificateRequestCreateWithCSRRequest.Validate(); vErr == nil {
			a.AlternativeCertificateRequestCreateWithCSRRequest = &alternativeCertificateRequestCreateWithCSRRequest
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *CreateCertificateRequestRequestBody) Validate() error {
	if a.AlternativeCertificateRequestCreateRequest != nil {
		return a.AlternativeCertificateRequestCreateRequest.Validate()
	}
	if a.AlternativeCertificateRequestCreateWithCSRRequest != nil {
		return a.AlternativeCertificateRequestCreateWithCSRRequest.Validate()
	}
	return errors.New("no alternative set")
}
