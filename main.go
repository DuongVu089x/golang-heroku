package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	// router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

	// router.Run(":" + port)



	http.HandleFunc("/api/v1/update", update)

	fmt.Println("Listenning on port", port, ".")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func update(w http.ResponseWriter, r *http.Request) {

}