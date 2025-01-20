package orderv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"github.com/google/uuid"
)

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "certificateRequestId":
//        type: "string"
//        format: "uuid"
//    "projectId":
//        type: "string"
//        format: "uuid"
//required:
//    - "projectId"
//    - "certificateRequestId"

//
type ExternalCertificateOrderPreview struct {
	CertificateRequestId uuid.UUID `json:"certificateRequestId"`
	ProjectId            uuid.UUID `json:"projectId"`
}

func (o *ExternalCertificateOrderPreview) Validate() error {
	return nil
}