package movie

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/LeeDDHH/go-crud/data"
	"github.com/LeeDDHH/go-crud/types"
	"github.com/gin-gonic/gin"
)

// POST
func CreateMovie(c *gin.Context) {
	body := c.Request.Body
	var movie types.Movie
	_ = json.NewDecoder(body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	data.Movies = append(data.Movies, movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}
