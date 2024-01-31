package main

import (
	"fmt"
	"urlshortener/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println(cfg)
}
