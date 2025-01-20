package conversation

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "accessToken":
//        type: "string"
//    "expiresAt":
//        type: "string"
//        format: "date-time"
//required:
//    - "accessToken"
//    - "expiresAt"

//
type GetFileAccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   string `json:"expiresAt"`
}

func (o *GetFileAccessTokenResponse) Validate() error {
	return nil
}
