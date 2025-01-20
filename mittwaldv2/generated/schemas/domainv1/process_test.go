package domainv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/domainv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Process", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"error\":\"string\",\"lastUpdate\":\"2006-01-02T15:04:05Z\",\"processType\":\"UNSPECIFIED\",\"state\":\"UNSPECIFIED\",\"status\":\"string\",\"statusCode\":\"string\",\"transactionId\":\"string\"}")

			sut := domainv1.Process{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
