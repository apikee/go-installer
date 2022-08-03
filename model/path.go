package model

import "gorm.io/gorm"

type Path struct {
	*gorm.Model
	Alias uint
	Path  string
}
