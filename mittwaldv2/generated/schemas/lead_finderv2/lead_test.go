package lead_finderv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/lead_finderv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lead", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"businessFields\":[\"string\"],\"company\":{\"city\":\"string\",\"employeeCount\":3.14,\"foundingYear\":3.14,\"salesVolume\":3.14},\"description\":\"string\",\"leadId\":\"string\",\"mainTechnology\":{\"name\":\"string\",\"version\":\"string\"},\"metrics\":{\"co2\":3.14,\"desktop\":{\"accessibility\":3.14,\"bestPractice\":3.14,\"performance\":3.14,\"seo\":3.14},\"mobile\":{\"accessibility\":3.14,\"bestPractice\":3.14,\"performance\":3.14,\"seo\":3.14},\"timeToFirstByteMs\":3.14},\"score\":3.14,\"screenshot\":\"string\",\"technologies\":[{\"name\":\"string\",\"version\":\"string\"}]}")

			sut := lead_finderv2.Lead{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
