package membershipv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "notset"
//    - "owner"
//    - "emailadmin"
//    - "external"

type ProjectRoles string

const ProjectRolesNotset ProjectRoles = "notset"
const ProjectRolesOwner ProjectRoles = "owner"
const ProjectRolesEmailadmin ProjectRoles = "emailadmin"
const ProjectRolesExternal ProjectRoles = "external"

func (e ProjectRoles) Validate() error {
	if e == ProjectRolesNotset || e == ProjectRolesOwner || e == ProjectRolesEmailadmin || e == ProjectRolesExternal {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
