package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/ceejay1000/blog-api/database"
	"github.com/ceejay1000/blog-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBlogPost(ctx *gin.Context) {

	id := ctx.Param("id")

	blog := new(models.Blog)

	// parsedId, err := strconv.ParseInt(id, 10, 64)

	// if err != nil {
	// 	log.Println("Cannot parse Id")
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "An error occured",
	// 	})
	// 	return
	// }

	// (*blog).ID = uint(parsedId)
	// queryResult := database.InitializeDB().First(blog)
	queryResult := database.InitializeDB().First(blog, id)

	if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {

		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Blog with ID '" + id + "' not found",
		})
		return
	}

	ctx.JSON(http.StatusFound, blog)
}

func GetBlogPosts(ctx *gin.Context) {

	var blogPosts []*models.Blog

	queryResult := database.InitializeDB().Find(&blogPosts)

	if errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {

		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "No blog post found",
		})
		return
	}

	if len(blogPosts) == 0 {

		ctx.JSON(http.StatusOK, gin.H{
			"posts": blogPosts,
		})
		return
	}

	ctx.JSON(http.StatusFound, blogPosts)
}

func GetBlogByName(ctx *gin.Context) {

	// blogPost := new(models.Blog)s

	var blogPosts []*models.Blog
	title := ctx.Param("title")

	if title == "" {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog title not specified",
		})

	}

	database.InitializeDB().Where("title = ?", title).First(blogPosts)
	// database.InitializeDB().Where(&models.Blog{Title: "Test", Content: "Some content"}, "title", "content").First(blogPosts)

	if len(blogPosts) == 0 {

		ctx.JSON(http.StatusOK, gin.H{
			"posts": blogPosts,
		})
		return
	}

	ctx.JSON(http.StatusFound, blogPosts)

}

func CreateBlog(ctx *gin.Context) {

	newBlog := new(models.Blog)

	err := ctx.ShouldBindJSON(newBlog)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid fields passed",
		})
	}

	result := database.InitializeDB().Create(newBlog)

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Blog already created",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Blog post created successfully",
	})
}

func UpdateBlog(ctx *gin.Context) {

	id := ctx.Param("id")

	blog := new(models.Blog)

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID passed",
		})
		return
	}

	blog.ID = uint(parsedId)

	if err := ctx.ShouldBindJSON(blog); err != nil {
		log.Fatal("Cannot bind data")

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "An error occured please try again",
		})
		return
	}

	// result := database.InitializeDB().Save(blog)
	// result := database.InitializeDB().Model(&models.Blog{}).Where("id = ?1", parsedId).Updates(blog);
	result := database.InitializeDB().Model(blog).Updates(blog)

	if result.RowsAffected > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Blog post updated successfully",
		})
	}
}

func DeleteBlog(ctx *gin.Context) {

	id := ctx.Param("id")

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Println("Invalid ID passed")
	}

	result := database.InitializeDB().Delete(&models.Blog{}, parsedId)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete blog, please try again",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
	})

}
