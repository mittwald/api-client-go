package marketplacev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/marketplacev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Contributor", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"customerId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"description\":\"string\",\"email\":\"string\",\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"logoRefId\":\"string\",\"name\":\"string\",\"phone\":\"string\",\"state\":\"enabled\",\"url\":\"string\"}")

			sut := marketplacev1.Contributor{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
