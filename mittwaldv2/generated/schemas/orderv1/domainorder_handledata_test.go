package orderv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DomainOrderHandleData", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"adminC\":[{\"name\":\"name\",\"value\":\"Ada Lovelace\"}],\"ownerC\":[{\"name\":\"name\",\"value\":\"Ada Lovelace\"}]}")

			sut := orderv1.DomainOrderHandleData{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})