package containerv2

import (
	"errors"
	"fmt"
)

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "command":
//        type: "array"
//        items:
//            type: "string"
//        description: "Defaults to image config on empty"
//        example: ["mysqld"]
//    "description":
//        type: "string"
//        example: "MySQL DB"
//    "entrypoint":
//        type: "array"
//        items:
//            type: "string"
//        description: "Defaults to image config on empty"
//        example: ["docker-entrypoint.sh"]
//    "envs":
//        type: "object"
//        additionalProperties:
//            type: "string"
//        example: {"MYSQL_DATABASE": "my_db", "MYSQL_PASSWORD": "my_password", "MYSQL_ROOT_PASSWORD": "my_root_password", "MYSQL_USER": "my_user"}
//    "image":
//        type: "string"
//        example: "mysql"
//    "ports":
//        type: "array"
//        items:
//            type: "string"
//        example: ["3306:3306/tcp"]
//    "volumes":
//        type: "array"
//        items:
//            type: "string"
//        example: ["data:/var/lib/mysql:ro"]
// required:
//    - "description"
//    - "image"
//    - "ports"

type ServiceDeclareRequest struct {
	Command     []string          `json:"command,omitempty"`
	Description string            `json:"description"`
	Entrypoint  []string          `json:"entrypoint,omitempty"`
	Envs        map[string]string `json:"envs,omitempty"`
	Image       string            `json:"image"`
	Ports       []string          `json:"ports"`
	Volumes     []string          `json:"volumes,omitempty"`
}

func (o *ServiceDeclareRequest) Validate() error {
	if err := func() error {
		if o.Command == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property command: %w", err)
	}
	if err := func() error {
		if o.Entrypoint == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property entrypoint: %w", err)
	}
	if o.Ports == nil {
		return errors.New("property ports is required, but not set")
	}
	if err := func() error {
		if o.Volumes == nil {
			return nil
		}
		return nil
	}(); err != nil {
		return fmt.Errorf("invalid property volumes: %w", err)
	}
	return nil
}
