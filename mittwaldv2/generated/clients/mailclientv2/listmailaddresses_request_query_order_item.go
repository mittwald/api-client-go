package mailclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "asc"
//    - "desc"

type ListMailAddressesRequestQueryOrderItem string

const ListMailAddressesRequestQueryOrderItemAsc ListMailAddressesRequestQueryOrderItem = "asc"
const ListMailAddressesRequestQueryOrderItemDesc ListMailAddressesRequestQueryOrderItem = "desc"

func (e ListMailAddressesRequestQueryOrderItem) Validate() error {
	if e == ListMailAddressesRequestQueryOrderItemAsc || e == ListMailAddressesRequestQueryOrderItemDesc {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
