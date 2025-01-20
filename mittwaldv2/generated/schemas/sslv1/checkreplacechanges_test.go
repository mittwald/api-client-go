package sslv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/sslv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CheckReplaceChanges", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"commonName\":{\"newValue\":\"string\",\"oldValue\":\"string\"},\"dnsNames\":{\"addedValues\":[\"string\"],\"removedValues\":[\"string\"],\"values\":[\"string\"]},\"issuer\":{\"newValue\":\"string\",\"oldValue\":\"string\"},\"validFrom\":{\"newValue\":\"string\",\"oldValue\":\"string\"},\"validTo\":{\"newValue\":\"string\",\"oldValue\":\"string\"}}")

			sut := sslv1.CheckReplaceChanges{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
