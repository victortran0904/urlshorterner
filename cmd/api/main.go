package main

import "github.com/victortran0904/urlshortener/internal/api"

func main() {

	app := api.NewEngine()
	err := app.Start()

	if err != nil {
		panic(err)
	}

}
