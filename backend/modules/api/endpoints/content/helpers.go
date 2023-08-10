package content

import (
	"backend/modules/api/endpoints/content/models"
	"backend/modules/api/endpoints/messages/content"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func QueryPages(dbConn *gorm.DB) *gorm.DB {
	return dbConn.Model(&models.Content{}).
		Preload("User").Preload("Category").
		Where(&models.Content{
			Type: "pages",
		}).Order("created_at DESC")
}

func QueryPosts(dbConn *gorm.DB) *gorm.DB {
	return dbConn.Model(&models.Content{}).
		Preload("User").Preload("Category").
		Where(&models.Content{
			Type: "posts",
		}).Order("created_at DESC")
}

func ContentToTeaser(item *models.ContentTeaser) *content.ContentTeaser {
	return &content.ContentTeaser{
		Slug:        item.Slug,
		Title:       item.Title,
		Description: item.Description,
		Author:      item.User.Handle,
		Category:    item.Category.Slug,
		Type:        item.Type,
		Monetized:   item.Monetized,
		Upvotes:     uint32(item.Upvotes),
		Published: &timestamppb.Timestamp{
			Seconds: item.CreatedAt.Unix(),
		},
		Edited: &timestamppb.Timestamp{
			Seconds: item.UpdatedAt.Unix(),
		},
	}
}

func ContentToDetails(item *models.Content) *content.ContentDetails {
	return &content.ContentDetails{
		Teaser: &content.ContentTeaser{
			Slug:        item.Slug,
			Title:       item.Title,
			Description: item.Description,
			Author:      item.User.Handle,
			Category:    item.Category.Slug,
			Monetized:   item.Monetized,
			Upvotes:     uint32(item.Upvotes),
			Published: &timestamppb.Timestamp{
				Seconds: item.CreatedAt.Unix(),
			},
			Edited: &timestamppb.Timestamp{
				Seconds: item.UpdatedAt.Unix(),
			},
		},
		Ltr:               item.Ltr,
		CompressedContent: item.CompressedHtml,
	}
}
