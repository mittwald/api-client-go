package orderv1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/orderv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServerOrder", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"customerId\":\"f3435305-fd26-470e-9f21-43d9be7e67e7\",\"description\":\"My first project\",\"diskspaceInGiB\":100,\"machineType\":\"shared.xlarge\",\"promotionCode\":\"123456\",\"recommendationCode\":\"mp-123456\",\"useFreeTrial\":true}")

			sut := orderv1.ServerOrder{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
