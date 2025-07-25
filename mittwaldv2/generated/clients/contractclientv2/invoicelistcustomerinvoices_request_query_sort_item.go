package contractclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "invoiceNumber"

type InvoiceListCustomerInvoicesRequestQuerySortItem string

const InvoiceListCustomerInvoicesRequestQuerySortItemInvoiceNumber InvoiceListCustomerInvoicesRequestQuerySortItem = "invoiceNumber"

func (e InvoiceListCustomerInvoicesRequestQuerySortItem) Validate() error {
	if e == InvoiceListCustomerInvoicesRequestQuerySortItemInvoiceNumber {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
