package ingressv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/ingressv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PathTarget", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeTargetDirectory", func() {
			exampleJSON := []byte("{\"directory\":\"string\"}")

			sut := ingressv1.PathTarget{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeTargetDirectory).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeTargetUrl", func() {
			exampleJSON := []byte("{\"url\":\"string\"}")

			sut := ingressv1.PathTarget{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeTargetUrl).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeTargetInstallation", func() {
			exampleJSON := []byte("{\"installationId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := ingressv1.PathTarget{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeTargetInstallation).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeTargetUseDefaultPage", func() {
			exampleJSON := []byte("{\"useDefaultPage\":true}")

			sut := ingressv1.PathTarget{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeTargetUseDefaultPage).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeTargetContainer", func() {
			exampleJSON := []byte("{\"container\":{\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"portProtocol\":\"string\"}}")

			sut := ingressv1.PathTarget{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeTargetContainer).NotTo(BeNil())
		})
	})
})
