package main

import (
	"github.com/ceejay1000/blog-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/blog/:id", handlers.GetBlogPost)
	router.GET("/blog", handlers.GetBlogPosts)
	router.POST("/blog", handlers.CreateBlog)
	router.PUT("/blog/:id", handlers.UpdateBlog)
	router.DELETE("/blog/:id", handlers.DeleteBlog)

	router.Run(":2000")
}
