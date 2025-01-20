package relocation

import (
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/directusv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "additionalServices":
//        type: "object"
//        properties:
//            "dataCompare":
//                type: "string"
//                enum:
//                    - "default"
//                    - "additionalCompare"
//        required:
//            - "dataCompare"
//    "allDomains":
//        type: "boolean"
//        description: "Should all project releated domains should be transferred to mittwald?"
//    "allowPasswordChange":
//        type: "boolean"
//        description: "Has to be true. Do you accept that our mittwald team can change and get password from your current provider?"
//    "articleType":
//        type: "string"
//        enum:
//            - "cms-hosting"
//            - "cms-hosting-express"
//            - "onlineshop"
//            - "onlineshop-express"
//        description: "Type of the article you want to relocate."
//    "contact":
//        type: "object"
//        properties:
//            "email":
//                type: "string"
//                minLength: 1
//                format: "email"
//            "firstName":
//                type: "string"
//                minLength: 1
//            "lastName":
//                type: "string"
//                minLength: 1
//            "phoneNumber":
//                type: "string"
//                pattern: "|^\\+([0-9]{2,3}|1)-[0-9]{2,5}-[0-9]+$"
//        required:
//            - "firstName"
//            - "lastName"
//            - "email"
//    "domains":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.directus.Domain"}
//        description: "List of domains which should be transferred (when allDomains is not checked)."
//    "notes":
//        type: "string"
//        description: "Anything our customer service needs to know for the relocation process."
//    "prices":
//        type: "object"
//        properties:
//            "positions":
//                type: "array"
//                items:
//                    type: "object"
//                    properties:
//                        "name":
//                            type: "string"
//                            minLength: 1
//                        "price":
//                            type: "number"
//                            minimum: 0
//                    required:
//                        - "name"
//                        - "price"
//            "total":
//                type: "number"
//                minimum: 0
//        required:
//            - "positions"
//            - "total"
//    "provider":
//        type: "object"
//        properties:
//            "loginUrl":
//                type: "string"
//                minLength: 1
//                description: "Url to the control panel of the provider"
//            "name":
//                oneOf:
//                    - type: "string"
//                      minLength: 1
//                    - type: "string"
//                      enum:
//                        - "1und1"
//                        - "strato"
//                description: "Name of your provider"
//            "password":
//                type: "string"
//                minLength: 1
//            "sourceAccount":
//                type: "string"
//                minLength: 1
//                description: "Which account of your provider should be moved?"
//            "userName":
//                type: "string"
//                minLength: 1
//                description: "Login name to your provider"
//        required:
//            - "name"
//            - "loginUrl"
//            - "userName"
//            - "password"
//            - "sourceAccount"
//    "target":
//        type: "object"
//        properties:
//            "organisation":
//                type: "string"
//                minLength: 1
//                description: "Your customer or organisation number"
//            "product":
//                type: "string"
//                oneOf:
//                    - type: "string"
//                      minLength: 1
//                    - type: "string"
//                      enum:
//                        - "Space-Server"
//                        - "proSpace"
//                        - "Agentur-Server"
//                        - "CMS-Hosting"
//                        - "Shop-Hosting"
//                description: "Help our customer service finding your target account"
//            "projectName":
//                type: "string"
//                minLength: 1
//                description: "In which p-account or short project id your project should be moved."
//            "system":
//                type: "string"
//                enum:
//                    - "kc"
//                    - "mstudio"
//                description: "Which mittwald system does the targetProject use?"
//        required:
//            - "organisation"
//            - "projectName"
//            - "system"
//            - "product"
//    "userId":
//        oneOf:
//            - type: "string"
//              minLength: 1
//              format: "uuid"
//            - type: "string"
//required:
//    - "articleType"
//    - "additionalServices"
//    - "prices"
//    - "provider"
//    - "contact"
//    - "target"
//    - "allowPasswordChange"

//
type CreateRelocationRequestBody struct {
	AdditionalServices  CreateRelocationRequestBodyAdditionalServices `json:"additionalServices"`
	AllDomains          *bool                                         `json:"allDomains,omitempty"`
	AllowPasswordChange bool                                          `json:"allowPasswordChange"`
	ArticleType         CreateRelocationRequestBodyArticleType        `json:"articleType"`
	Contact             CreateRelocationRequestBodyContact            `json:"contact"`
	Domains             []directusv1.Domain                           `json:"domains,omitempty"`
	Notes               *string                                       `json:"notes,omitempty"`
	Prices              CreateRelocationRequestBodyPrices             `json:"prices"`
	Provider            CreateRelocationRequestBodyProvider           `json:"provider"`
	Target              CreateRelocationRequestBodyTarget             `json:"target"`
	UserId              *CreateRelocationRequestBodyUserID            `json:"userId,omitempty"`
}

func (o *CreateRelocationRequestBody) Validate() error {
	if err := o.AdditionalServices.Validate(); err != nil {
		return fmt.Errorf("invalid property additionalServices: %w", err)
	}
	if err := o.ArticleType.Validate(); err != nil {
		return fmt.Errorf("invalid property articleType: %w", err)
	}
	if err := o.Contact.Validate(); err != nil {
		return fmt.Errorf("invalid property contact: %w", err)
	}
	if err := func() error {
		if o.Domains == nil {
			return nil
		}
		return func() error {
			for i := range o.Domains {
				if err := o.Domains[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property domains: %w", err)
	}
	if err := o.Prices.Validate(); err != nil {
		return fmt.Errorf("invalid property prices: %w", err)
	}
	if err := o.Provider.Validate(); err != nil {
		return fmt.Errorf("invalid property provider: %w", err)
	}
	if err := o.Target.Validate(); err != nil {
		return fmt.Errorf("invalid property target: %w", err)
	}
	if err := func() error {
		if o.UserId == nil {
			return nil
		}
		return o.UserId.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property userId: %w", err)
	}
	return nil
}