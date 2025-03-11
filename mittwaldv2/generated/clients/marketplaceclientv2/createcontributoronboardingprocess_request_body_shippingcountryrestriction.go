package marketplaceclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "onlyDomestic"
// description: "You can restrict your customers to certain countries."

// You can restrict your customers to certain countries.
type CreateContributorOnboardingProcessRequestBodyShippingCountryRestriction string

const CreateContributorOnboardingProcessRequestBodyShippingCountryRestrictionOnlyDomestic CreateContributorOnboardingProcessRequestBodyShippingCountryRestriction = "onlyDomestic"

func (e CreateContributorOnboardingProcessRequestBodyShippingCountryRestriction) Validate() error {
	if e == CreateContributorOnboardingProcessRequestBodyShippingCountryRestrictionOnlyDomestic {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
