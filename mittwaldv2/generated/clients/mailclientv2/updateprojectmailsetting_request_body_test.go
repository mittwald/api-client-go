package mailclientv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/mailclientv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateProjectMailSettingRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal into AlternativeUpdateProjectMailSettingRequestBodyAlternative1", func() {
			exampleJSON := []byte("{\"blacklist\":[\"string\"]}")

			sut := mailclientv2.UpdateProjectMailSettingRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeUpdateProjectMailSettingRequestBodyAlternative1).NotTo(BeNil())
		})
		It("should unmarshal into AlternativeUpdateProjectMailSettingRequestBodyAlternative2", func() {
			exampleJSON := []byte("{\"whitelist\":[\"string\"]}")

			sut := mailclientv2.UpdateProjectMailSettingRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
			Expect(sut.AlternativeUpdateProjectMailSettingRequestBodyAlternative2).NotTo(BeNil())
		})
	})
})
