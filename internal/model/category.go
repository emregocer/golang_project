package model

type Category struct {
	Base
	SoftDelete
	Name        string `json:"name"`
	Description string `json:"description"`
}
