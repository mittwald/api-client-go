package contractv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/contractv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Contact", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"address\":{\"addressPrefix\":\"c/o Ada Lovelace\",\"city\":\"Espelkamp\",\"countryCode\":\"DE\",\"houseNumber\":\"4-6\",\"street\":\"Königsberger Straße\",\"zip\":\"32339\"},\"company\":\"Musterfirma\",\"emailAddress\":\"string\",\"firstName\":\"Ada\",\"lastName\":\"Lovelace\",\"phoneNumbers\":[\"+49 123 4567890\"],\"salutation\":\"mr\",\"title\":\"Dr.\",\"useFormalTerm\":true}")

			sut := contractv1.Contact{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
