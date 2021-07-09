package leagueliveapi_test

import (
	"fmt"

	"github.com/Jonchun/leagueliveapi"
)

func ExampleNew() {
	a := leagueliveapi.New()
	data, err := a.GetGameData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
