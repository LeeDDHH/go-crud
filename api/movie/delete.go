package movie

import (
	"net/http"

	"github.com/LeeDDHH/go-crud/data"
	"github.com/gin-gonic/gin"
)

// DELETE
func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	for index, item := range data.Movies {
		if item.ID == id {
			data.Movies = append(data.Movies[:index], data.Movies[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"data": data.Movies})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "there's no id of movies"})
}
