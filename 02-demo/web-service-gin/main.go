package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var albums = []student{
	{"张三", 23},
	{"李四", 24},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
