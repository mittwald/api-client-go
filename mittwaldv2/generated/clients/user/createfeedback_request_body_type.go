package user

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "feedback"
//    - "bug"
// description: "Type of feedback."

// Type of feedback.
type CreateFeedbackRequestBodyType string

const CreateFeedbackRequestBodyTypeFeedback CreateFeedbackRequestBodyType = "feedback"
const CreateFeedbackRequestBodyTypeBug CreateFeedbackRequestBodyType = "bug"

func (e CreateFeedbackRequestBodyType) Validate() error {
	if e == CreateFeedbackRequestBodyTypeFeedback || e == CreateFeedbackRequestBodyTypeBug {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
