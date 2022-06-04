package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

var albums = []album{
	{ID: "1", Title: "Go", Artist: "Google", Price: 1.99},
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var routes = []route{
	{Method: "GET", Path: "/"},
	{Method: "GET", Path: "/albums"},
	{Method: "GET", Path: "/albums/:id"},
	{Method: "GET", Path: "/ping"},
	{Method: "POST", Path: "/albums"},
}

func main() {
	// GET
	router := gin.Default()
	router.GET("/", getRoutes)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/ping", getPong)

	// POST
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

// GET

func getRoutes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"available routes": routes})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
}

func getPong(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// POST

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
