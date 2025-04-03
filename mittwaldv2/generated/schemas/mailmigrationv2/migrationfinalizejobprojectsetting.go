package mailmigrationv2

import "errors"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "blacklistEntries":
//        type: "array"
//        items:
//            type: "string"
//    "whitelistEntries":
//        type: "array"
//        items:
//            type: "string"
// required:
//    - "blacklistEntries"
//    - "whitelistEntries"

type MigrationFinalizeJobProjectSetting struct {
	BlacklistEntries []string `json:"blacklistEntries"`
	WhitelistEntries []string `json:"whitelistEntries"`
}

func (o *MigrationFinalizeJobProjectSetting) Validate() error {
	if o.BlacklistEntries == nil {
		return errors.New("property blacklistEntries is required, but not set")
	}
	if o.WhitelistEntries == nil {
		return errors.New("property whitelistEntries is required, but not set")
	}
	return nil
}
