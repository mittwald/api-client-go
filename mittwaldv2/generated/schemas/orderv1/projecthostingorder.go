package orderv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "customerId":
//        type: "string"
//        example: "0f5ec9cd-1b1b-4850-9061-fcebe765c62d"
//    "description":
//        type: "string"
//        example: "My first project"
//    "diskspaceInGiB":
//        type: "number"
//        example: 20
//    "promotionCode":
//        type: "string"
//        example: "123456"
//    "recommendationCode":
//        type: "string"
//        example: "mp-123456"
//    "spec":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.MachineTypeSpec"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.HardwareSpec"}
//    "useFreeTrial":
//        type: "boolean"
// required:
//    - "customerId"
//    - "diskspaceInGiB"
//    - "description"
//    - "spec"

type ProjectHostingOrder struct {
	CustomerId         string                  `json:"customerId"`
	Description        string                  `json:"description"`
	DiskspaceInGiB     float64                 `json:"diskspaceInGiB"`
	PromotionCode      *string                 `json:"promotionCode,omitempty"`
	RecommendationCode *string                 `json:"recommendationCode,omitempty"`
	Spec               ProjectHostingOrderSpec `json:"spec"`
	UseFreeTrial       *bool                   `json:"useFreeTrial,omitempty"`
}

func (o *ProjectHostingOrder) Validate() error {
	if err := o.Spec.Validate(); err != nil {
		return fmt.Errorf("invalid property spec: %w", err)
	}
	return nil
}
