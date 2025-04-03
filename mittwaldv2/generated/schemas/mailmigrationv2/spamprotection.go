package mailmigrationv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "active":
//        type: "boolean"
//    "deleteSensitivity":
//        type: "integer"
//    "folder":
//        type: "integer"
//        enum:
//            - 0
//            - 1
//        description: "SPAM_FOLDER_INBOX_UNSPECIFIED = 0 SPAM_FOLDER_SPAM = 1"
//        default: 0
//    "keepDays":
//        type: "integer"
//    "relocateSensitivity":
//        type: "integer"
// required:
//    - "active"
//    - "folder"
//    - "keepDays"

type SpamProtection struct {
	Active              bool   `json:"active"`
	DeleteSensitivity   *int64 `json:"deleteSensitivity,omitempty"`
	Folder              int64  `json:"folder"`
	KeepDays            int64  `json:"keepDays"`
	RelocateSensitivity *int64 `json:"relocateSensitivity,omitempty"`
}

func (o *SpamProtection) Validate() error {
	return nil
}
