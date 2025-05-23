package stracev2_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/stracev2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DataDbQueriesItem", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"query\":\"SELECT * FROM my_table;\",\"stats\":{\"kernelMs\":1.2345,\"occurrences\":25,\"syscallCount\":4321,\"userspaceMs\":1.2345},\"warnLevel\":\"NO\"}")

			sut := stracev2.DataDbQueriesItem{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
