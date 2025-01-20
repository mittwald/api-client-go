package domain_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/domain"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TLSDeprecatedRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeTlsAcme", func() {
			exampleJSON := []byte("{\"acme\":true,\"isCreated\":true,\"requestDeadline\":\"string\"}")

			sut := domain.TLSDeprecatedRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeTlsAcme).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeTlsCertificate", func() {
			exampleJSON := []byte("{\"certificateId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := domain.TLSDeprecatedRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeTlsCertificate).NotTo(BeNil())
		})
	})
})
