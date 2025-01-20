package appv1

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

//This data type was generated from the following JSON schema:
//type: "string"
//enum:
//    - "installation"
//    - "update"
//    - "reconfigure"
//description: "The AppInstallationLifecycle can be used to express a specific point in the AppInstallation Lifecycle, e.g. while installing a new AppInstallation."

//The AppInstallationLifecycle can be used to express a specific point in the AppInstallation Lifecycle, e.g. while installing a new AppInstallation.
type AppInstallationLifecycle string

const AppInstallationLifecycleInstallation AppInstallationLifecycle = "installation"
const AppInstallationLifecycleUpdate AppInstallationLifecycle = "update"
const AppInstallationLifecycleReconfigure AppInstallationLifecycle = "reconfigure"

func (e AppInstallationLifecycle) Validate() error {
	if e == AppInstallationLifecycleInstallation || e == AppInstallationLifecycleUpdate || e == AppInstallationLifecycleReconfigure {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}