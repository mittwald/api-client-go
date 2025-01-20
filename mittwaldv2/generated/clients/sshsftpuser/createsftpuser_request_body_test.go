package sshsftpuser_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/sshsftpuser"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateSFTPUserRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"accessLevel\":\"read\",\"authentication\":{\"password\":\"string\"},\"description\":\"string\",\"directories\":[\"string\"],\"expiresAt\":\"2006-01-02T15:04:05Z\"}")

			sut := sshsftpuser.CreateSFTPUserRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})