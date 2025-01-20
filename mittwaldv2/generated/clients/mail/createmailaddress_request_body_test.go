package mail_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/mail"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateMailAddressRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeCreateForwardAddress", func() {
			exampleJSON := []byte("{\"address\":\"string\",\"forwardAddresses\":[\"string\"]}")

			sut := mail.CreateMailAddressRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCreateForwardAddress).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCreateMailAddress", func() {
			exampleJSON := []byte("{\"address\":\"string\",\"isCatchAll\":true,\"mailbox\":{\"enableSpamProtection\":true,\"password\":\"string\",\"quotaInBytes\":2147483648}}")

			sut := mail.CreateMailAddressRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCreateMailAddress).NotTo(BeNil())
		})
	})
})
