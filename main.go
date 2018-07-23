package main

import (
	"github.com/joho/godotenv"
	"github.com/opencrypter/api/infrastructure"
	"github.com/opencrypter/api/ui"
)

func init() {
	godotenv.Load()
	infrastructure.MigrateDb()
}

func main() {
	ui.NewRouter().Run()
}
