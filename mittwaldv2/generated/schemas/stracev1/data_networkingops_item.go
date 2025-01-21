package stracev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "connectionType":
//        type: "string"
//        enum:
//            - "UNKNOWN"
//            - "PRIVATE"
//            - "EXTERNAL"
//    "description":
//        type: "string"
//        description: "A short description of the network connection to provide additional context."
//    "ip":
//        type: "string"
//        description: "IP address to which a connection was established."
//    "port":
//        type: "integer"
//        description: "Port to which a connection was established."
//    "stats": {"$ref": "#/components/schemas/de.mittwald.v1.strace.Statistics"}
//    "warnLevel":
//        type: "string"
//        enum:
//            - "NO"
//            - "WARN"
//            - "SEVERE"
//        description: "Alerts when the time, syscall count or occurrence count of this group are abnormal."
//required:
//    - "ip"
//    - "port"
//    - "description"
//    - "connectionType"
//    - "stats"
//    - "warnLevel"

type DataNetworkingOpsItem struct {
	ConnectionType DataNetworkingOpsItemConnectionType `json:"connectionType"`
	Description    string                              `json:"description"`
	Ip             string                              `json:"ip"`
	Port           int64                               `json:"port"`
	Stats          Statistics                          `json:"stats"`
	WarnLevel      DataNetworkingOpsItemWarnLevel      `json:"warnLevel"`
}

func (o *DataNetworkingOpsItem) Validate() error {
	if err := o.ConnectionType.Validate(); err != nil {
		return fmt.Errorf("invalid property connectionType: %w", err)
	}
	if err := o.Stats.Validate(); err != nil {
		return fmt.Errorf("invalid property stats: %w", err)
	}
	if err := o.WarnLevel.Validate(); err != nil {
		return fmt.Errorf("invalid property warnLevel: %w", err)
	}
	return nil
}
