package contractclientv2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv2"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.order.ProjectHostingOrderPreview"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.order.ServerOrderPreview"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.order.DomainOrderPreview"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.order.ExternalCertificateOrderPreview"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.order.LeadFyndrOrderPreview"}

type PreviewOrderRequestBodyOrderData struct {
	AlternativeProjectHostingOrderPreview      *orderv2.ProjectHostingOrderPreview
	AlternativeServerOrderPreview              *orderv2.ServerOrderPreview
	AlternativeDomainOrderPreview              *orderv2.DomainOrderPreview
	AlternativeExternalCertificateOrderPreview *orderv2.ExternalCertificateOrderPreview
	AlternativeLeadFyndrOrderPreview           *orderv2.LeadFyndrOrderPreview
}

func (a *PreviewOrderRequestBodyOrderData) MarshalJSON() ([]byte, error) {
	if a.AlternativeProjectHostingOrderPreview != nil {
		return json.Marshal(a.AlternativeProjectHostingOrderPreview)
	}
	if a.AlternativeServerOrderPreview != nil {
		return json.Marshal(a.AlternativeServerOrderPreview)
	}
	if a.AlternativeDomainOrderPreview != nil {
		return json.Marshal(a.AlternativeDomainOrderPreview)
	}
	if a.AlternativeExternalCertificateOrderPreview != nil {
		return json.Marshal(a.AlternativeExternalCertificateOrderPreview)
	}
	if a.AlternativeLeadFyndrOrderPreview != nil {
		return json.Marshal(a.AlternativeLeadFyndrOrderPreview)
	}
	return []byte("null"), nil
}

func (a *PreviewOrderRequestBodyOrderData) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeProjectHostingOrderPreview orderv2.ProjectHostingOrderPreview
	if err := dec.Decode(&alternativeProjectHostingOrderPreview); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeProjectHostingOrderPreview.Validate(); vErr == nil {
			a.AlternativeProjectHostingOrderPreview = &alternativeProjectHostingOrderPreview
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeServerOrderPreview orderv2.ServerOrderPreview
	if err := dec.Decode(&alternativeServerOrderPreview); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeServerOrderPreview.Validate(); vErr == nil {
			a.AlternativeServerOrderPreview = &alternativeServerOrderPreview
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeDomainOrderPreview orderv2.DomainOrderPreview
	if err := dec.Decode(&alternativeDomainOrderPreview); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeDomainOrderPreview.Validate(); vErr == nil {
			a.AlternativeDomainOrderPreview = &alternativeDomainOrderPreview
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeExternalCertificateOrderPreview orderv2.ExternalCertificateOrderPreview
	if err := dec.Decode(&alternativeExternalCertificateOrderPreview); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeExternalCertificateOrderPreview.Validate(); vErr == nil {
			a.AlternativeExternalCertificateOrderPreview = &alternativeExternalCertificateOrderPreview
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeLeadFyndrOrderPreview orderv2.LeadFyndrOrderPreview
	if err := dec.Decode(&alternativeLeadFyndrOrderPreview); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeLeadFyndrOrderPreview.Validate(); vErr == nil {
			a.AlternativeLeadFyndrOrderPreview = &alternativeLeadFyndrOrderPreview
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *PreviewOrderRequestBodyOrderData) Validate() error {
	if a.AlternativeProjectHostingOrderPreview != nil {
		return a.AlternativeProjectHostingOrderPreview.Validate()
	}
	if a.AlternativeServerOrderPreview != nil {
		return a.AlternativeServerOrderPreview.Validate()
	}
	if a.AlternativeDomainOrderPreview != nil {
		return a.AlternativeDomainOrderPreview.Validate()
	}
	if a.AlternativeExternalCertificateOrderPreview != nil {
		return a.AlternativeExternalCertificateOrderPreview.Validate()
	}
	if a.AlternativeLeadFyndrOrderPreview != nil {
		return a.AlternativeLeadFyndrOrderPreview.Validate()
	}
	return errors.New("no alternative set")
}
