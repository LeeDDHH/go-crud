package router

import (
	"github.com/LeeDDHH/go-crud/api"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	h := r.Group("/")
	{
		h.GET("/movies", api.GetMovies)
		h.GET("/movies/:id", api.GetMovie)
		h.POST("/movies", api.CreateMovie)
		h.PUT("/movies/:id", api.UpdateMovie)
		h.DELETE("/movies/:id", api.DeleteMovie)
	}
}
