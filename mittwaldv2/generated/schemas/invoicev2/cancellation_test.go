package invoicev2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/invoicev2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cancellation", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"cancellationId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"cancelledAt\":\"2006-01-02T15:04:05Z\",\"correctionNumber\":\"RG1234567\",\"pdfId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"reason\":\"Kulanz\"}")

			sut := invoicev2.Cancellation{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
