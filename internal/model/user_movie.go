package model

import "github.com/lib/pq"

type UserMovie struct {
	Base
	UserId       int         `json:"user_id" db:"user_id"`
	MovieId      int         `json:"movie_id" db:"movie_id"`
	FavouritedAt pq.NullTime `json:"favourited_at" db:"favourited_at"`
}
