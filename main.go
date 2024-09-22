package main

import (
	"fmt"
	"groupie-tracker/api"
)

func main() {
	fmt.Println(string(api.FetchAPI("https://groupietrackers.herokuapp.com/api/artists")))
}
