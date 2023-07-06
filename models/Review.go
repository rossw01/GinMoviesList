package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	AuthorId int
	Author   string
	MovieId  string // Needs to be string, this isn't a mistake 😼
	Rating   float32
	Subject  string
	Body     string
	Date     string
}
