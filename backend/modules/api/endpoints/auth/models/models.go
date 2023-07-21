package models

import (
	"fmt"

	"gorm.io/gorm"
)

func AutoMigrate(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		&Resource{},
		&AccessType{},
		&Permission{},
		&Role{},
		&Plan{},
		&OAuth2Token{},
		&DiscordUser{},
		&GithubUser{},
		&User{},
		&Session{},
	)
}

func CreateDefaultData(dbConn *gorm.DB) error {
	if err := dbConn.Transaction(func(tx *gorm.DB) error {
		karmaResource := Resource{
			Name:        "karma",
			Description: "Karma Leaderboards",
		}
		postResource := Resource{
			Name:        "post",
			Description: "Blog Posts System",
		}
		commentResource := Resource{
			Name:        "comment",
			Description: "Blog Post Comments System",
		}
		memberResource := Resource{
			Name:        "member",
			Description: "Member Page",
		}
		resources := []*Resource{
			&karmaResource,
			&postResource,
			&commentResource,
			&memberResource,
		}
		for _, resource := range resources {
			if err := tx.FirstOrCreate(resource, &Resource{Name: resource.Name}).Error; err != nil {
				return fmt.Errorf("error creating default resources: %w", err)
			}
		}

		readAccessType := AccessType{
			Name:        "read",
			Description: "Read (view) the object",
		}
		createAccessType := AccessType{
			Name:        "create",
			Description: "Create more of this object",
		}
		editAccessType := AccessType{
			Name:        "edit",
			Description: "Edit objects of this type",
		}
		deleteAccessType := AccessType{
			Name:        "delete",
			Description: "Delete objects of this type",
		}
		reactAccessType := AccessType{
			Name:        "react",
			Description: "Be able to react (likes, comments) on this object",
		}
		accessTypes := []*AccessType{
			&readAccessType,
			&createAccessType,
			&editAccessType,
			&deleteAccessType,
			&reactAccessType,
		}
		for _, accessType := range accessTypes {
			if err := tx.FirstOrCreate(accessType, &AccessType{Name: accessType.Name}).Error; err != nil {
				return fmt.Errorf("error creating default access types: %w", err)
			}
		}

		karmaReadPermission := Permission{
			ResourceID:   karmaResource.ID,
			AccessTypeID: readAccessType.ID,
			Description:  "View the karma leaderboards",
		}
		postCreatePermission := Permission{
			ResourceID:   postResource.ID,
			AccessTypeID: createAccessType.ID,
			Description:  "Create a new blog post",
		}
		postEditPermission := Permission{
			ResourceID:   postResource.ID,
			AccessTypeID: editAccessType.ID,
			Description:  "Edit a blog post",
		}
		postDeletePermission := Permission{
			ResourceID:   postResource.ID,
			AccessTypeID: deleteAccessType.ID,
			Description:  "Delete a blog post",
		}
		postReactPermission := Permission{
			ResourceID:   postResource.ID,
			AccessTypeID: reactAccessType.ID,
			Description:  "React & comment a blog post",
		}
		commentReactPermission := Permission{
			ResourceID:   commentResource.ID,
			AccessTypeID: reactAccessType.ID,
			Description:  "React & comment a other comments",
		}
		commentDeletePermission := Permission{
			ResourceID:   commentResource.ID,
			AccessTypeID: deleteAccessType.ID,
			Description:  "Delete blog posts comments",
		}
		commentEditPermission := Permission{
			ResourceID:   commentResource.ID,
			AccessTypeID: editAccessType.ID,
			Description:  "Edit blog posts comments",
		}
		permissions := []*Permission{
			&karmaReadPermission,
			&postCreatePermission,
			&postEditPermission,
			&postDeletePermission,
			&postReactPermission,
			&commentReactPermission,
			&commentDeletePermission,
			&commentEditPermission,
		}
		for _, permission := range permissions {
			if err := tx.FirstOrCreate(permission, &Permission{ResourceID: permission.ResourceID, AccessTypeID: permission.AccessTypeID}).Error; err != nil {
				return fmt.Errorf("error creating default permissions: %w", err)
			}
		}

		memberRole := Role{
			Name:        "member",
			Description: "Community Member",
			Permissions: []Permission{
				karmaReadPermission,
				postReactPermission,
				commentReactPermission,
			},
		}
		editorRole := Role{
			Name:        "editor",
			Description: "Community Editor",
			Permissions: []Permission{
				postCreatePermission,
				postEditPermission,
			},
		}
		adminRole := Role{
			Name:        "admin",
			Description: "Community Admin",
			Permissions: []Permission{
				postDeletePermission,
				commentEditPermission,
				commentDeletePermission,
			},
		}
		roles := []*Role{
			&memberRole,
			&editorRole,
			&adminRole,
		}
		for _, role := range roles {
			if err := tx.FirstOrCreate(role, &Role{Name: role.Name}).Error; err != nil {
				return fmt.Errorf("error creating default roles: %w", err)
			}
		}

		basicPlan := Plan{
			Name:        "basic",
			Description: "Not supporting on any platform",
		}
		supporterPlan := Plan{
			Name:        "supporter",
			Description: "Supporter subscription on Discord",
		}
		vipPlan := Plan{
			Name:        "vip",
			Description: "VIP subscription on Discord",
		}
		plans := []*Plan{
			&basicPlan,
			&supporterPlan,
			&vipPlan,
		}
		for _, plan := range plans {
			if err := tx.FirstOrCreate(plan, &Plan{Name: plan.Name}).Error; err != nil {
				return fmt.Errorf("error creating default plans: %w", err)
			}
		}

		return nil
	}); err != nil {
		return fmt.Errorf("create default data tx: %w", err)
	}
	return nil
}
