package controller

import (
	Entity "server/entity"
	"server/services"
)

type Movie = Entity.Movie
type MovieController interface {
	GetAllMovies() []Movie
	GetMovieByID(id int64) Movie
}
type controller struct {
	service services.MovieService
}

func New(service services.MovieService) MovieController {
	return controller{
		service: service,
	}
}
func (c *controller) GetAllMovies() []Movie {
	return c.service.GetAllMovies()
}
func (c *controller) GetMovieByID(id int64) Movie {
	return c.service.GetMovieByID(id)
}
