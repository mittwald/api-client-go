package lead_finderv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/lead_finderv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UnlockedLead", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"businessFields\":[\"string\"],\"company\":null,\"contact\":{\"address\":{\"addressPrefix\":null,\"city\":\"Espelkamp\",\"countryCode\":\"DE\",\"houseNumber\":\"4-6\",\"street\":\"Königsberger Straße\",\"zip\":\"32339\"},\"company\":\"string\",\"emailAddress\":\"string\",\"phoneNumbers\":[\"string\"]},\"description\":\"string\",\"domain\":\"string\",\"hoster\":{\"mailServer\":[\"string\"],\"nameServer\":[\"string\"],\"server\":[\"string\"]},\"leadId\":\"string\",\"mainTechnology\":{\"name\":\"string\",\"version\":\"string\"},\"metrics\":{\"additionalMetrics\":{\"string\":{\"category\":3.14,\"score\":3.14,\"unit\":\"string\",\"value\":3.14}},\"basic\":{\"co2\":3.14,\"desktop\":{\"accessibility\":3.14,\"bestPractice\":3.14,\"performance\":3.14,\"seo\":3.14},\"mobile\":{\"accessibility\":3.14,\"bestPractice\":3.14,\"performance\":3.14,\"seo\":3.14},\"timeToFirstByteMs\":3.14}},\"reservedUntil\":\"2006-01-02T15:04:05Z\",\"score\":3.14,\"screenshot\":\"string\",\"socialMedia\":[{\"network\":\"string\",\"url\":\"string\"}],\"technologies\":[{\"name\":\"string\",\"version\":\"string\"}],\"unlockedAt\":\"2006-01-02T15:04:05Z\"}")

			sut := lead_finderv2.UnlockedLead{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
