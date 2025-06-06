package userclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "accessTokenRetrievalKey":
//        type: "string"
//        maxLength: 37
//        minLength: 37
//    "userId":
//        type: "string"
//        format: "uuid"
// required:
//    - "userId"
//    - "accessTokenRetrievalKey"
// description: AuthenticateWithAccessTokenRetrievalKeyRequestBody models the JSON body of a 'user-authenticate-with-access-token-retrieval-key' request

// AuthenticateWithAccessTokenRetrievalKeyRequestBody models the JSON body of a 'user-authenticate-with-access-token-retrieval-key' request
type AuthenticateWithAccessTokenRetrievalKeyRequestBody struct {
	AccessTokenRetrievalKey string `json:"accessTokenRetrievalKey"`
	UserId                  string `json:"userId"`
}

func (o *AuthenticateWithAccessTokenRetrievalKeyRequestBody) Validate() error {
	return nil
}
