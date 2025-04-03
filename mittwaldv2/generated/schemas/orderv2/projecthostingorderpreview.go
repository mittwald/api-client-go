package orderv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "customerId":
//        type: "string"
//    "description":
//        type: "string"
//        example: "My first project"
//    "diskspaceInGiB":
//        type: "number"
//        example: 20
//    "spec":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.MachineTypeSpec"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.HardwareSpec"}
// required:
//    - "diskspaceInGiB"
//    - "spec"

type ProjectHostingOrderPreview struct {
	CustomerId     *string                        `json:"customerId,omitempty"`
	Description    *string                        `json:"description,omitempty"`
	DiskspaceInGiB float64                        `json:"diskspaceInGiB"`
	Spec           ProjectHostingOrderPreviewSpec `json:"spec"`
}

func (o *ProjectHostingOrderPreview) Validate() error {
	if err := o.Spec.Validate(); err != nil {
		return fmt.Errorf("invalid property spec: %w", err)
	}
	return nil
}
