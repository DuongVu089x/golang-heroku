package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"

	//"gitlab.ghn.vn/common-projects/go-sdk/sdk"
)

const (
	// TOKEN telegram
	TOKEN = "904350232:AAHGK4iwOaKlr1ujT7FDdKeHLzYIwEQASVs"
	// URL telegram
	URL = "https://api.telegram.org/bot"
	// PORT local
	// PORT = os.Getenv("PORT")
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/demo", func(c echo.Context) error {
		log.Print(c)
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":"+port))
}

