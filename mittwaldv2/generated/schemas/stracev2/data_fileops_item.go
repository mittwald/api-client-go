package stracev2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "filename":
//        type: "string"
//        example: "my_file.php"
//    "filepath":
//        type: "string"
//        example: "/html/my-project"
//    "stats": {"$ref": "#/components/schemas/de.mittwald.v1.strace.Statistics"}
//    "warnLevel":
//        type: "string"
//        enum:
//            - "NO"
//            - "WARN"
//            - "SEVERE"
//        description: "Alerts when the time, syscall count or occurrence count of this group are abnormal."
// required:
//    - "stats"
//    - "warnLevel"

type DataFileOpsItem struct {
	Filename  *string                  `json:"filename,omitempty"`
	Filepath  *string                  `json:"filepath,omitempty"`
	Stats     Statistics               `json:"stats"`
	WarnLevel DataFileOpsItemWarnLevel `json:"warnLevel"`
}

func (o *DataFileOpsItem) Validate() error {
	if err := o.Stats.Validate(); err != nil {
		return fmt.Errorf("invalid property stats: %w", err)
	}
	if err := o.WarnLevel.Validate(); err != nil {
		return fmt.Errorf("invalid property warnLevel: %w", err)
	}
	return nil
}
