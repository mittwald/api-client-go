package databasev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/databasev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RedisDatabase", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"configuration\":{\"additionalFlags\":[],\"maxMemory\":\"64Mi\",\"maxMemoryPolicy\":\"allkeys-lru\",\"persistent\":true},\"createdAt\":\"string\",\"description\":\"string\",\"finalizers\":[\"string\"],\"hostname\":\"string\",\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"name\":\"string\",\"port\":42,\"projectId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"status\":\"pending\",\"statusSetAt\":\"string\",\"storageUsageInBytes\":42,\"storageUsageInBytesSetAt\":\"string\",\"updatedAt\":\"string\",\"version\":\"string\"}")

			sut := databasev1.RedisDatabase{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
