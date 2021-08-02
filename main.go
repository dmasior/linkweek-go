package main

import (
	"linkweek-go/cmd"
	"log"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
