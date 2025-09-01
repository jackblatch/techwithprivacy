package main

import (
	"fmt"
	"os"
	"techwithprivacy/components"
	"techwithprivacy/markdown"
	"techwithprivacy/pages"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		markdownContent, err := os.ReadFile("content.md")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		html := markdown.ToHTML(markdownContent)
		component := pages.Index(string(html))
		page := components.RootLayout("Tech with Privacy", component)
		c.Status(200)
		c.Header("Content-Type", "text/html; charset=utf-8")

		err = page.Render(c.Request.Context(), c.Writer)
		if err != nil {
			c.String(500, err.Error())
		}
	})

	return r
}

func main() {
	r := setupRouter()
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
