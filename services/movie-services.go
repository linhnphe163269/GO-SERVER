package services

import Entity "server/entity"

type Movie = Entity.Movie
type MovieService interface {
	GetAllMovies() []Movie
	GetMovieByID(id int64) Movie
}
type movieService struct {
	movies []Movie
}

func New() MovieService {
	return &movieService{}
}
func (s *movieService) GetAllMovies() []Movie {
	return s.movies
}
func (s *movieService) GetMovieByID(id int64) Movie {
	for _, movie := range s.movies {
		if movie.ID == id {
			return movie
		}
	}
	return Movie{}
}
