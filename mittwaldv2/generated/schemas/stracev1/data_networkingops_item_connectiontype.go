package stracev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "UNKNOWN"
//    - "PRIVATE"
//    - "EXTERNAL"

type DataNetworkingOpsItemConnectionType string

const DataNetworkingOpsItemConnectionTypeUNKNOWN DataNetworkingOpsItemConnectionType = "UNKNOWN"
const DataNetworkingOpsItemConnectionTypePRIVATE DataNetworkingOpsItemConnectionType = "PRIVATE"
const DataNetworkingOpsItemConnectionTypeEXTERNAL DataNetworkingOpsItemConnectionType = "EXTERNAL"

func (e DataNetworkingOpsItemConnectionType) Validate() error {
	if e == DataNetworkingOpsItemConnectionTypeUNKNOWN || e == DataNetworkingOpsItemConnectionTypePRIVATE || e == DataNetworkingOpsItemConnectionTypeEXTERNAL {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
