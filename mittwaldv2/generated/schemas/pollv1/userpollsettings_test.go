package pollv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/pollv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserPollSettings", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"completedAt\":\"2006-01-02T15:04:05Z\",\"dontShowUntil\":\"2006-01-02T15:04:05Z\",\"ignoredAt\":\"2006-01-02T15:04:05Z\",\"shouldShow\":true,\"status\":\"completed\",\"userId\":\"string\"}")

			sut := pollv1.UserPollSettings{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})