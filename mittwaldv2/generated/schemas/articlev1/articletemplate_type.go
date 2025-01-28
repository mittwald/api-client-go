package articlev1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "miscellaneous"
//    - "base"
//    - "additional"
//    - "modifier"
//    - "setup_fee"

type ArticleTemplateType string

const ArticleTemplateTypeMiscellaneous ArticleTemplateType = "miscellaneous"
const ArticleTemplateTypeBase ArticleTemplateType = "base"
const ArticleTemplateTypeAdditional ArticleTemplateType = "additional"
const ArticleTemplateTypeModifier ArticleTemplateType = "modifier"
const ArticleTemplateTypeSetupfee ArticleTemplateType = "setup_fee"

func (e ArticleTemplateType) Validate() error {
	if e == ArticleTemplateTypeMiscellaneous || e == ArticleTemplateTypeBase || e == ArticleTemplateTypeAdditional || e == ArticleTemplateTypeModifier || e == ArticleTemplateTypeSetupfee {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
