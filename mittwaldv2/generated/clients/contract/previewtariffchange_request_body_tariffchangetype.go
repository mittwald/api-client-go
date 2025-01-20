package contract

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "projectHosting"
//    - "server"

type PreviewTariffChangeRequestBodyTariffChangeType string

const PreviewTariffChangeRequestBodyTariffChangeTypeProjectHosting PreviewTariffChangeRequestBodyTariffChangeType = "projectHosting"
const PreviewTariffChangeRequestBodyTariffChangeTypeServer PreviewTariffChangeRequestBodyTariffChangeType = "server"

func (e PreviewTariffChangeRequestBodyTariffChangeType) Validate() error {
	if e == PreviewTariffChangeRequestBodyTariffChangeTypeProjectHosting || e == PreviewTariffChangeRequestBodyTariffChangeTypeServer {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}