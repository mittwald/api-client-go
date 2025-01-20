package storagespacev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/storagespacev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StatisticsCategory", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"kind\":\"webspace\",\"resources\":[{\"description\":null,\"id\":\"169cea81-2c11-46a4-8f0b-5b0b47caeb78\",\"name\":\"mysql-xyz\",\"usageInBytes\":1000,\"usageInBytesSetAt\":\"2023-12-22T13:46:52.000Z\"}],\"totalUsageInBytes\":1000}")

			sut := storagespacev1.StatisticsCategory{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
