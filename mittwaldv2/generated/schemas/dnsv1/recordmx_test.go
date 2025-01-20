package dnsv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/dnsv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecordMX", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeRecordUnset", func() {
			exampleJSON := []byte("{}")

			sut := dnsv1.RecordMX{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordUnset).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeRecordMXManaged", func() {
			exampleJSON := []byte("{\"managed\":true}")

			sut := dnsv1.RecordMX{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordMXManaged).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeRecordMXCustom", func() {
			exampleJSON := []byte("{\"records\":[{\"fqdn\":\"string\",\"priority\":42}],\"settings\":{\"ttl\":{\"seconds\":42}}}")

			sut := dnsv1.RecordMX{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeRecordMXCustom).NotTo(BeNil())
		})
	})
})
