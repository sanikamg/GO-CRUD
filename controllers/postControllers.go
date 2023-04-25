package controllers

import (
	"go_crud/initializers"
	"go_crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// get data of request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}
	//return it

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"post": posts,
	})
}

func PostUpdate(c *gin.Context) {
	//get the id off url
	id := c.Param("id")

	//get the data of req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	//respond with it
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	//get the post
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//respond with them
	c.Status(200)
}
