package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.

// func Index(c *gin.Context) {
// 	app.Save("123", "b")
// 	c.IndentedJSON(http.StatusOK, albums)
// }

//	func Health(c *gin.Context) {
//		c.IndentedJSON(http.StatusOK, albumss)
//	}
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func AddUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func UpdateUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func DeleteUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
