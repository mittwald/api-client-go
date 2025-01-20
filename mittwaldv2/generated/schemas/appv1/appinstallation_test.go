package appv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/appv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppInstallation", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"appId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"appVersion\":{\"current\":\"string\",\"desired\":\"string\"},\"customDocumentRoot\":\"string\",\"description\":\"string\",\"disabled\":true,\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"installationPath\":\"string\",\"linkedDatabases\":[{\"databaseId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"databaseUserIds\":null,\"kind\":\"mysql\",\"purpose\":\"primary\"}],\"processes\":[\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"],\"projectId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"screenshotId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"screenshotRef\":\"string\",\"shortId\":\"a-XXXXXX\",\"systemSoftware\":[{\"systemSoftwareId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"systemSoftwareVersion\":{\"current\":\"string\",\"desired\":\"string\"},\"updatePolicy\":\"none\"}],\"updatePolicy\":\"none\",\"userInputs\":[{\"name\":\"string\",\"value\":\"string\"}]}")

			sut := appv1.AppInstallation{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
