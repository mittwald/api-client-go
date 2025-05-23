package leadfyndrv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "address_prefix":
//        type: "string"
//    "city":
//        type: "string"
//        minLength: 1
//    "country_code":
//        type: "string"
//        maxLength: 2
//        minLength: 2
//    "house_number":
//        type: "string"
//        minLength: 1
//    "street":
//        type: "string"
//        minLength: 1
//    "zip":
//        type: "string"
//        minLength: 1

type ContactAddress struct {
	Address_prefix *string `json:"address_prefix,omitempty"`
	City           *string `json:"city,omitempty"`
	Country_code   *string `json:"country_code,omitempty"`
	House_number   *string `json:"house_number,omitempty"`
	Street         *string `json:"street,omitempty"`
	Zip            *string `json:"zip,omitempty"`
}

func (o *ContactAddress) Validate() error {
	return nil
}
