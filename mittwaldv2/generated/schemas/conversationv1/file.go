package conversationv1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//oneOf:
//    - {"$ref": "#/components/schemas/de.mittwald.v1.conversation.RequestedFile"}
//    - {"$ref": "#/components/schemas/de.mittwald.v1.conversation.UploadedFile"}

type File struct {
	AlternativeRequestedFile *RequestedFile
	AlternativeUploadedFile  *UploadedFile
}

func (a *File) MarshalJSON() ([]byte, error) {
	if a.AlternativeRequestedFile != nil {
		return json.Marshal(a.AlternativeRequestedFile)
	}
	if a.AlternativeUploadedFile != nil {
		return json.Marshal(a.AlternativeUploadedFile)
	}
	return []byte("null"), nil
}

func (a *File) UnmarshalJSON(input []byte) error {
	reader := bytes.NewReader(input)
	decodedAtLeastOnce := false
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	reader.Reset(input)
	var alternativeRequestedFile RequestedFile
	if err := dec.Decode(&alternativeRequestedFile); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeRequestedFile.Validate(); vErr == nil {
			a.AlternativeRequestedFile = &alternativeRequestedFile
			decodedAtLeastOnce = true
		}
	}

	reader.Reset(input)
	var alternativeUploadedFile UploadedFile
	if err := dec.Decode(&alternativeUploadedFile); err == nil {
		//subtype: *generator.ReferenceType
		if vErr := alternativeUploadedFile.Validate(); vErr == nil {
			a.AlternativeUploadedFile = &alternativeUploadedFile
			decodedAtLeastOnce = true
		}
	}

	if !decodedAtLeastOnce {
		return fmt.Errorf("could not unmarshal into any alternative for type %T", a)
	}
	return nil
}

func (a *File) Validate() error {
	if a.AlternativeRequestedFile != nil {
		return a.AlternativeRequestedFile.Validate()
	}
	if a.AlternativeUploadedFile != nil {
		return a.AlternativeUploadedFile.Validate()
	}
	return errors.New("no alternative set")
}