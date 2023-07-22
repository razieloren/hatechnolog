package models

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const (
	PermissionStringSeparator = ":"
)

type Resource struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"UNIQUE"`
	Description string
}

func (resource *Resource) TableName() string {
	return "api.resource"
}

type AccessType struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"UNIQUE"`
	Description string
}

func (accessType *AccessType) TableName() string {
	return "api.access_type"
}

type Permission struct {
	ID           uint `gorm:"primarykey"`
	ResourceID   uint `gorm:"NOT NULL"`
	Resource     Resource
	AccessTypeID uint `gorm:"NOT NULL"`
	AccessType   AccessType
	Description  string `gorm:"NOT NULL"`
}

func (permission *Permission) TableName() string {
	return "api.permission"
}

func (permission *Permission) FromPermissionString(dbConn *gorm.DB, permissionString string) error {
	if !strings.Contains(permissionString, PermissionStringSeparator) {
		return fmt.Errorf("no permission separator")
	}
	parts := strings.Split(permissionString, PermissionStringSeparator)
	if len(parts) != 2 {
		return fmt.Errorf("bad permissions streucture")
	}
	resourceName, accessTypeName := parts[0], parts[1]
	var resource Resource
	if err := dbConn.Where(
		&Resource{Name: resourceName}).Take(&resource).Error; err != nil {
		return fmt.Errorf("take resource: %w", err)
	}
	var accessType AccessType
	if err := dbConn.Where(
		&AccessType{Name: accessTypeName}).Take(&accessType).Error; err != nil {
		return fmt.Errorf("take access type: %w", err)
	}
	if err := dbConn.Where(
		&Permission{ResourceID: resource.ID, AccessTypeID: accessType.ID}).Take(permission).Error; err != nil {
		return fmt.Errorf("take permission: %w", err)
	}
	return nil
}
