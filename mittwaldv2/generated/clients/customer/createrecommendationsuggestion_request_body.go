package customer

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "suggestion":
//        type: "string"
//        example: "a mittwald t-shirt"
//required:
//    - "suggestion"

type CreateRecommendationSuggestionRequestBody struct {
	Suggestion string `json:"suggestion"`
}

func (o *CreateRecommendationSuggestionRequestBody) Validate() error {
	return nil
}
