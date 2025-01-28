package contractv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "amount":
//        type: "integer"
//        minimum: 1
//        example: 1
//    "articleTemplateId":
//        type: "string"
//        example: "a1b8f0e9-904f-4716-a1c0-81ccf5342a56"
//    "description":
//        type: "string"
//        example: "Musterbeschreibung"
//    "id":
//        type: "string"
//        example: "a1b8f0e9-904f-4716-a1c0-81ccf5342a56"
//    "name":
//        type: "string"
//        example: "Musterartikel"
//    "unitPrice": {"$ref": "#/components/schemas/de.mittwald.v1.contract.Price"}
// required:
//    - "id"
//    - "name"
//    - "articleTemplateId"
//    - "amount"
//    - "unitPrice"

type Article struct {
	Amount            int64   `json:"amount"`
	ArticleTemplateId string  `json:"articleTemplateId"`
	Description       *string `json:"description,omitempty"`
	Id                string  `json:"id"`
	Name              string  `json:"name"`
	UnitPrice         Price   `json:"unitPrice"`
}

func (o *Article) Validate() error {
	if err := o.UnitPrice.Validate(); err != nil {
		return fmt.Errorf("invalid property unitPrice: %w", err)
	}
	return nil
}
