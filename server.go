package main

import (
	"net/http"
	"os"
	"post-app/configs"
	"post-app/entities"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}

	db := configs.DBConfig()

	// migrate all schema
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Post{})
	db.AutoMigrate(&entities.Comment{})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}