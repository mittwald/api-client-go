package filev2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/filev2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FileUploadRules", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"extensions\":[\"\"],\"fileTypes\":[{\"extensions\":[\"\"],\"mimeType\":\"image/jpeg\"}],\"maxNameLength\":80,\"maxSizeInBytes\":1000000,\"maxSizeInKB\":1000,\"maxSizeInKb\":1000,\"mimeTypes\":[\"\"],\"properties\":{\"imageDimensions\":{\"max\":null,\"min\":null}}}")

			sut := filev2.FileUploadRules{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
