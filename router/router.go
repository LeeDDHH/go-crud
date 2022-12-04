package router

import (
	"github.com/LeeDDHH/go-crud/api/movie"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	h := r.Group("/")
	{
		h.GET("/movies", movie.GetMovies)
		h.GET("/movies/:id", movie.GetMovie)
		h.POST("/movies", movie.CreateMovie)
		h.PUT("/movies/:id", movie.UpdateMovie)
		h.DELETE("/movies/:id", movie.DeleteMovie)
	}
}
