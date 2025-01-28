package mailv1

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "imapClusterId":
//        type: "string"
//    "mailDirectory":
//        type: "string"
//    "rateLimitId":
//        type: "string"
// required:
//    - "imapClusterId"
//    - "mailDirectory"
//    - "rateLimitId"

type MailsystemSettings struct {
	ImapClusterId string `json:"imapClusterId"`
	MailDirectory string `json:"mailDirectory"`
	RateLimitId   string `json:"rateLimitId"`
}

func (o *MailsystemSettings) Validate() error {
	return nil
}
