package movie

import (
	"encoding/json"
	"net/http"

	"github.com/LeeDDHH/go-crud/data"
	"github.com/LeeDDHH/go-crud/types"
	"github.com/gin-gonic/gin"
)

// PUT
func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	body := c.Request.Body
	for index, item := range data.Movies {
		if item.ID == id {
			data.Movies = append(data.Movies[:index], data.Movies[index+1:]...)
			var movie types.Movie
			_ = json.NewDecoder(body).Decode(&movie)
			movie.ID = id
			data.Movies = append(data.Movies, movie)
			c.JSON(http.StatusOK, gin.H{"data": data.Movies})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "no update"})
}
