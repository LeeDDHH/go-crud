package movie

import (
	"net/http"

	"github.com/LeeDDHH/go-crud/data"
	"github.com/gin-gonic/gin"
)

// GET
func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": data.Movies})
}

func GetMovie(c *gin.Context) {
	id := c.Param("id")
	for _, item := range data.Movies {
		if item.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": item})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "there's no id of movies"})
}
