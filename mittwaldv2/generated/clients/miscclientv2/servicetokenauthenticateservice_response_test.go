package miscclientv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/miscclientv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServicetokenAuthenticateServiceResponse", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"accessToken\":{\"id\":\"string\",\"jwtClaims\":{\"exp\":\"string\",\"iat\":\"string\",\"iss\":\"string\",\"sub\":\"string\"},\"publicToken\":\"string\"}}")

			sut := miscclientv2.ServicetokenAuthenticateServiceResponse{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
