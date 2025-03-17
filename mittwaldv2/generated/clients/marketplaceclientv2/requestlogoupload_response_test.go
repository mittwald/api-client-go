package marketplaceclientv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/marketplaceclientv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RequestLogoUploadResponse", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"logoRefId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"rules\":{\"extensions\":[\".png\"],\"fileTypes\":[{\"extensions\":[\".png\"],\"mimeType\":\"image/png\"}],\"maxSizeInBytes\":42,\"mimeTypes\":[\"image/png\"],\"properties\":{\"imageDimensions\":{\"max\":{\"height\":42,\"width\":42},\"min\":{\"height\":42,\"width\":42}}}}}")

			sut := marketplaceclientv2.RequestLogoUploadResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
