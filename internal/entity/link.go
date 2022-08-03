package entity

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	URL  string `db:"url"`
	Slug string `db:"slug" gorm:"unique"`
}
