package contract_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/contract"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PreviewOrderResponse", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeHostingOrderPreviewResponse", func() {
			exampleJSON := []byte("{\"machineTypePrice\":500,\"storagePrice\":1000,\"totalPrice\":1500}")

			sut := contract.PreviewOrderResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeHostingOrderPreviewResponse).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeDomainOrderPreviewResponse", func() {
			exampleJSON := []byte("{\"domainContractDuration\":12,\"domainPrice\":800,\"feePrice\":100,\"totalPrice\":900}")

			sut := contract.PreviewOrderResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeDomainOrderPreviewResponse).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeExternalCertificateOrderPreviewResponse", func() {
			exampleJSON := []byte("{\"feePrice\":900,\"recurringPrice\":490,\"totalPrice\":1390}")

			sut := contract.PreviewOrderResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeExternalCertificateOrderPreviewResponse).NotTo(BeNil())
		})
	})
})