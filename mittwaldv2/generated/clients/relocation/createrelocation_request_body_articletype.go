package relocation

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "cms-hosting"
//    - "cms-hosting-express"
//    - "onlineshop"
//    - "onlineshop-express"
// description: "Type of the article you want to relocate."

// Type of the article you want to relocate.
type CreateRelocationRequestBodyArticleType string

const CreateRelocationRequestBodyArticleTypeCmsHosting CreateRelocationRequestBodyArticleType = "cms-hosting"
const CreateRelocationRequestBodyArticleTypeCmsHostingExpress CreateRelocationRequestBodyArticleType = "cms-hosting-express"
const CreateRelocationRequestBodyArticleTypeOnlineshop CreateRelocationRequestBodyArticleType = "onlineshop"
const CreateRelocationRequestBodyArticleTypeOnlineshopExpress CreateRelocationRequestBodyArticleType = "onlineshop-express"

func (e CreateRelocationRequestBodyArticleType) Validate() error {
	if e == CreateRelocationRequestBodyArticleTypeCmsHosting || e == CreateRelocationRequestBodyArticleTypeCmsHostingExpress || e == CreateRelocationRequestBodyArticleTypeOnlineshop || e == CreateRelocationRequestBodyArticleTypeOnlineshopExpress {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
