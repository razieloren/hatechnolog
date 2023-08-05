package models

import (
	"backend/modules/api/endpoints/auth/models"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Slug        string `gorm:"NOT NULL; UNIQUE"`
	Name        string `gorm:"NOT NULL"`
	Description string `gorm:"NOT NULL"`
}

type CategoryTeaser struct {
	Slug        string
	Name        string
	Description string
	CreatedAt   time.Time
}

func (category *Category) TableName() string {
	return "api.category"
}

type Content struct {
	gorm.Model
	Slug           string `gorm:"NOT NULL; UNIQUE"`
	Type           string `gorm:"NOT NULL;"`
	Title          string `gorm:"NOT NULL"`
	UserID         uint   `gorm:"NOT NULL"`
	User           models.User
	Ltr            bool   `gorm:"NOT NULL"`
	CompressedHtml []byte `gorm:"NOT NULL"`
	Upvotes        uint   `gorm:"NOT NULL"`
	CategoryID     uint   `gorm:"NOT NULL"`
	Category       Category
	Monetized      bool   `gorm:"NOT NULL"`
	Hash           []byte `gorm:"NOT NULL"`
}

type ContentTeaser struct {
	Slug       string
	Title      string
	UserID     uint
	User       models.User
	Upvotes    uint
	CategoryID uint
	Category   Category
	Monetized  bool
	CreatedAt  time.Time
}

func (content *Content) TableName() string {
	return "api.content"
}
