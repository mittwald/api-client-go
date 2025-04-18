package invoicev2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/invoicev2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PaymentSettings", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativePaymentSettingsDebit", func() {
			exampleJSON := []byte("{\"accountHolder\":\"Ada Lovelace\",\"bic\":\"DEUTDEDB123\",\"iban\":\"DE12345678901234567890\",\"method\":\"debit\"}")

			sut := invoicev2.PaymentSettings{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativePaymentSettingsDebit).NotTo(BeNil())
		})
		It("should unmarshal into AlternativePaymentSettingsInvoice", func() {
			exampleJSON := []byte("{\"method\":\"invoice\"}")

			sut := invoicev2.PaymentSettings{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativePaymentSettingsInvoice).NotTo(BeNil())
		})
	})
})
