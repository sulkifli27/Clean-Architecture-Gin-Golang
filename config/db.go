package config

import (
	"fmt"
	"github/com/sulkifli27/golang_api/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func SetupDatabaseConnection() *gorm.DB{
    errEnv := godotenv.Load()
    if errEnv != nil {
        panic("Failed to load env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST") 
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&Local" , dbUser, dbPassword, dbHost, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to create a connection to database")
    }

    db.AutoMigrate(&model.User{}, &model.Book{} )
    return db
}

func CloseDatabaseConnection(db *gorm.DB){
    dbSQL, err := db.DB()
    if err != nil {
        panic("Failed to close connection from database")
    }
    dbSQL.Close()
}