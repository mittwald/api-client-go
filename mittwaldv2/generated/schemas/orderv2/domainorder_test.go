package orderv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DomainOrder", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"authCode\":\"XXXXXXX\",\"domain\":\"mittwald.example\",\"handleData\":{\"adminC\":[{\"name\":\"name\",\"value\":\"Ada Lovelace\"}],\"ownerC\":[{\"name\":\"name\",\"value\":\"Ada Lovelace\"}]},\"projectId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := orderv2.DomainOrder{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
