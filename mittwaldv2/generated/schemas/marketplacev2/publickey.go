package marketplacev2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "algorithm":
//        type: "string"
//        example: "Ed25519"
//    "key":
//        anyOf:
//            - type: "string"
//              format: "byte"
//              description: "raw base64 public key"
//            - type: "string"
//              description: "SPKI, ASN.1 DER serialized, PEM encoded public key"
//    "serial":
//        type: "string"
//        format: "uuid"
// required:
//    - "serial"
//    - "algorithm"
//    - "key"

type PublicKey struct {
	Algorithm string `json:"algorithm"`
	Key       any    `json:"key"`
	Serial    string `json:"serial"`
}

func (o *PublicKey) Validate() error {
	return nil
}
