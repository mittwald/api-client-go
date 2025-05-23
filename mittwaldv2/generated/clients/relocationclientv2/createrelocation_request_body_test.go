package relocationclientv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/relocationclientv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateRelocationRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"additionalServices\":{\"dataCompare\":\"default\"},\"allDomains\":true,\"allowPasswordChange\":true,\"articleType\":\"cms-hosting\",\"contact\":{\"email\":\"string\",\"firstName\":\"string\",\"lastName\":\"string\",\"phoneNumber\":\"string\"},\"domains\":[{\"authCode\":null,\"domainOwnerData\":null,\"name\":\"string\"}],\"notes\":\"string\",\"prices\":{\"positions\":[{\"name\":\"string\",\"price\":3.14}],\"total\":3.14},\"provider\":{\"loginUrl\":\"string\",\"name\":\"string\",\"password\":\"string\",\"sourceAccount\":\"string\",\"userName\":\"string\"},\"target\":{\"organisation\":\"string\",\"product\":\"string\",\"projectName\":\"string\",\"system\":\"kc\"},\"userId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := relocationclientv2.CreateRelocationRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
