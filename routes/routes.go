package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tusher-A/go-rest-api/models"
)



func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum models.Album 

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, currentAlbum := range models.Albums {
		if currentAlbum.ID == id {
			c.IndentedJSON(http.StatusOK, currentAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func InitiateRoutes(router *gin.Engine){
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
}