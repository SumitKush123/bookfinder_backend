package main

import (
	"core/config"
	"core/route"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	// Middleware
	config.DbInit()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e = route.InitHttp()

	// Get port from config
	cfg := config.GetConfig()
	port := ":" + cfg.Port

	// Start server
	e.Logger.Fatal(e.Start(port))
}
