package leadfyndrv2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/leadfyndrv2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UnlockedLead", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"actualUrl\":\"string\",\"businessFields\":[\"string\"],\"company\":null,\"contact\":{\"address\":{\"address_prefix\":null,\"city\":null,\"country_code\":null,\"house_number\":null,\"street\":null,\"zip\":null}},\"description\":\"string\",\"domain\":\"string\",\"hoster\":{\"mailServer\":[\"string\"],\"nameServer\":[\"string\"],\"server\":[\"string\"]},\"languages\":[\"string\"],\"leadId\":\"string\",\"mainTechnology\":{\"categories\":[],\"categoryPriority\":42,\"name\":\"string\",\"version\":\"string\"},\"metrics\":{\"additionalMetrics\":{\"string\":{\"category\":\"string\",\"name\":\"string\",\"score\":3.14,\"unit\":\"string\",\"value\":3.14}},\"basic\":{\"co2\":null,\"contentLoaded\":null,\"desktop\":{\"accessibility\":3.14,\"bestPractice\":3.14,\"cumulativeLayoutShift\":3.14,\"firstContentfulPaint\":3.14,\"largestContentfulPaint\":3.14,\"performance\":3.14,\"seo\":3.14,\"totalBlockingTime\":3.14},\"mobile\":{\"accessibility\":3.14,\"bestPractice\":3.14,\"cumulativeLayoutShift\":3.14,\"firstContentfulPaint\":3.14,\"largestContentfulPaint\":3.14,\"performance\":3.14,\"seo\":3.14,\"totalBlockingTime\":3.14},\"timeToFirstByteMs\":null}},\"potential\":3.14,\"reservationAllowed\":true,\"reservedAt\":\"2006-01-02T15:04:05Z\",\"scannedAt\":\"2006-01-02T15:04:05Z\",\"screenshot\":\"string\",\"socialMedia\":[{\"network\":\"string\",\"url\":\"string\"}],\"technologies\":[{\"categories\":[],\"categoryPriority\":42,\"name\":\"string\",\"version\":\"string\"}],\"unlockedAt\":\"2006-01-02T15:04:05Z\"}")

			sut := leadfyndrv2.UnlockedLead{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
