package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/LeeDDHH/go-crud/data"
	"github.com/LeeDDHH/go-crud/types"
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

// POST
func CreateMovie(c *gin.Context) {
	body := c.Request.Body
	var movie types.Movie
	_ = json.NewDecoder(body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	data.Movies = append(data.Movies, movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

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
