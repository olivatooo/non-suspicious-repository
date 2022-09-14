package main

import (
	"log"

	"account/actions"
	"account/actions/users"
)

func main() {
	app := actions.App()
	users.Router(app)
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
