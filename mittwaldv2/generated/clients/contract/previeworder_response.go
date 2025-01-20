package contract

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.order.HostingOrderPreviewResponse"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.order.DomainOrderPreviewResponse"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.order.ExternalCertificateOrderPreviewResponse"}

type PreviewOrderResponse struct {
	AlternativeHostingOrderPreviewResponse             *orderv1.HostingOrderPreviewResponse
	AlternativeDomainOrderPreviewResponse              *orderv1.DomainOrderPreviewResponse
	AlternativeExternalCertificateOrderPreviewResponse *orderv1.ExternalCertificateOrderPreviewResponse
}

func (a *PreviewOrderResponse) MarshalJSON() ([]byte, error) {
	if a.AlternativeHostingOrderPreviewResponse != nil {
		return json.Marshal(a.AlternativeHostingOrderPreviewResponse)
	}
	if a.AlternativeDomainOrderPreviewResponse != nil {
		return json.Marshal(a.AlternativeDomainOrderPreviewResponse)
	}
	if a.AlternativeExternalCertificateOrderPreviewResponse != nil {
		return json.Marshal(a.AlternativeExternalCertificateOrderPreviewResponse)
	}
	return []byte("null"), nil
}

func (a *PreviewOrderResponse) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeHostingOrderPreviewResponse orderv1.HostingOrderPreviewResponse
	if err := dec.Decode(&alternativeHostingOrderPreviewResponse); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeHostingOrderPreviewResponse.Validate(); vErr == nil {
			a.AlternativeHostingOrderPreviewResponse = &alternativeHostingOrderPreviewResponse
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeDomainOrderPreviewResponse orderv1.DomainOrderPreviewResponse
	if err := dec.Decode(&alternativeDomainOrderPreviewResponse); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeDomainOrderPreviewResponse.Validate(); vErr == nil {
			a.AlternativeDomainOrderPreviewResponse = &alternativeDomainOrderPreviewResponse
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeExternalCertificateOrderPreviewResponse orderv1.ExternalCertificateOrderPreviewResponse
	if err := dec.Decode(&alternativeExternalCertificateOrderPreviewResponse); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeExternalCertificateOrderPreviewResponse.Validate(); vErr == nil {
			a.AlternativeExternalCertificateOrderPreviewResponse = &alternativeExternalCertificateOrderPreviewResponse
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *PreviewOrderResponse) Validate() error {
	if a.AlternativeHostingOrderPreviewResponse != nil {
		return a.AlternativeHostingOrderPreviewResponse.Validate()
	}
	if a.AlternativeDomainOrderPreviewResponse != nil {
		return a.AlternativeDomainOrderPreviewResponse.Validate()
	}
	if a.AlternativeExternalCertificateOrderPreviewResponse != nil {
		return a.AlternativeExternalCertificateOrderPreviewResponse.Validate()
	}
	return errors.New("no alternative set")
}
