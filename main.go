package main

import (
	"fmt"
	"log"

	"github.com/sethlittleford/budget-bot/version"
)

func main() {
	v, err := version.Version()
	if err != nil {
		log.Fatalln("Version() failed")
	}
	fmt.Println(v)
}
