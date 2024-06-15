package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const webPort string = ":80"

//go:embed templates
var templatesFS embed.FS

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
		render(c, "test.page.gohtml", &templatesFS)
	})

	fmt.Println("Starting front end service on port 8080")
	mux.Run(webPort)
}

func render(c *gin.Context, t string, templatesFS *embed.FS) {

	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", t))

	templateSlice = append(templateSlice, partials...)

	// Parse the templates from the embedded filesystem
	tmpl, err := template.ParseFS(templatesFS, templateSlice...)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := tmpl.Execute(c.Writer, nil); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}
