package containerv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/containerv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ContainerImageConfig", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"command\":[\"mysqld\"],\"entrypoint\":[\"docker-entrypoint.sh\"],\"env\":[{\"description\":null,\"isAiGenerated\":true,\"isSensitive\":null,\"key\":\"MYSQL_PASSWORD\",\"value\":null}],\"exposedPorts\":[{\"description\":null,\"isAiGenerated\":true,\"port\":\"3306/tcp\"}],\"hasAiGeneratedData\":true,\"isAiAvailable\":true,\"isUserRoot\":false,\"overwritingUser\":1000,\"user\":\"mysql\",\"userId\":0,\"volumes\":[{\"description\":null,\"isAiGenerated\":true,\"volume\":\"/var/lib/mysql\"}]}")

			sut := containerv2.ContainerImageConfig{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
