package main

import (
	"net/http"
	"os"
	"post-app/models"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	// connect to database
	dsn := db_user + ":" + db_password + "@tcp(127.0.0.1:3306)/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrate all schema
	db.AutoMigrate(&models.User{})

	// create
	db.Create(&models.User{ Username: "damian", Password: "damian", Name: "Damian", Phone: "08123456789"})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}