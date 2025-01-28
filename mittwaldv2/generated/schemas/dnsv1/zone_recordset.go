package dnsv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "cname": {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordCNAME"}
//    "combinedARecords": {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordCombinedA"}
//    "mx": {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordMX"}
//    "srv": {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordSRV"}
//    "txt": {"$ref": "#/components/schemas/de.mittwald.v1.dns.RecordTXT"}
// required:
//    - "combinedARecords"
//    - "cname"
//    - "mx"
//    - "txt"
//    - "srv"

type ZoneRecordSet struct {
	Cname            RecordCNAME     `json:"cname"`
	CombinedARecords RecordCombinedA `json:"combinedARecords"`
	Mx               RecordMX        `json:"mx"`
	Srv              RecordSRV       `json:"srv"`
	Txt              RecordTXT       `json:"txt"`
}

func (o *ZoneRecordSet) Validate() error {
	if err := o.Cname.Validate(); err != nil {
		return fmt.Errorf("invalid property cname: %w", err)
	}
	if err := o.CombinedARecords.Validate(); err != nil {
		return fmt.Errorf("invalid property combinedARecords: %w", err)
	}
	if err := o.Mx.Validate(); err != nil {
		return fmt.Errorf("invalid property mx: %w", err)
	}
	if err := o.Srv.Validate(); err != nil {
		return fmt.Errorf("invalid property srv: %w", err)
	}
	if err := o.Txt.Validate(); err != nil {
		return fmt.Errorf("invalid property txt: %w", err)
	}
	return nil
}
