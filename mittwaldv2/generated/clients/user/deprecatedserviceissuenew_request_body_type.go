package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "feedback"
//    - "bug"

type DeprecatedServiceIssueNewRequestBodyType string

const DeprecatedServiceIssueNewRequestBodyTypeFeedback DeprecatedServiceIssueNewRequestBodyType = "feedback"
const DeprecatedServiceIssueNewRequestBodyTypeBug DeprecatedServiceIssueNewRequestBodyType = "bug"

func (e DeprecatedServiceIssueNewRequestBodyType) Validate() error {
	if e == DeprecatedServiceIssueNewRequestBodyTypeFeedback || e == DeprecatedServiceIssueNewRequestBodyTypeBug {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
