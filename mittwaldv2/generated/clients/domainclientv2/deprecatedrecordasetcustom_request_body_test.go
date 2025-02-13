package domainclientv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/domainclientv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeprecatedRecordASetCustomRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeRecordUnset", func() {
			exampleJSON := []byte("{}")

			sut := domainclientv2.DeprecatedRecordASetCustomRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordUnset).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCombinedACustom", func() {
			exampleJSON := []byte("{\"a\":[\"string\"],\"aaaa\":[\"string\"],\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := domainclientv2.DeprecatedRecordASetCustomRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCombinedACustom).NotTo(BeNil())
		})
	})
})
