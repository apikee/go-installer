package model

type Record struct {
	ID    uint `gorm:"primaryKey"`
	Alias string
	Path  string
}
