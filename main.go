package main

import (
	"fmt"
	test_client "github.com/pavelRostovtsev/test_api/pkg/test_any_client"
)

func main() {
	input := test_client.Input{
		AnyData: `any`,
	}
	testClient := test_client.NewClient(input)

	anyData := testClient.GetPreparedAnyData()

	fmt.Print(anyData)
}
