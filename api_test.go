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
	fmt.Printf("%#v\n", data)
	// Output: leagueliveapi.GameData{GameMode:"PRACTICETOOL", GameTime:13.315361022949219, MapName:"Map11", MapNumber:11, MapTerrain:"Default"}
}
