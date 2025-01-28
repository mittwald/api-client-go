package sshuserv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "comment":
//        type: "string"
//    "key":
//        type: "string"
// required:
//    - "comment"
//    - "key"
// description: "A representation of an ssh-public-key."

// A representation of an ssh-public-key.
type PublicKey struct {
	Comment string `json:"comment"`
	Key     string `json:"key"`
}

func (o *PublicKey) Validate() error {
	return nil
}
