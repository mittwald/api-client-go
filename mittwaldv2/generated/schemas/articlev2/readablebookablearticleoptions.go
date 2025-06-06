package articlev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "articleId":
//        type: "string"
//        minLength: 1
//        example: "WH25-0001"
//    "info":
//        type: "object"
//        properties:
//            "articleName":
//                type: "string"
//                example: "proSpace lite"
//            "articleTemplateName":
//                type: "string"
//                example: "proSpace"
//            "fromArticleTemplate":
//                type: "boolean"
//    "maxArticleCount":
//        type: "number"
//        example: 10
// required:
//    - "articleId"

type ReadableBookableArticleOptions struct {
	ArticleId       string                              `json:"articleId"`
	Info            *ReadableBookableArticleOptionsInfo `json:"info,omitempty"`
	MaxArticleCount *float64                            `json:"maxArticleCount,omitempty"`
}

func (o *ReadableBookableArticleOptions) Validate() error {
	if err := func() error {
		if o.Info == nil {
			return nil
		}
		return o.Info.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property info: %w", err)
	}
	return nil
}
