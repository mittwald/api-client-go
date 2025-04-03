package orderv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "NEW_ORDER"
//    - "CONTRACT_CHANGE"

type OrderType string

const OrderTypeNEWORDER OrderType = "NEW_ORDER"
const OrderTypeCONTRACTCHANGE OrderType = "CONTRACT_CHANGE"

func (e OrderType) Validate() error {
	if e == OrderTypeNEWORDER || e == OrderTypeCONTRACTCHANGE {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
