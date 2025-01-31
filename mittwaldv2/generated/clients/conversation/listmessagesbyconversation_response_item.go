package conversation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2/generated/schemas/conversationv1"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.conversation.Message"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.conversation.StatusUpdate"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.conversation.ServiceRequest"}

type ListMessagesByConversationResponseItem struct {
	AlternativeMessage        *conversationv1.Message
	AlternativeStatusUpdate   *conversationv1.StatusUpdate
	AlternativeServiceRequest *conversationv1.ServiceRequest
}

func (a *ListMessagesByConversationResponseItem) MarshalJSON() ([]byte, error) {
	if a.AlternativeMessage != nil {
		return json.Marshal(a.AlternativeMessage)
	}
	if a.AlternativeStatusUpdate != nil {
		return json.Marshal(a.AlternativeStatusUpdate)
	}
	if a.AlternativeServiceRequest != nil {
		return json.Marshal(a.AlternativeServiceRequest)
	}
	return []byte("null"), nil
}

func (a *ListMessagesByConversationResponseItem) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeMessage conversationv1.Message
	if err := dec.Decode(&alternativeMessage); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeMessage.Validate(); vErr == nil {
			a.AlternativeMessage = &alternativeMessage
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeStatusUpdate conversationv1.StatusUpdate
	if err := dec.Decode(&alternativeStatusUpdate); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeStatusUpdate.Validate(); vErr == nil {
			a.AlternativeStatusUpdate = &alternativeStatusUpdate
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeServiceRequest conversationv1.ServiceRequest
	if err := dec.Decode(&alternativeServiceRequest); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeServiceRequest.Validate(); vErr == nil {
			a.AlternativeServiceRequest = &alternativeServiceRequest
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *ListMessagesByConversationResponseItem) Validate() error {
	if a.AlternativeMessage != nil {
		return a.AlternativeMessage.Validate()
	}
	if a.AlternativeStatusUpdate != nil {
		return a.AlternativeStatusUpdate.Validate()
	}
	if a.AlternativeServiceRequest != nil {
		return a.AlternativeServiceRequest.Validate()
	}
	return errors.New("no alternative set")
}
