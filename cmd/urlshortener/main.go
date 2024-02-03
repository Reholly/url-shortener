package main

import "urlshortener/internal/app"

func main() {
	app := app.New()
	app.Run()
}
