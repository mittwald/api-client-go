package mittwaldv2_test

import (
	"context"
	"fmt"

	"github.com/mittwald/api-client-go/mittwaldv2"
	"github.com/mittwald/api-client-go/mittwaldv2/generated/clients/userclientv2"
)

func ExampleNew_accessTokenFromEnvironment() {
	ctx := context.Background()

	client, err := mittwaldv2.New(ctx, mittwaldv2.WithAccessTokenFromEnv())
	if err != nil {
		panic(err)
	}

	me, _, err := client.User().GetUser(ctx, userclientv2.GetUserRequest{UserID: "self"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!", me.Person.FirstName)
}
