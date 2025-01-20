package databasev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/databasev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MySqlUser", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"accessIpMask\":\"string\",\"accessLevel\":\"full\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"databaseId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"description\":\"string\",\"disabled\":true,\"externalAccess\":true,\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"mainUser\":true,\"name\":\"string\",\"passwordUpdatedAt\":\"2006-01-02T15:04:05Z\",\"status\":\"pending\",\"statusSetAt\":\"2006-01-02T15:04:05Z\",\"updatedAt\":\"2006-01-02T15:04:05Z\"}")

			sut := databasev1.MySqlUser{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})