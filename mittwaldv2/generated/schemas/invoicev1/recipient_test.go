package invoicev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/invoicev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Recipient", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"address\":{\"addressPrefix\":\"c/o Ada Lovelace\",\"city\":\"Espelkamp\",\"countryCode\":\"DE\",\"houseNumber\":\"4-6\",\"street\":\"Königsberger Straße\",\"zip\":\"32339\"},\"company\":\"Mittwald CM Service GmbH \\u0026 Co. KG\",\"emailAddress\":\"string\",\"firstName\":\"Ada\",\"lastName\":\"Lovelace\",\"phoneNumbers\":[\"+49 123 4567890\"],\"salutation\":\"mr\",\"title\":\"Dr.\",\"useFormalTerm\":true}")

			sut := invoicev1.Recipient{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
