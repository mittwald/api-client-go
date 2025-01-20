package mail_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/mail"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateMailAddressSpamProtectionV2DeprecatedRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"spamProtection\":{\"active\":true,\"autoDeleteSpam\":true,\"folder\":\"inbox\",\"relocationMinSpamScore\":42}}")

			sut := mail.UpdateMailAddressSpamProtectionV2DeprecatedRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
