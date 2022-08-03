package model

import "gorm.io/gorm"

type Alias struct {
	*gorm.Model
	Alias string
}
