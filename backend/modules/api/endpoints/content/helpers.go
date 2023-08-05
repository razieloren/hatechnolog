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
		})
}

func QueryPosts(dbConn *gorm.DB) *gorm.DB {
	return dbConn.Model(&models.Content{}).
		Preload("User").Preload("Category").
		Where(&models.Content{
			Type: "posts",
		})
}

func ContentToTeaser(item *models.ContentTeaser) *content.ContentTeaser {
	return &content.ContentTeaser{
		Slug:      item.Slug,
		Title:     item.Title,
		Author:    item.User.Handle,
		Category:  item.Category.Slug,
		Monetized: item.Monetized,
		Upvotes:   uint32(item.Upvotes),
		Published: &timestamppb.Timestamp{
			Seconds: item.CreatedAt.Unix(),
		},
	}
}

func ContentToDetails(item *models.Content) *content.ContentDetails {
	return &content.ContentDetails{
		Teaser: &content.ContentTeaser{
			Slug:      item.Slug,
			Title:     item.Title,
			Author:    item.User.Handle,
			Category:  item.Category.Slug,
			Monetized: item.Monetized,
			Upvotes:   uint32(item.Upvotes),
			Published: &timestamppb.Timestamp{
				Seconds: item.CreatedAt.Unix(),
			},
		},
		Ltr:               item.Ltr,
		CompressedContent: item.CompressedHtml,
		Edited: &timestamppb.Timestamp{
			Seconds: item.UpdatedAt.Unix(),
		},
	}
}
