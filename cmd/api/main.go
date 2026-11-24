// Command api starts the HTTP API using environment configuration.
package main

import "github.com/victortran0904/urlshortener/internal/api"

func main() {
	cfg, err := api.NewConfig()
	if err != nil {
		panic(err)
	}

	app := api.NewEngine(cfg)
	err = app.Start()

	if err != nil {
		panic(err)
	}

}
