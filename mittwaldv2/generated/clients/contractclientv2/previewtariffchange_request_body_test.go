package contractclientv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/contractclientv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PreviewTariffChangeRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"tariffChangeData\":{\"contractId\":\"string\",\"diskspaceInGiB\":20,\"spec\":{\"machineType\":\"prospace.2cpu.4gb\"}},\"tariffChangeType\":\"projectHosting\"}")

			sut := contractclientv2.PreviewTariffChangeRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
