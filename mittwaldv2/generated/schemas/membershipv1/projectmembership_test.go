package membershipv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/membershipv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProjectMembership", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"email\":\"string\",\"expiresAt\":\"2006-01-02T15:04:05Z\",\"id\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"inherited\":true,\"inviteId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"memberSince\":\"2006-01-02T15:04:05Z\",\"mfa\":true,\"projectId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\",\"role\":\"notset\",\"userId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := membershipv1.ProjectMembership{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})