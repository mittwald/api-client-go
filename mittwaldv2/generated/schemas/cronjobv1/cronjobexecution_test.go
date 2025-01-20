package cronjobv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/cronjobv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CronjobExecution", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"abortedBy\":{\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"},\"cronjobId\":\"string\",\"durationInMilliseconds\":12374,\"end\":\"2006-01-02T15:04:05Z\",\"executionEnd\":\"2006-01-02T15:04:05Z\",\"executionStart\":\"2006-01-02T15:04:05Z\",\"id\":\"cron-bd26li-28027320\",\"logPath\":\"/var/log/cronjobs/cron-bd26li-28027320.log\",\"start\":\"2006-01-02T15:04:05Z\",\"status\":\"Complete\",\"successful\":true,\"triggeredBy\":{\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}}")

			sut := cronjobv1.CronjobExecution{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
