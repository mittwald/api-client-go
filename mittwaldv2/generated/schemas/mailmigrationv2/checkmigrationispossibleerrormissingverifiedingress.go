package mailmigrationv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "hostname":
//        type: "string"
//        format: "idn-hostname"
// required:
//    - "hostname"

type CheckMigrationIsPossibleErrorMissingVerifiedIngress struct {
	Hostname string `json:"hostname"`
}

func (o *CheckMigrationIsPossibleErrorMissingVerifiedIngress) Validate() error {
	return nil
}
