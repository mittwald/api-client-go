package ingressv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "ERROR_UNSPECIFIED"
//    - "ERROR_QUAD_A"
//    - "ERROR_NO_A_RECORD"
//    - "ERROR_ACME_CERTIFICATE_REQUEST_DEADLINE_EXCEEDED"

type IngressDNSValidationErrorsItem string

const IngressDNSValidationErrorsItemERRORUNSPECIFIED IngressDNSValidationErrorsItem = "ERROR_UNSPECIFIED"
const IngressDNSValidationErrorsItemERRORQUADA IngressDNSValidationErrorsItem = "ERROR_QUAD_A"
const IngressDNSValidationErrorsItemERRORNOARECORD IngressDNSValidationErrorsItem = "ERROR_NO_A_RECORD"
const IngressDNSValidationErrorsItemERRORACMECERTIFICATEREQUESTDEADLINEEXCEEDED IngressDNSValidationErrorsItem = "ERROR_ACME_CERTIFICATE_REQUEST_DEADLINE_EXCEEDED"

func (e IngressDNSValidationErrorsItem) Validate() error {
	if e == IngressDNSValidationErrorsItemERRORUNSPECIFIED || e == IngressDNSValidationErrorsItemERRORQUADA || e == IngressDNSValidationErrorsItemERRORNOARECORD || e == IngressDNSValidationErrorsItemERRORACMECERTIFICATEREQUESTDEADLINEEXCEEDED {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
