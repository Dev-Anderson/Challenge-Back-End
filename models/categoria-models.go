package models

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model
	Titulo string `json:"titulo"`
	Cor    string `json:"cor"`
}
