package container_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/container"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateStackRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"services\":{\"string\":{\"command\":null,\"description\":null,\"entrypoint\":null,\"envs\":null,\"image\":null,\"ports\":null,\"volumes\":null}},\"volumes\":{\"string\":{\"name\":null}}}")

			sut := container.UpdateStackRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
