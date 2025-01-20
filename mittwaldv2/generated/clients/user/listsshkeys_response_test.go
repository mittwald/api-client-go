package user_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/user"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ListSSHKeysResponse", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"sshKeys\":[{\"algorithm\":\"ssh-rsa\",\"comment\":\"a.lovelace@example.com\",\"createdAt\":\"2006-01-02T15:04:05Z\",\"expiresAt\":null,\"fingerprint\":\"a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1:a1\",\"key\":\"string\",\"sshKeyId\":\"7a9d8971-09b0-4c39-8c64-546b6e1875ce\"}]}")

			sut := user.ListSSHKeysResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})