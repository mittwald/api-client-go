package backupv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/backupv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectBackupSchedule", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"createdAt\":\"string\",\"description\":\"I'm a ProjectBackupSchedule\",\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"isSystemBackup\":true,\"projectId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"schedule\":\"5 4 * * *\",\"ttl\":\"7d\",\"updatedAt\":\"string\"}")

			sut := backupv1.ProjectBackupSchedule{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
