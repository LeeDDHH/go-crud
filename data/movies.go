package data

import (
	"github.com/LeeDDHH/go-crud/types"
)

var Movies []types.Movie

func Init () {
	Movies = append(Movies, types.Movie{ID: "1", Isbn: "438222", Title: "Movie One", Director: &types.Director{Firstname: "John", Lastname: "Doe"}})
	Movies = append(Movies, types.Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &types.Director{Firstname: "Steve", Lastname: "Smith"}})
}
