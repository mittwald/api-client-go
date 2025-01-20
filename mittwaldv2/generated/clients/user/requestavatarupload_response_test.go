package user_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/user"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RequestAvatarUploadResponse", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"refId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"rules\":{\"maxSizeInKB\":3000,\"mimeTypes\":[\"image/png\"],\"properties\":{\"imageDimensions\":{\"max\":{\"height\":42,\"width\":42},\"min\":{\"height\":42,\"width\":42}}}}}")

			sut := user.RequestAvatarUploadResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})