package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "completed"
//    - "muted"
//    - "ignored"

type PostPollStatusRequestBodyStatus string

const PostPollStatusRequestBodyStatusCompleted PostPollStatusRequestBodyStatus = "completed"
const PostPollStatusRequestBodyStatusMuted PostPollStatusRequestBodyStatus = "muted"
const PostPollStatusRequestBodyStatusIgnored PostPollStatusRequestBodyStatus = "ignored"

func (e PostPollStatusRequestBodyStatus) Validate() error {
	if e == PostPollStatusRequestBodyStatusCompleted || e == PostPollStatusRequestBodyStatusMuted || e == PostPollStatusRequestBodyStatusIgnored {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
