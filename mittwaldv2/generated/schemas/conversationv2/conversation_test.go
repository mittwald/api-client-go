package conversationv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/conversationv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conversation", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"category\":{\"categoryId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"name\":\"string\",\"referenceType\":[]},\"conversationId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"createdBy\":{\"active\":true,\"avatarRefId\":\"string\",\"clearName\":\"string\",\"department\":\"development\",\"userId\":\"string\"},\"lastMessage\":{\"createdAt\":\"2006-01-02T15:04:05Z\",\"createdBy\":{\"active\":true,\"avatarRefId\":\"string\",\"clearName\":\"string\",\"department\":\"development\",\"userId\":\"string\"}},\"lastMessageAt\":\"2006-01-02T15:04:05Z\",\"lastMessageBy\":{\"active\":true,\"avatarRefId\":\"string\",\"clearName\":\"string\",\"department\":\"development\",\"userId\":\"string\"},\"mainUser\":{\"active\":true,\"avatarRefId\":\"string\",\"clearName\":\"string\",\"department\":\"development\",\"userId\":\"string\"},\"notificationRoles\":[\"customer_owner\"],\"relatedTo\":{\"aggregate\":\"user\",\"domain\":\"user\",\"id\":\"string\"},\"relations\":[{\"aggregate\":\"string\",\"domain\":\"string\",\"id\":\"string\"}],\"sharedWith\":{\"aggregate\":\"user\",\"domain\":\"user\",\"id\":\"string\"},\"shortId\":\"string\",\"status\":\"open\",\"title\":\"string\",\"visibility\":\"shared\"}")

			sut := conversationv2.Conversation{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
