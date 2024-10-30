package model

import "github.com/tiksup/tiksup-kafka-worker/pkg/movie"

type RandoMovie struct {
	UserID string        `json:"user_id"`
	Movies []movie.Movie `json:"movies"`
}
