package ingressv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "container":
//        type: "object"
//        properties:
//            "id":
//                type: "string"
//                format: "uuid"
//            "portProtocol":
//                type: "string"
//                format: "docker-network-port"
//                description: "docker-compose port specification in format port/protocol (e.g. 8080/TCP)"
//        required:
//            - "id"
//            - "portProtocol"
// required:
//    - "container"

type TargetContainer struct {
	Container TargetContainerContainer `json:"container"`
}

func (o *TargetContainer) Validate() error {
	if err := o.Container.Validate(); err != nil {
		return fmt.Errorf("invalid property container: %w", err)
	}
	return nil
}
