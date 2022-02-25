/*
Copyright © 2022 Seth Littleford <seth.littleford@gmail.com>

*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sethlittleford/budget-bot/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file from root")
	}

	cmd.Execute()
}
