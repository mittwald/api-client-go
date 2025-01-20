package user

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "object"
//properties:
//    "expires":
//        type: "string"
//        format: "date-time"
//        description: "The expiration date of the token."
//    "refreshToken":
//        type: "string"
//        description: "Refresh token to refresh your access token even after it has expired."
//    "token":
//        type: "string"
//        description: "Public token to identify yourself against the api gateway."
//required:
//    - "token"
//    - "refreshToken"
//    - "expires"

//
type AuthenticateMFAResponse struct {
	Expires      string `json:"expires"`
	RefreshToken string `json:"refreshToken"`
	Token        string `json:"token"`
}

func (o *AuthenticateMFAResponse) Validate() error {
	return nil
}
