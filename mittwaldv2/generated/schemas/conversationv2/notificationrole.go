package conversationv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "customer_owner"
//    - "customer_accountant"
//    - "customer_member"
//    - "project_owner"
//    - "project_emailadmin"
//    - "project_external"

type NotificationRole string

const NotificationRoleCustomerowner NotificationRole = "customer_owner"
const NotificationRoleCustomeraccountant NotificationRole = "customer_accountant"
const NotificationRoleCustomermember NotificationRole = "customer_member"
const NotificationRoleProjectowner NotificationRole = "project_owner"
const NotificationRoleProjectemailadmin NotificationRole = "project_emailadmin"
const NotificationRoleProjectexternal NotificationRole = "project_external"

func (e NotificationRole) Validate() error {
	if e == NotificationRoleCustomerowner || e == NotificationRoleCustomeraccountant || e == NotificationRoleCustomermember || e == NotificationRoleProjectowner || e == NotificationRoleProjectemailadmin || e == NotificationRoleProjectexternal {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}
