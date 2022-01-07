package movie

import "github.com/emregocer/golang_project/internal/category"

type MovieResource struct {
	Id         int                         `json:"id"`
	Name       string                      `json:"name"`
	Plot       string                      `json:"plot"`
	Categories []category.CategoryResource `json:"categories,omitempty"`
}
