package conversationv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/conversationv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Message", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"conversationId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"createdBy\":{\"active\":true,\"avatarRefId\":\"string\",\"clearName\":\"string\",\"department\":\"development\",\"userId\":\"string\"},\"files\":[{\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"status\":\"requested\"}],\"internal\":true,\"messageContent\":\"string\",\"messageId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"type\":\"MESSAGE\"}")

			sut := conversationv2.Message{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
