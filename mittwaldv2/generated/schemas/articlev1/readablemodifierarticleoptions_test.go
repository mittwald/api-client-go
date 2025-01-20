package articlev1_test

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

import (
	"encoding/json"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/articlev1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReadableModifierArticleOptions", func() {
	When("unmarshaling from JSON", func() {
		It("should unmarshal", func() {
			exampleJSON := []byte("{\"articleId\":\"string\",\"info\":{\"articleName\":\"proSpace lite\",\"articleTemplateName\":\"proSpace\",\"fromArticleTemplate\":true},\"maxArticleCount\":3.14}")

			sut := articlev1.ReadableModifierArticleOptions{}
			Expect(json.Unmarshal(exampleJSON, &sut)).To(Succeed())
			Expect(sut.Validate()).To(Succeed())
		})
	})
})
