package mailmigrationv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/mailmigrationv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MigrationMailbox", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"description\":\"string\",\"finished\":true,\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"migrationJobs\":{\"migrate\":{\"requirements\":{\"mailbox\":{\"mailsystem\":{\"imapClusterId\":\"string\",\"mailDirectory\":\"string\",\"rateLimitId\":\"string\"},\"name\":\"string\",\"quotaInBytes\":42,\"spamProtection\":{\"active\":true,\"deleteSensitivity\":42,\"folder\":42,\"keepDays\":42,\"relocateSensitivity\":42}},\"projectId\":\"string\"}}},\"name\":\"string\"}")

			sut := mailmigrationv1.MigrationMailbox{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
