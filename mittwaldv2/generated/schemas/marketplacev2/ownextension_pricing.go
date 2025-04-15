package marketplacev2

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
//    - {"$ref": "#/components/schemas/de.mittwald.v1.marketplace.MonthlyPricingStrategy"}

type OwnExtensionPricing struct {
	AlternativeMonthlyPricingStrategy *MonthlyPricingStrategy
}

func (a *OwnExtensionPricing) MarshalJSON() ([]byte, error) {
	if a.AlternativeMonthlyPricingStrategy != nil {
		return json.Marshal(a.AlternativeMonthlyPricingStrategy)
	}
	return []byte("null"), nil
}

func (a *OwnExtensionPricing) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeMonthlyPricingStrategy MonthlyPricingStrategy
	if err := dec.Decode(&alternativeMonthlyPricingStrategy); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeMonthlyPricingStrategy.Validate(); vErr == nil {
			a.AlternativeMonthlyPricingStrategy = &alternativeMonthlyPricingStrategy
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *OwnExtensionPricing) Validate() error {
	if a.AlternativeMonthlyPricingStrategy != nil {
		return a.AlternativeMonthlyPricingStrategy.Validate()
	}
	return errors.New("no alternative set")
}
