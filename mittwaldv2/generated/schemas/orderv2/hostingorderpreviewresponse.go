package orderv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "machineTypePrice":
//        type: "number"
//        example: 500
//    "storagePrice":
//        type: "number"
//        example: 1000
//    "totalPrice":
//        type: "number"
//        example: 1500
// required:
//    - "totalPrice"
//    - "storagePrice"
//    - "machineTypePrice"

type HostingOrderPreviewResponse struct {
	MachineTypePrice float64 `json:"machineTypePrice"`
	StoragePrice     float64 `json:"storagePrice"`
	TotalPrice       float64 `json:"totalPrice"`
}

func (o *HostingOrderPreviewResponse) Validate() error {
	return nil
}
