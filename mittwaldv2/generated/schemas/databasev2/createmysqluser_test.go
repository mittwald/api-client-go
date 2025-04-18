package databasev2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/databasev2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateMySqlUser", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"accessIpMask\":\"string\",\"accessLevel\":\"full\",\"databaseId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"description\":\"string\",\"externalAccess\":true,\"password\":\"string\"}")

			sut := databasev2.CreateMySqlUser{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
