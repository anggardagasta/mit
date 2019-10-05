package main

import (
	"github.com/anggardagasta/mit/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg := config.Init()

	cfg.Start()
}
