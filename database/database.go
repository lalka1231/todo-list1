package database

import (
    "fmt"
    "log"
    "os"
    
    "github.com/lalka1231/todo-list1/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        dbHost, dbUser, dbPassword, dbName, dbPort)
    
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    log.Println("Database connected successfully")
    
    DB.AutoMigrate(&models.Task{})
    log.Println("Database migrated")
}
