package stracev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "query":
//        type: "string"
//        description: "The whole DB query."
//        example: "SELECT * FROM my_table;"
//    "stats": {"$ref": "#/components/schemas/de.mittwald.v1.strace.Statistics"}
//    "warnLevel":
//        type: "string"
//        enum:
//            - "NO"
//            - "WARN"
//            - "SEVERE"
//        description: "Alerts when the time, syscall count or occurrence count of this group are abnormal."
// required:
//    - "query"
//    - "stats"
//    - "warnLevel"

type DataDbQueriesItem struct {
	Query     string                     `json:"query"`
	Stats     Statistics                 `json:"stats"`
	WarnLevel DataDbQueriesItemWarnLevel `json:"warnLevel"`
}

func (o *DataDbQueriesItem) Validate() error {
	if err := o.Stats.Validate(); err != nil {
		return fmt.Errorf("invalid property stats: %w", err)
	}
	if err := o.WarnLevel.Validate(); err != nil {
		return fmt.Errorf("invalid property warnLevel: %w", err)
	}
	return nil
}
