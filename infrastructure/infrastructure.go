package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/opencrypter/api/domain"
	"os"
)

func init() {
	godotenv.Load("../.env")
}

// Opens a new repository connection.
func openDb() *gorm.DB {
	var parameters string

	if os.Getenv("GIN_MODE") != "release" {
		parameters += "sslmode=disable"
	}

	parameters += " host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" dbname=" + os.Getenv("DB_NAME") +
		" password=" + os.Getenv("DB_PASSWORD")

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
	)
}
