package sslv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/sslv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CertificateErrorMessage", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative1", func() {
			exampleJSON := []byte("\"certificate_read_failed\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative1).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative2", func() {
			exampleJSON := []byte("\"certificate_decode_failed\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative2).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative3", func() {
			exampleJSON := []byte("\"certificate_parsing_failed\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative3).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative4", func() {
			exampleJSON := []byte("\"certificate_self_signed\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative4).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative5", func() {
			exampleJSON := []byte("\"certificate_not_authorized_to_sign\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative5).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative6", func() {
			exampleJSON := []byte("\"certificate_expired\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative6).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative7", func() {
			exampleJSON := []byte("\"ca_not_authorized_for_this_name\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative7).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative8", func() {
			exampleJSON := []byte("\"too_many_intermediates\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative8).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative9", func() {
			exampleJSON := []byte("\"incompatible_usage\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative9).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative10", func() {
			exampleJSON := []byte("\"unknown_authority\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative10).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative11", func() {
			exampleJSON := []byte("\"private_key_read_failed\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative11).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative12", func() {
			exampleJSON := []byte("\"private_key_decode_failed\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative12).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative13", func() {
			exampleJSON := []byte("\"private_key_parse_failed\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative13).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative14", func() {
			exampleJSON := []byte("\"private_key_encrypted\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative14).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative15", func() {
			exampleJSON := []byte("\"private_key_not_rsa\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative15).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative16", func() {
			exampleJSON := []byte("\"private_key_mismatch\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative16).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative17", func() {
			exampleJSON := []byte("\"unknown_cloudflare_error\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative17).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeCertificateErrorMessageAlternative18", func() {
			exampleJSON := []byte("\"unknown\"")

			sut := sslv2.CertificateErrorMessage{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeCertificateErrorMessageAlternative18).NotTo(BeNil())
		})
	})
})
