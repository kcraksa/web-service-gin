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

var albums = []album{
	{ID: "1", Title: "Eight Day's a Week", Artist: "The Beatles", Price: 24.00},
	{ID: "2", Title: "Meraih Mimpi", Artist: "J-Rocks", Price: 15.50},
	{ID: "3", Title: "Balonku Ada Lima", Artist: "Kak Seto", Price: 5.00},
	{ID: "4", Title: "Si Doel Anak Sekolahan", Artist: "Rano Karno", Price: 10.77},
}

func main() {
	router := gin.Default()

	// Router
	router.GET("/albums", getAlbum)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

// getAlbum responds with the list of all albums as JSON
func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbum adds an album from JSON received in the request body.
func postAlbum(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
