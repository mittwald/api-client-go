package articlev1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "articleName":
//        type: "string"
//        example: "proSpace lite"
//    "articleTemplateName":
//        type: "string"
//        example: "proSpace"
//    "fromArticleTemplate":
//        type: "boolean"

type ReadableChangeArticleOptionsInfo struct {
	ArticleName         *string `json:"articleName,omitempty"`
	ArticleTemplateName *string `json:"articleTemplateName,omitempty"`
	FromArticleTemplate *bool   `json:"fromArticleTemplate,omitempty"`
}

func (o *ReadableChangeArticleOptionsInfo) Validate() error {
	return nil
}
