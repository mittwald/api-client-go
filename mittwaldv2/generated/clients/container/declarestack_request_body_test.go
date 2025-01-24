package container_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/container"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeclareStackRequestBody", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"services\":{\"string\":{\"command\":null,\"description\":\"MySQL DB\",\"entrypoint\":null,\"envs\":null,\"image\":\"mysql\",\"ports\":[],\"volumes\":null}},\"volumes\":{\"string\":{\"name\":\"mysql-volume\"}}}")

			sut := container.DeclareStackRequestBody{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})