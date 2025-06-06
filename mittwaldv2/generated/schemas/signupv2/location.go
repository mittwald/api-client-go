package signupv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "city":
//        type: "string"
//        example: "Espelkamp"
//    "country":
//        type: "string"
//        example: "DE"
//    "ipAddress":
//        type: "string"
//        format: "ipv4"

type Location struct {
	City      *string `json:"city,omitempty"`
	Country   *string `json:"country,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
}

func (o *Location) Validate() error {
	return nil
}
