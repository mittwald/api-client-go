package stracev2

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "kernelMs":
//        type: "number"
//        description: "Elapsed kernel space time in milliseconds."
//        example: 1.2345
//    "occurrences":
//        type: "integer"
//        description: "The number of times this group occurred."
//        example: 25
//    "syscallCount":
//        type: "integer"
//        description: "Syscall count."
//        example: 4321
//    "userspaceMs":
//        type: "number"
//        description: "Elapsed user space time in milliseconds."
//        example: 1.2345
// required:
//    - "kernelMs"
//    - "userspaceMs"
//    - "syscallCount"
//    - "occurrences"

type Statistics struct {
	KernelMs     float64 `json:"kernelMs"`
	Occurrences  int64   `json:"occurrences"`
	SyscallCount int64   `json:"syscallCount"`
	UserspaceMs  float64 `json:"userspaceMs"`
}

func (o *Statistics) Validate() error {
	return nil
}
