package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Slug        string `gorm:"NOT NULL;UNIQUE"`
	Title       string `gorm:"NOT NULL;"`
	Description string `gorm:"NOT NULL;"`
	PlaylistID  string `gorm:"NOT NULL;"`
	MainVideoID string `gorm:"NOT NULL;"`
}

type CourseTeaser struct {
	Slug        string
	Title       string
	Description string
	PlaylistID  string
	MainVideoID string
}

func (course *Course) TableName() string {
	return "api.course"
}
