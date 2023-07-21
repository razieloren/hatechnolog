package models

type Role struct {
	ID          uint         `gorm:"primarykey"`
	Name        string       `gorm:"UNIQUE;NOT NULL"`
	Permissions []Permission `gorm:"many2many:api.role_permission"`
	Description string       `gorm:"NOT NULL"`
}

func (role *Role) TableName() string {
	return "api.role"
}
