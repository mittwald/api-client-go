package databasev1

import "errors"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "collations":
//        type: "array"
//        items:
//            type: "string"
//        example: ["utf8_general_ci", "utf8_bin", "utf8_unicode_ci"]
//    "name":
//        type: "string"
//        example: "utf8"
//    "versionId":
//        type: "string"
//        example: "mysql57"
// required:
//    - "name"
//    - "collations"
//    - "versionId"

type MySqlCharacterSettings struct {
	Collations []string `json:"collations"`
	Name       string   `json:"name"`
	VersionId  string   `json:"versionId"`
}

func (o *MySqlCharacterSettings) Validate() error {
	if o.Collations == nil {
		return errors.New("property collations is required, but not set")
	}
	return nil
}
