package ingressv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "acme":
//        type: "boolean"
// required:
//    - "acme"
// additionalProperties: false

type TlsAcmeDeprecated struct {
	Acme bool `json:"acme"`
}

func (o *TlsAcmeDeprecated) Validate() error {
	return nil
}
