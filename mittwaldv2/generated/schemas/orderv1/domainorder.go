package orderv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "authCode":
//        type: "string"
//        example: "XXXXXXX"
//    "domain":
//        type: "string"
//        example: "mittwald.example"
//    "handleData":
//        type: "object"
//        properties:
//            "adminC":
//                type: "array"
//                items: {"$ref": "#/components/schemas/de.mittwald.v1.order.DomainHandleField"}
//                deprecated: true
//            "ownerC":
//                type: "array"
//                items: {"$ref": "#/components/schemas/de.mittwald.v1.order.DomainHandleField"}
//                minItems: 1
//        required:
//            - "ownerC"
//    "projectId":
//        type: "string"
//        format: "uuid"
// required:
//    - "projectId"
//    - "domain"
//    - "handleData"

type DomainOrder struct {
	AuthCode   *string               `json:"authCode,omitempty"`
	Domain     string                `json:"domain"`
	HandleData DomainOrderHandleData `json:"handleData"`
	ProjectId  string                `json:"projectId"`
}

func (o *DomainOrder) Validate() error {
	if err := o.HandleData.Validate(); err != nil {
		return fmt.Errorf("invalid property handleData: %w", err)
	}
	return nil
}
