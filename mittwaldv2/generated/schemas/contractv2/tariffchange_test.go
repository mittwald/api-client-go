package contractv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/contractv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TariffChange", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"newArticles\":[{\"amount\":1,\"articleTemplateId\":\"a1b8f0e9-904f-4716-a1c0-81ccf5342a56\",\"description\":\"Musterbeschreibung\",\"id\":\"a1b8f0e9-904f-4716-a1c0-81ccf5342a56\",\"name\":\"Musterartikel\",\"unitPrice\":{\"currency\":\"EUR\",\"value\":100}}],\"scheduledAtDate\":\"2006-01-02T15:04:05Z\",\"scheduledByUserId\":\"string\",\"targetDate\":\"2006-01-02T15:04:05Z\"}")

			sut := contractv2.TariffChange{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
