package marketplacev2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/marketplacev2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Extension", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"assets\":[{\"assetType\":\"image\",\"fileName\":\"myFile.png\",\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"index\":1}],\"blocked\":true,\"context\":\"project\",\"contributorId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"deprecation\":{\"deprecatedAt\":\"2006-01-02T15:04:05Z\",\"note\":\"This extension is no longer actively maintained. Please Use the successor extension instead.\",\"successorId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"},\"description\":\"string\",\"detailedDescriptions\":{\"de\":{\"markdown\":\"string\",\"plain\":\"string\"},\"en\":{\"markdown\":\"string\",\"plain\":\"string\"}},\"disabled\":true,\"externalFrontends\":[{\"name\":\"string\",\"url\":\"string\"}],\"frontendComponents\":[{\"name\":\"string\",\"url\":\"string\"}],\"frontendFragments\":{\"string\":null},\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"logoRefId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"name\":\"MyPingExtension\",\"pricing\":{\"netPrice\":500},\"published\":true,\"scopes\":[\"string\"],\"state\":\"enabled\",\"statistics\":{\"amountOfInstances\":42},\"subTitle\":{\"de\":\"Ping deine App an\",\"en\":\"Ping your app\"},\"support\":{\"email\":\"a.lovelace@example.com\",\"phone\":\"string\"},\"tags\":[\"string\"]}")

			sut := marketplacev2.Extension{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
