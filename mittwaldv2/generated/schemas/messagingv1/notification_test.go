package messagingv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/messagingv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Notification", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"createdAt\":\"string\",\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"read\":true,\"reference\":{\"aggregate\":\"string\",\"domain\":\"string\",\"id\":\"string\",\"parents\":[{\"aggregate\":\"string\",\"domain\":\"string\",\"id\":\"string\"}]},\"severity\":\"success\",\"type\":\"string\"}")

			sut := messagingv1.Notification{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
