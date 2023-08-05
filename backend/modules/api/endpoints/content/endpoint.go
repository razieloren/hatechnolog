package content

import (
	"backend/modules/api/endpoints/content/models"
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/messages/content"
	"backend/x/web"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndpointGetPageRequest(dbConn *gorm.DB, c echo.Context, request *content.GetPageRequest) error {
	var contentItem models.Content
	if err := QueryPages(dbConn).Where(&models.Content{
		Slug: request.Slug,
	}).Take(&contentItem).Error; err != nil {
		c.Logger().Error("Error fetching page: ", err)
		return web.GenerateInternalServerError()
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetPageResponse{
			GetPageResponse: &content.GetPageResponse{
				Details: ContentToDetails(&contentItem),
			},
		},
	})
}

func EndpointGetPostsTeasersRequest(dbConn *gorm.DB, c echo.Context, request *content.GetPostsTeasersRequest) error {
	var teasers []models.ContentTeaser
	if err := QueryPosts(dbConn).Find(&teasers).Error; err != nil {
		c.Logger().Error("Error finding post teasers: ", err)
		return web.GenerateInternalServerError()
	}
	var contentMessage content.GetPostsTeasersResponse
	for _, teaser := range teasers {
		contentMessage.Teasers = append(contentMessage.Teasers, ContentToTeaser(&teaser))
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetPostsTeasersResponse{
			GetPostsTeasersResponse: &contentMessage,
		},
	})
}

func EndpointGetPostRequest(dbConn *gorm.DB, c echo.Context, request *content.GetPostRequest) error {
	var contentItem models.Content
	if err := QueryPosts(dbConn).Where(&models.Content{
		Slug: request.Slug,
	}).Take(&contentItem).Error; err != nil {
		c.Logger().Error("Error fetching post: ", err)
		return web.GenerateInternalServerError()
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetPostResponse{
			GetPostResponse: &content.GetPostResponse{
				Details: ContentToDetails(&contentItem),
			},
		},
	})
}

func EndpointGetCategoriesTeasersRequest(dbConn *gorm.DB, c echo.Context, request *content.GetCategoriesTeasersRequest) error {
	var teasers []models.CategoryTeaser
	if err := dbConn.Model(&models.Category{}).Find(&teasers).Error; err != nil {
		c.Logger().Error("Error finding category teasers: ", err)
		return web.GenerateInternalServerError()
	}
	var categoryMessage content.GetCategoriesTeasersResponse
	for _, teaser := range teasers {
		categoryMessage.Teasers = append(categoryMessage.Teasers, &content.CategoryTeaser{
			Slug:        teaser.Slug,
			Name:        teaser.Name,
			Description: teaser.Description,
		})
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetCategoriesTeasersResponse{
			GetCategoriesTeasersResponse: &categoryMessage,
		},
	})
}

func EndpointGetCategoryRequest(dbConn *gorm.DB, c echo.Context, request *content.GetCategoryRequest) error {
	var categoryItem models.Category
	if err := dbConn.Where(&models.Category{
		Slug: request.Slug,
	}).Take(&categoryItem).Error; err != nil {
		c.Logger().Error("Error fetching category: ", err)
		return web.GenerateInternalServerError()
	}
	var teasers []models.ContentTeaser
	if err := dbConn.Model(&models.Content{}).
		Preload("User").Preload("Category").Where(&models.Content{
		CategoryID: categoryItem.ID,
	}).Find(&teasers).Error; err != nil {
		c.Logger().Error("Error finding post teasers: ", err)
		return web.GenerateInternalServerError()
	}
	var teasersList []*content.ContentTeaser
	for _, teaser := range teasers {
		teasersList = append(teasersList, ContentToTeaser(&teaser))
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetCategoryResponse{
			GetCategoryResponse: &content.GetCategoryResponse{
				Details: &content.CategoryDetails{
					Teaser: &content.CategoryTeaser{
						Slug:        categoryItem.Slug,
						Name:        categoryItem.Name,
						Description: categoryItem.Description,
					},
					Contents: teasersList,
				},
			},
		},
	})
}
