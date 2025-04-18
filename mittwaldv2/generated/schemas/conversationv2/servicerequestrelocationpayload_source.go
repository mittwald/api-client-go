package conversationv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "providerName":
//        type: "string"
//    "providerPassword":
//        type: "string"
//    "providerUrl":
//        type: "string"
//    "providerUsername":
//        type: "string"
//    "sourceAccount":
//        type: "string"
// required:
//    - "sourceAccount"
//    - "providerName"
//    - "providerUrl"
//    - "providerUsername"
//    - "providerPassword"

type ServiceRequestRelocationPayloadSource struct {
	ProviderName     string `json:"providerName"`
	ProviderPassword string `json:"providerPassword"`
	ProviderUrl      string `json:"providerUrl"`
	ProviderUsername string `json:"providerUsername"`
	SourceAccount    string `json:"sourceAccount"`
}

func (o *ServiceRequestRelocationPayloadSource) Validate() error {
	return nil
}
