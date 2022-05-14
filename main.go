package main

import (
	"net/http"

	controller "notice/controllers"
	service "notice/services"

	"github.com/gin-gonic/gin"
)

func Api(c *gin.Context) {
	c.JSON(200, gin.H{
		"api": "notice",
	})
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./ui/build/index.html")
	r.Static("/public", "./ui/build")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.Use(service.CORS())

	r.GET("/api", Api)

	rApi := r.Group("/api")
	{
		rApi.GET("/notes", controller.GetNotes)
		rNote := rApi.Group("/note")
		{
			rNote.POST("/add", controller.AddNote)
			rNote.GET("/:id", controller.GetNote)
			rNote.PUT("/edit", controller.EditNote)
			rNote.DELETE("/:id", controller.DelNote)
		}
	}

	r.Run("0.0.0.0:9090")
}
