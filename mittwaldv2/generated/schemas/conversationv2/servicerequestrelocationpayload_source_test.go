package conversationv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/conversationv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceRequestRelocationPayloadSource", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"providerName\":\"string\",\"providerPassword\":\"string\",\"providerUrl\":\"string\",\"providerUsername\":\"string\",\"sourceAccount\":\"string\"}")

			sut := conversationv2.ServiceRequestRelocationPayloadSource{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
