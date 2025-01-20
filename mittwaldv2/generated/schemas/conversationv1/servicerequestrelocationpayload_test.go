package conversationv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/conversationv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceRequestRelocationPayload", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"contact\":{\"email\":\"string\",\"firstname\":\"string\",\"lastname\":\"string\",\"phone\":\"string\"},\"domain\":{\"allDomains\":true,\"domains\":[{\"authCode\":\"string\",\"domainName\":\"string\"}]},\"notes\":\"string\",\"preferredRelocationDate\":\"2006-01-02T15:04:05Z\",\"redirectusKey\":3.14,\"source\":{\"providerName\":\"string\",\"providerPassword\":\"string\",\"providerUrl\":\"string\",\"providerUsername\":\"string\",\"sourceAccount\":\"string\"},\"target\":{\"accountShortId\":\"string\",\"application\":\"string\",\"articleType\":\"string\",\"organisation\":\"string\",\"prices\":[{\"name\":null,\"price\":null}],\"product\":\"string\",\"withDataCompare\":true},\"userId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := conversationv1.ServiceRequestRelocationPayload{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
