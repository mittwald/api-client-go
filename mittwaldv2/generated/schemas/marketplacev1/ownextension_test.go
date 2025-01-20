package marketplacev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/marketplacev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OwnExtension", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"backendComponents\":{\"extensionAddedToContext\":{\"url\":\"string\"},\"extensionInstanceRemovedFromContext\":{\"url\":\"string\"},\"extensionInstanceSecretRotated\":{\"url\":\"string\"},\"extensionInstanceUpdated\":{\"url\":\"string\"}},\"blocked\":true,\"context\":\"project\",\"contributorId\":\"string\",\"deprecation\":{\"deprecatedAt\":\"2006-01-02T15:04:05Z\"},\"description\":\"string\",\"detailedDescriptions\":{\"de\":{\"markdown\":\"string\",\"plain\":\"string\"},\"en\":{\"markdown\":\"string\",\"plain\":\"string\"}},\"disabled\":true,\"frontendComponents\":[{\"name\":\"string\",\"url\":\"string\"}],\"frontendFragments\":{\"string\":{\"url\":\"string\"}},\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"name\":\"string\",\"scopes\":[\"string\"],\"state\":\"enabled\",\"support\":{\"email\":\"a.lovelace@example.com\",\"phone\":\"string\"},\"tags\":[\"string\"]}")

			sut := marketplacev1.OwnExtension{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})