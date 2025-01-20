package orderv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "contractId":
//        type: "string"
//    "diskspaceInGiB":
//        type: "number"
//        example: 20
//    "spec":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.MachineTypeSpec"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.order.HardwareSpec"}
//required:
//    - "contractId"
//    - "diskspaceInGiB"
//    - "spec"

//
type ProjectHostingTariffChange struct {
	ContractId     string                         `json:"contractId"`
	DiskspaceInGiB float64                        `json:"diskspaceInGiB"`
	Spec           ProjectHostingTariffChangeSpec `json:"spec"`
}

func (o *ProjectHostingTariffChange) Validate() error {
	if err := o.Spec.Validate(); err != nil {
		return fmt.Errorf("invalid property spec: %w", err)
	}
	return nil
}
