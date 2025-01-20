package orderv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectHostingOrderSpec", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeMachineTypeSpec", func() {
			exampleJSON := []byte("{\"machineType\":\"prospace.2cpu.4gb\"}")

			sut := orderv1.ProjectHostingOrderSpec{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeMachineTypeSpec).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeHardwareSpec", func() {
			exampleJSON := []byte("{\"ram\":2,\"vcpu\":1}")

			sut := orderv1.ProjectHostingOrderSpec{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeHardwareSpec).NotTo(BeNil())
		})
	})
})
