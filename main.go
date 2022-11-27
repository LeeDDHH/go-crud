package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

func deleteMovie(c *gin.Context) {
	id := c.Param("id")
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"data": movies})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "there's no id of movies"})
}

func getMovie(c *gin.Context) {
	id := c.Param("id")
	for _, item := range movies {
		if item.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": item})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "there's no id of movies"})
}

func createMovie(c *gin.Context) {
	body := c.Request.Body
	var movie Movie
	_ = json.NewDecoder(body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func updateMovie(c *gin.Context) {
	id := c.Param("id")
	body := c.Request.Body
	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(body).Decode(&movie)
			movie.ID = id
			movies = append(movies, movie)
			c.JSON(http.StatusOK, gin.H{"data": movies})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": "no update"})
}

func main() {
	r := gin.Default()

	movies = append(movies, Movie{ID: "1", Isbn: "438222", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	h := r.Group("/")
	{
		h.GET("/movies", getMovies)
		h.GET("/movies/:id", getMovie)
		h.POST("/movies", createMovie)
		h.PUT("/movies/:id", updateMovie)
		h.DELETE("/movies/:id", deleteMovie)
	}

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
