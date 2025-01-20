package conversation_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/conversation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListMessagesByConversationResponseItem", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeMessage", func() {
			exampleJSON := []byte("{\"conversationId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"createdBy\":{\"active\":null,\"avatarRefId\":null,\"clearName\":null,\"department\":null,\"userId\":\"string\"},\"files\":[{\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"status\":\"requested\"}],\"internal\":true,\"messageContent\":\"string\",\"messageId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"type\":\"MESSAGE\"}")

			sut := conversation.ListMessagesByConversationResponseItem{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeMessage).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeStatusUpdate", func() {
			exampleJSON := []byte("{\"conversationId\":\"string\",\"createdAt\":\"string\",\"internal\":true,\"messageContent\":\"string\",\"meta\":{\"user\":{\"active\":true,\"avatarRefId\":\"string\",\"clearName\":\"string\",\"department\":\"development\",\"userId\":\"string\"}},\"type\":\"STATUS_UPDATE\"}")

			sut := conversation.ListMessagesByConversationResponseItem{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeStatusUpdate).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeServiceRequest", func() {
			exampleJSON := []byte("{\"conversationId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"files\":[{\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"status\":\"requested\"}],\"internal\":true,\"messageContent\":\"relocation\",\"messageId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"meta\":{\"contact\":{\"email\":\"string\",\"firstname\":\"string\",\"lastname\":\"string\",\"phone\":null},\"domain\":{\"allDomains\":true,\"domains\":[{\"authCode\":\"string\",\"domainName\":\"string\"}]},\"notes\":\"string\",\"preferredRelocationDate\":\"2006-01-02T15:04:05Z\",\"redirectusKey\":3.14,\"source\":{\"providerName\":\"string\",\"providerPassword\":\"string\",\"providerUrl\":\"string\",\"providerUsername\":\"string\",\"sourceAccount\":\"string\"},\"target\":{\"accountShortId\":\"string\",\"application\":null,\"articleType\":\"string\",\"organisation\":\"string\",\"prices\":null,\"product\":\"string\",\"withDataCompare\":null},\"userId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"},\"type\":\"SERVICE_REQUEST\"}")

			sut := conversation.ListMessagesByConversationResponseItem{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeServiceRequest).NotTo(BeNil())
		})
	})
})