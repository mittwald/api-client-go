package pollv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "completed"
//    - "muted"
//    - "ignored"
//    - "new"

type UserPollSettingsStatus string

const UserPollSettingsStatusCompleted UserPollSettingsStatus = "completed"
const UserPollSettingsStatusMuted UserPollSettingsStatus = "muted"
const UserPollSettingsStatusIgnored UserPollSettingsStatus = "ignored"
const UserPollSettingsStatusNew UserPollSettingsStatus = "new"

func (e UserPollSettingsStatus) Validate() error {
	if e == UserPollSettingsStatusCompleted || e == UserPollSettingsStatusMuted || e == UserPollSettingsStatusIgnored || e == UserPollSettingsStatusNew {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
