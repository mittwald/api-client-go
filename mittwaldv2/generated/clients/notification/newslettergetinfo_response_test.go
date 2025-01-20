package notification_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/notification"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewsletterGetInfoResponse", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"active\":true,\"email\":\"a.lovelace@example.com\",\"registered\":true}")

			sut := notification.NewsletterGetInfoResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
