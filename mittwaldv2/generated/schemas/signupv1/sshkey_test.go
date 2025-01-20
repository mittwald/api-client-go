package signupv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/signupv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SshKey", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"algorithm\":\"ssh-rsa\",\"comment\":\"a.lovelace@example.com\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"expiresAt\":\"2006-01-02T15:04:05Z\",\"fingerprint\":\"a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1\",\"key\":\"string\",\"sshKeyId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}")

			sut := signupv1.SshKey{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})