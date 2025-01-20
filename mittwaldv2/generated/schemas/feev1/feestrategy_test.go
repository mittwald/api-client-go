package feev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/feev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FeeStrategy", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeOneTimePaymentFeeStrategy", func() {
			exampleJSON := []byte("{\"price\":3.14}")

			sut := feev1.FeeStrategy{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeOneTimePaymentFeeStrategy).NotTo(BeNil())
		})
		It("should unmarshal into AlternativePeriodBasedFeeStrategy", func() {
			exampleJSON := []byte("{\"periods\":[{\"feeValidFrom\":\"string\",\"feeValidUntil\":\"string\",\"monthlyPrice\":3.14}]}")

			sut := feev1.FeeStrategy{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativePeriodBasedFeeStrategy).NotTo(BeNil())
		})
	})
})
