package projectv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "redis"
//    - "node"
//    - "container"

type ProjectFeature string

const ProjectFeatureRedis ProjectFeature = "redis"
const ProjectFeatureNode ProjectFeature = "node"
const ProjectFeatureContainer ProjectFeature = "container"

func (e ProjectFeature) Validate() error {
	if e == ProjectFeatureRedis || e == ProjectFeatureNode || e == ProjectFeatureContainer {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
