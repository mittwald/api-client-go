package feev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/feev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PeriodBasedFeeStrategy", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"periods\":[{\"feeValidFrom\":\"2006-01-02T15:04:05Z\",\"feeValidUntil\":\"2006-01-02T15:04:05Z\",\"monthlyPrice\":3.14}]}")

			sut := feev1.PeriodBasedFeeStrategy{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
