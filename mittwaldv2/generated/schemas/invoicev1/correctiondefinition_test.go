package invoicev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/invoicev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CorrectionDefinition", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"contractItemId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"creditPeriod\":{\"end\":\"2006-01-02T15:04:05Z\",\"start\":\"2006-01-02T15:04:05Z\"}}")

			sut := invoicev1.CorrectionDefinition{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
