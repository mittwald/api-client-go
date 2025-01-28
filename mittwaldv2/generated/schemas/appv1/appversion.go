package appv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "appId":
//        type: "string"
//        format: "uuid"
//    "breakingNote": {"$ref": "#/components/schemas/de.mittwald.v1.app.BreakingNote"}
//    "databases":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.app.DatabaseDependency"}
//    "docRoot":
//        type: "string"
//    "docRootUserEditable":
//        type: "boolean"
//    "externalVersion":
//        type: "string"
//    "id":
//        type: "string"
//        format: "uuid"
//    "internalVersion":
//        type: "string"
//    "recommended":
//        type: "boolean"
//    "requestHandler": {"$ref": "#/components/schemas/de.mittwald.v1.app.RequestHandlerRequirement"}
//    "systemSoftwareDependencies":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.app.SystemSoftwareDependency"}
//        uniqueItems: true
//    "userInputs":
//        type: "array"
//        items: {"$ref": "#/components/schemas/de.mittwald.v1.app.UserInput"}
//        uniqueItems: true
// required:
//    - "id"
//    - "appId"
//    - "externalVersion"
//    - "internalVersion"
//    - "docRoot"
//    - "docRootUserEditable"
// description: "An AppVersion is an officially supported version of an App, containing the necessary and recommended configuration und dependencies."

// An AppVersion is an officially supported version of an App, containing the necessary and recommended configuration und dependencies.
type AppVersion struct {
	AppId                      string                     `json:"appId"`
	BreakingNote               *BreakingNote              `json:"breakingNote,omitempty"`
	Databases                  []DatabaseDependency       `json:"databases,omitempty"`
	DocRoot                    string                     `json:"docRoot"`
	DocRootUserEditable        bool                       `json:"docRootUserEditable"`
	ExternalVersion            string                     `json:"externalVersion"`
	Id                         string                     `json:"id"`
	InternalVersion            string                     `json:"internalVersion"`
	Recommended                *bool                      `json:"recommended,omitempty"`
	RequestHandler             *RequestHandlerRequirement `json:"requestHandler,omitempty"`
	SystemSoftwareDependencies []SystemSoftwareDependency `json:"systemSoftwareDependencies,omitempty"`
	UserInputs                 []UserInput                `json:"userInputs,omitempty"`
}

func (o *AppVersion) Validate() error {
	if err := func() error {
		if o.BreakingNote == nil {
			return nil
		}
		return o.BreakingNote.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property breakingNote: %w", err)
	}
	if err := func() error {
		if o.Databases == nil {
			return nil
		}
		return func() error {
			for i := range o.Databases {
				if err := o.Databases[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property databases: %w", err)
	}
	if err := func() error {
		if o.RequestHandler == nil {
			return nil
		}
		return o.RequestHandler.Validate()
	}(); err != nil {
		return fmt.Errorf("invalid property requestHandler: %w", err)
	}
	if err := func() error {
		if o.SystemSoftwareDependencies == nil {
			return nil
		}
		return func() error {
			for i := range o.SystemSoftwareDependencies {
				if err := o.SystemSoftwareDependencies[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property systemSoftwareDependencies: %w", err)
	}
	if err := func() error {
		if o.UserInputs == nil {
			return nil
		}
		return func() error {
			for i := range o.UserInputs {
				if err := o.UserInputs[i].Validate(); err != nil {
					return fmt.Errorf("item %d is invalid %w", i, err)
				}
			}
			return nil
		}()
	}(); err != nil {
		return fmt.Errorf("invalid property userInputs: %w", err)
	}
	return nil
}
