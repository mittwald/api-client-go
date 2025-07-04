package marketplaceclientv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/marketplaceclientv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExtensionRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeExtensionRequestBodyAlternative1", func() {
			exampleJSON := []byte("{\"consentedScopes\":[\"string\"],\"customerId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := marketplaceclientv2.ExtensionRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeExtensionRequestBodyAlternative1).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeExtensionRequestBodyAlternative2", func() {
			exampleJSON := []byte("{\"consentedScopes\":[\"string\"],\"projectId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := marketplaceclientv2.ExtensionRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeExtensionRequestBodyAlternative2).NotTo(BeNil())
		})
	})
})
