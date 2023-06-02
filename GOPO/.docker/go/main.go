package main

// A good tutorial for go and databases: https://eli.thegreenplace.net/2021/accessing-postgresql-databases-in-go/

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	httpport := os.Getenv("HTTP_PORT")
	ConnectInit()
	gin.SetMode("debug")
	fmt.Println("Server started at localhost:" + httpport)
	router := gin.Default()
	router.GET("/", getIdeas)
	router.POST("/", postIdea)
	router.LoadHTMLGlob("templates/*")
	router.Run(":" + httpport)
}

func getIdeas(c *gin.Context) {
	renderIdeasPage(c, "")
}

func renderIdeasPage(c *gin.Context, errMsg string) {
	Ideas, err := queryAllIdeas()
	if err != nil {

		c.HTML(http.StatusOK, "ideas.html", gin.H{
			"Title": "Ideas",
			"Error": "Failed to get ideas: " + err.Error(),
			"Ideas": Ideas,
		})
		return
	}
	c.HTML(http.StatusOK, "ideas.html", gin.H{
		"Title": "Ideas",
		"Error": errMsg,
		"Ideas": Ideas,
	})
}

func postIdea(c *gin.Context) {
	cat := c.PostForm("Category")
	desc := c.PostForm("Description")
	if cat == "" {
		renderIdeasPage(c, "You must select a category!")
		return
	}
	if desc == "" {
		renderIdeasPage(c, "You must set a description!")
		return
	}
	err := queryAddIdea(cat, desc)
	if err != nil {
		renderIdeasPage(c, err.Error())
	}
	renderIdeasPage(c, "")
}
