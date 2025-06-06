package contractclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "orderData":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.ProjectHostingOrderPreview"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.ServerOrderPreview"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.DomainOrderPreview"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.ExternalCertificateOrderPreview"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.LeadFyndrOrderPreview"}
//    "orderType":
//        type: "string"
//        enum:
//            - "domain"
//            - "projectHosting"
//            - "server"
//            - "externalCertificate"
//            - "leadFyndr"
// description: PreviewOrderRequestBody models the JSON body of a 'order-preview-order' request

// PreviewOrderRequestBody models the JSON body of a 'order-preview-order' request
type PreviewOrderRequestBody struct {
	OrderData *PreviewOrderRequestBodyOrderData `json:"orderData,omitempty"`
	OrderType *PreviewOrderRequestBodyOrderType `json:"orderType,omitempty"`
}

func (o *PreviewOrderRequestBody) Validate() error {
	if err := func() error {
		if o.OrderData == nil {
			return nil
		}
		return o.OrderData.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property orderData: %w", err)
	}
	if err := func() error {
		if o.OrderType == nil {
			return nil
		}
		return o.OrderType.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property orderType: %w", err)
	}
	return nil
}
