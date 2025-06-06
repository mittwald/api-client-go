package fileclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "avatar"
//    - "conversation"
// example: "avatar"

type DeprecatedGetFileTypeRulesRequestPathName string

const DeprecatedGetFileTypeRulesRequestPathNameAvatar DeprecatedGetFileTypeRulesRequestPathName = "avatar"
const DeprecatedGetFileTypeRulesRequestPathNameConversation DeprecatedGetFileTypeRulesRequestPathName = "conversation"

func (e DeprecatedGetFileTypeRulesRequestPathName) Validate() error {
	if e == DeprecatedGetFileTypeRulesRequestPathNameAvatar || e == DeprecatedGetFileTypeRulesRequestPathNameConversation {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
