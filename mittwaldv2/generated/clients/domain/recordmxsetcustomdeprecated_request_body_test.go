package domain_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/domain"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecordMxSetCustomDeprecatedRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeRecordUnset", func() {
			exampleJSON := []byte("{}")

			sut := domain.RecordMxSetCustomDeprecatedRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordUnset).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeRecordMXCustom", func() {
			exampleJSON := []byte("{\"records\":[{\"fqdn\":\"string\",\"priority\":42}],\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := domain.RecordMxSetCustomDeprecatedRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordMXCustom).NotTo(BeNil())
		})
	})
})