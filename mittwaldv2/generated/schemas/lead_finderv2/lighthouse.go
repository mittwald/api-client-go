package lead_finderv2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "accessibility":
//        type: "number"
//    "bestPractice":
//        type: "number"
//    "performance":
//        type: "number"
//    "seo":
//        type: "number"
// required:
//    - "performance"
//    - "bestPractice"
//    - "accessibility"
//    - "seo"

type Lighthouse struct {
	Accessibility float64 `json:"accessibility"`
	BestPractice  float64 `json:"bestPractice"`
	Performance   float64 `json:"performance"`
	Seo           float64 `json:"seo"`
}

func (o *Lighthouse) Validate() error {
	return nil
}
