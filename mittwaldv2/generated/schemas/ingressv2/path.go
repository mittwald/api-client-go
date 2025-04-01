package ingressv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "object"
// properties:
//    "path":
//        type: "string"
//    "target":
//        oneOf:
//            - {"$ref": "#/components/schemas/de.mittwald.v1.ingress.TargetUrl"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.ingress.TargetInstallation"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.ingress.TargetUseDefaultPage"}
//            - {"$ref": "#/components/schemas/de.mittwald.v1.ingress.TargetContainer"}
// required:
//    - "path"
//    - "target"

type Path struct {
	Path   string     `json:"path"`
	Target PathTarget `json:"target"`
}

func (o *Path) Validate() error {
	if err := o.Target.Validate(); err != nil {
		return fmt.Errorf("invalid property target: %w", err)
	}
	return nil
}
