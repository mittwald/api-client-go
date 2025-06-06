package leadfyndrv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "category":
//        type: "string"
//    "name":
//        type: "string"
//    "score":
//        type: "number"
//    "unit":
//        type: "string"
//    "value":
//        type: "number"
// required:
//    - "category"
//    - "name"

type Metric struct {
	Category string   `json:"category"`
	Name     string   `json:"name"`
	Score    *float64 `json:"score,omitempty"`
	Unit     *string  `json:"unit,omitempty"`
	Value    *float64 `json:"value,omitempty"`
}

func (o *Metric) Validate() error {
	return nil
}
