package model

type MovieCategory struct {
	Base
	MovieId    int `json:"movie_id" db:"movie_id"`
	CategoryId int `json:"category_id" db:"category_id"`
}
