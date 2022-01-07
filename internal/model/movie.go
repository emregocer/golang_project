package model

type Movie struct {
	Base
	SoftDelete
	Name       string     `json:"name"`
	Plot       string     `json:"plot"`
	Categories []Category `json:"categories"`
}
