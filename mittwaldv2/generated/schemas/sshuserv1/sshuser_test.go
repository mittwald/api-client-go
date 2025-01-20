package sshuserv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/sshuserv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SshUser", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"active\":true,\"authUpdatedAt\":\"2006-01-02T15:04:05Z\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"description\":\"string\",\"expiresAt\":\"2006-01-02T15:04:05Z\",\"hasPassword\":true,\"id\":\"string\",\"projectId\":\"string\",\"publicKeys\":[{\"comment\":\"string\",\"key\":\"string\"}],\"updatedAt\":\"2006-01-02T15:04:05Z\",\"userName\":\"string\"}")

			sut := sshuserv1.SshUser{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
