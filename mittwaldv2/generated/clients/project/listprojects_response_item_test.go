package project_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/project"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListProjectsResponseItem", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"backupStorageUsageInBytes\":42,\"backupStorageUsageInBytesSetAt\":\"2006-01-02T15:04:05Z\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"customerId\":\"string\",\"customerMeta\":{\"id\":\"string\"},\"description\":\"string\",\"disableReason\":\"maliciousCode\",\"disabledAt\":\"2006-01-02T15:04:05Z\",\"enabled\":true,\"id\":\"string\",\"imageRefId\":\"string\",\"isReady\":true,\"projectHostingId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"readiness\":\"creating\",\"serverId\":\"string\",\"shortId\":\"string\",\"status\":\"pending\",\"statusSetAt\":\"2006-01-02T15:04:05Z\",\"webStorageUsageInBytes\":42,\"webStorageUsageInBytesSetAt\":\"2006-01-02T15:04:05Z\"}")

			sut := project.ListProjectsResponseItem{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
