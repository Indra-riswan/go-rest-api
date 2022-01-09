package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupConnectionDb() *gorm.DB {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error Load Env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erorr connection Database")
	}
	//AutoMigrate Database
	// db.AutoMigrate()
	return db

}

func CloseConnectionDatabase(db *gorm.DB) {
	closeDB, err := db.DB()
	if err != nil {
		panic("Failed Close Connectin From Database")
	}
	closeDB.Close()
}
