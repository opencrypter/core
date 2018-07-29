package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/opencrypter/api/domain"
	"os"
)

func init() {
	MigrateDb()
}

// Opens a new repository connection.
func openDb() *gorm.DB {
	var parameters string
	var ok bool
	var host string
	var port string
	var dbName string
	var user string
	var password string

	if host, ok = os.LookupEnv("DB_HOST"); !ok {
		host = "127.0.0.1"
	}
	if port, ok = os.LookupEnv("DB_PORT"); !ok {
		port = "5432"
	}
	if dbName, ok = os.LookupEnv("DB_NAME"); !ok {
		dbName = "opencrypter"
	}
	if user, ok = os.LookupEnv("DB_USER"); !ok {
		user = "postgres"
	}
	if os.Getenv("GIN_MODE") != "release" {
		parameters += "sslmode=disable "
	}
	password = os.Getenv("DB_PASSWORD")

	parameters += fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbName, password)

	db, err := gorm.Open("postgres", parameters)
	if err != nil {
		panic(err)
	}

	return db
}

// Run auto migration for given models
func MigrateDb() {
	repository := openDb()
	defer repository.Close()

	repository.AutoMigrate(
		domain.Device{},
		domain.Account{},
	)
}
