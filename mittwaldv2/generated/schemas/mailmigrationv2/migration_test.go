package mailmigrationv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/mailmigrationv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Migration", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"aborted\":true,\"addresses\":[{\"address\":\"string\",\"finished\":true,\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"migrationJobs\":{\"migrate\":{\"finished\":true,\"requirements\":{\"address\":\"string\",\"autoResponder\":{\"active\":true,\"expiresAt\":null,\"message\":\"string\",\"startsAt\":null},\"forwardAddresses\":[\"string\"],\"isCatchAll\":true,\"mailbox\":{\"mailsystem\":{\"imapClusterId\":\"string\",\"mailDirectory\":\"string\",\"rateLimitId\":\"string\"},\"name\":\"string\",\"quotaInBytes\":42,\"spamProtection\":{\"active\":true,\"deleteSensitivity\":42,\"folder\":42,\"keepDays\":42,\"relocateSensitivity\":42}},\"projectId\":\"string\"}}},\"preMigrationJobs\":{\"aliasSet\":[{\"finished\":true,\"sourceCoabMailboxName\":\"string\"}],\"deliveryMigrations\":[{\"finished\":true,\"sourceCoabDeliveryMailbox\":\"string\",\"sourceCoabDeliveryUid\":42,\"targetDeliveryAddress\":\"string\"}]}}],\"finalizers\":{\"disableLegacyEntities\":{\"addresses\":[\"string\"],\"mailboxNames\":[\"string\"]},\"projectSettingMigrations\":{\"blacklistEntries\":[\"string\"],\"whitelistEntries\":[\"string\"]}},\"finished\":true,\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"mailboxes\":[{\"description\":\"string\",\"finished\":true,\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"migrationJobs\":{\"migrate\":{\"requirements\":{\"mailbox\":{\"mailsystem\":{\"imapClusterId\":\"string\",\"mailDirectory\":\"string\",\"rateLimitId\":\"string\"},\"name\":\"string\",\"quotaInBytes\":42,\"spamProtection\":{\"active\":true,\"deleteSensitivity\":42,\"folder\":42,\"keepDays\":42,\"relocateSensitivity\":42}},\"projectId\":\"string\"}}},\"name\":\"string\"}],\"sourceCoabProjectId\":\"string\",\"targetNexusProjectId\":\"string\"}")

			sut := mailmigrationv2.Migration{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
