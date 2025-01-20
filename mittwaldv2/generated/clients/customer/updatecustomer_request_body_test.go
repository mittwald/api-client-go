package customer_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/customer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateCustomerRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"customerId\":\"string\",\"name\":\"string\",\"owner\":{\"address\":{\"addressPrefix\":\"c/o Ada Lovelace\",\"city\":\"Espelkamp\",\"countryCode\":\"DE\",\"houseNumber\":\"4-6\",\"street\":\"Königsberger Straße\",\"zip\":\"32339\"},\"company\":\"string\",\"emailAddress\":\"string\",\"firstName\":\"string\",\"lastName\":\"string\",\"phoneNumbers\":[],\"salutation\":\"mr\",\"title\":\"string\",\"useFormalTerm\":true},\"vatId\":\"string\"}")

			sut := customer.UpdateCustomerRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})