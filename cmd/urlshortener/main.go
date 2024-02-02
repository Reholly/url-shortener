package main

import (
	"fmt"
	"urlshortener/internal/config"
	"urlshortener/internal/storage"
)

func main() {
	cfg := config.LoadConfig()

	db := storage.New(cfg.ConnectionString)
	fmt.Println(db)
}
