package domainv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/domainv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Domain", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"authCode\":{\"expires\":\"2006-01-02T15:04:05Z\",\"value\":\"string\"},\"authCode2\":{\"expires\":\"2006-01-02T15:04:05Z\"},\"connected\":true,\"contactHash\":\"string\",\"deleted\":true,\"domain\":\"string\",\"domainId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"handles\":{\"adminC\":{\"current\":{\"handleFields\":[{\"name\":\"string\",\"value\":\"string\"}],\"handleRef\":\"string\"},\"desired\":null},\"ownerC\":{\"current\":{\"handleFields\":[{\"name\":\"string\",\"value\":\"string\"}],\"handleRef\":\"string\"},\"desired\":{\"handleFields\":[{\"name\":\"string\",\"value\":\"string\"}],\"handleRef\":\"string\"}}},\"nameservers\":[\"string\"],\"processes\":[{\"error\":null,\"lastUpdate\":\"2006-01-02T15:04:05Z\",\"processType\":\"UNSPECIFIED\",\"state\":\"UNSPECIFIED\",\"status\":null,\"statusCode\":null,\"transactionId\":\"string\"}],\"projectId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"transferInAuthCode\":\"string\",\"usesDefaultNameserver\":true}")

			sut := domainv1.Domain{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
