package orderv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "machineType":
//        type: "string"
//        example: "prospace.2cpu.4gb"

//
type MachineTypeSpec struct {
	MachineType *string `json:"machineType,omitempty"`
}

func (o *MachineTypeSpec) Validate() error {
	return nil
}