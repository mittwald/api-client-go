package commonsv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/commonsv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Address", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"addressPrefix\":\"c/o Ada Lovelace\",\"city\":\"Espelkamp\",\"countryCode\":\"DE\",\"houseNumber\":\"4-6\",\"street\":\"Königsberger Straße\",\"zip\":\"32339\"}")

			sut := commonsv1.Address{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})