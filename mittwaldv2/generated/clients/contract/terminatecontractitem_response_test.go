package contract_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/contract"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TerminateContractItemResponse", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"contractId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"contractItemId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"reason\":\"Server wird nicht mehr benötigt\",\"terminationTargetDate\":\"2006-01-02T15:04:05Z\"}")

			sut := contract.TerminateContractItemResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
