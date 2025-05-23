package marketplaceclientv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "accessTokenRetrievalKey":
//        type: "string"
//    "userId":
//        type: "string"
//        format: "uuid"
// required:
//    - "accessTokenRetrievalKey"
//    - "userId"

type CreateRetrievalKeyResponse struct {
	AccessTokenRetrievalKey string `json:"accessTokenRetrievalKey"`
	UserId                  string `json:"userId"`
}

func (o *CreateRetrievalKeyResponse) Validate() error {
	return nil
}
