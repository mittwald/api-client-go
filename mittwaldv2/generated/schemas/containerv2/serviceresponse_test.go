package containerv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/containerv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceResponse", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"deployedState\":{\"command\":[\"string\"],\"entrypoint\":[\"string\"],\"envs\":{\"string\":\"string\"},\"image\":\"string\",\"ports\":[\"string\"],\"volumes\":[\"string\"]},\"description\":\"MySQL DB\",\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"pendingState\":{\"command\":[\"string\"],\"entrypoint\":[\"string\"],\"envs\":{\"string\":\"string\"},\"image\":\"string\",\"ports\":[\"string\"],\"volumes\":[\"string\"]},\"projectId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"serviceName\":\"mysql-db\",\"shortId\":\"c-12e4u6\",\"stackId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"status\":\"running\"}")

			sut := containerv2.ServiceResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
