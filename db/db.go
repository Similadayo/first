package db

import (
	"log"
	"os"

	"github.com/Similadayo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitializeDB() {
	Err := godotenv.Load()
	if Err != nil {
		log.Fatal("Error loading .env file")
	}

	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")

	// dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPassword)

	dbUri := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB.AutoMigrate(&models.User{}, &models.BlacklistToken{}, &models.Suspension{},
		&models.Category{}, &models.Product{})
}

func GetDB() (*gorm.DB, error) {
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	dbUri := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}
	return db, nil
}
