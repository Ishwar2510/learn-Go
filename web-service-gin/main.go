package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "10", Title: "Blue train", Artist: "john", Price: 56.9},
	{ID: "20", Title: "Blue train", Artist: "john", Price: 56.9},
	{ID: "30", Title: "Blue train", Artist: "john", Price: 56.9},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbum(c *gin.Context) {
	fmt.Println("gin context", c)
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Println("the error is ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	albums = append(albums, newAlbum)
	fmt.Println("the new album is", albums)
	c.IndentedJSON(http.StatusCreated, albums)
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.Run("localhost:8080")
}
