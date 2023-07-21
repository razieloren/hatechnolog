package models

type Plan struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"UNIQUE;NOT NULL"`
	Description string `gorm:"NOT NULL"`
}

func (plan *Plan) TableName() string {
	return "api.plan"
}
