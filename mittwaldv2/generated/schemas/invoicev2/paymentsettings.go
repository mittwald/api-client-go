package invoicev2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.invoice.PaymentSettingsDebit"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.invoice.PaymentSettingsInvoice"}

type PaymentSettings struct {
	AlternativePaymentSettingsDebit   *PaymentSettingsDebit
	AlternativePaymentSettingsInvoice *PaymentSettingsInvoice
}

func (a *PaymentSettings) MarshalJSON() ([]byte, error) {
	if a.AlternativePaymentSettingsDebit != nil {
		return json.Marshal(a.AlternativePaymentSettingsDebit)
	}
	if a.AlternativePaymentSettingsInvoice != nil {
		return json.Marshal(a.AlternativePaymentSettingsInvoice)
	}
	return []byte("null"), nil
}

func (a *PaymentSettings) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativePaymentSettingsDebit PaymentSettingsDebit
	if err := dec.Decode(&alternativePaymentSettingsDebit); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativePaymentSettingsDebit.Validate(); vErr == nil {
			a.AlternativePaymentSettingsDebit = &alternativePaymentSettingsDebit
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativePaymentSettingsInvoice PaymentSettingsInvoice
	if err := dec.Decode(&alternativePaymentSettingsInvoice); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativePaymentSettingsInvoice.Validate(); vErr == nil {
			a.AlternativePaymentSettingsInvoice = &alternativePaymentSettingsInvoice
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *PaymentSettings) Validate() error {
	if a.AlternativePaymentSettingsDebit != nil {
		return a.AlternativePaymentSettingsDebit.Validate()
	}
	if a.AlternativePaymentSettingsInvoice != nil {
		return a.AlternativePaymentSettingsInvoice.Validate()
	}
	return errors.New("no alternative set")
}
