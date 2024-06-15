package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const webPort string = ":80"


func main() {

	mux := gin.Default()

	// setup cross-origin resourse sharing middleware
	mux.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// routes
	mux.GET("/", func(c *gin.Context) {
		render(c, "test.page.gohtml")
	})

	fmt.Println("Starting front end service on port 8080")
	mux.Run(webPort)
}

func render(c *gin.Context, t string) {

	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", t))

	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := tmpl.Execute(c.Writer, nil); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
