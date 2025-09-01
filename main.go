package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"techwithprivacy/web/routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		page, err := routes.GetIndex()

		if err != nil {
			log.Fatalf("failed to get index page: %v", err)
			c.String(500, "Error fetching page")
			return
		}

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
