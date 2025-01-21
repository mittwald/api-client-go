package domain_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/domain"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateRecordSetRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeRecordUnset", func() {
			exampleJSON := []byte("{}")

			sut := domain.UpdateRecordSetRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordUnset).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCombinedACustom", func() {
			exampleJSON := []byte("{\"a\":[\"string\"],\"aaaa\":[\"string\"],\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := domain.UpdateRecordSetRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCombinedACustom).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeRecordMXCustom", func() {
			exampleJSON := []byte("{\"records\":[{\"fqdn\":\"string\",\"priority\":42}],\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := domain.UpdateRecordSetRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordMXCustom).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeRecordTXTComponent", func() {
			exampleJSON := []byte("{\"entries\":[\"string\"],\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := domain.UpdateRecordSetRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordTXTComponent).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeRecordSRVComponent", func() {
			exampleJSON := []byte("{\"records\":[{\"fqdn\":\"string\",\"port\":42,\"priority\":null,\"weight\":null}],\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := domain.UpdateRecordSetRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordSRVComponent).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeRecordCNAMEComponent", func() {
			exampleJSON := []byte("{\"fqdn\":\"string\",\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := domain.UpdateRecordSetRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordCNAMEComponent).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeRecordCAAComponent", func() {
			exampleJSON := []byte("{\"records\":[{\"flags\":42,\"tag\":\"issue\",\"value\":\"string\"}],\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := domain.UpdateRecordSetRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordCAAComponent).NotTo(BeNil())
		})
	})
})
