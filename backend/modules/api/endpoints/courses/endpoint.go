package courses

import (
	"backend/modules/api/endpoints/courses/models"
	"backend/modules/api/endpoints/messages"
	"backend/modules/api/endpoints/messages/courses"
	"backend/x/web"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EndpointGetCoursesTeasers(dbConn *gorm.DB, c echo.Context, request *courses.GetCoursesTeasersRequest) error {
	var teasers []models.CourseTeaser
	if err := dbConn.Model(&models.Course{}).Find(&teasers).Error; err != nil {
		c.Logger().Error("Error getting courses teasers: ", err)
		return web.GenerateInternalServerError()
	}
	var coursesMessage courses.GetCoursesTeasersResponse
	for _, teaser := range teasers {
		coursesMessage.Teasers = append(coursesMessage.Teasers, &courses.CourseTeaser{
			Slug:        teaser.Slug,
			Title:       teaser.Title,
			MainVideoId: teaser.MainVideoID,
			Description: teaser.Description,
		})
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetCoursesTeasersResponse{
			GetCoursesTeasersResponse: &coursesMessage,
		},
	})
}

func EndpointGetCourse(dbConn *gorm.DB, c echo.Context, request *courses.GetCourseRequest) error {
	var course models.Course
	if err := dbConn.Where(&models.Course{
		Slug: request.Slug,
	}).Take(&course).Error; err != nil {
		c.Logger().Error("Error getting course: ", err)
		return web.GenerateInternalServerError()
	}
	return web.GenerateResponse(c, &messages.Wrapper{
		Message: &messages.Wrapper_GetCourseResponse{
			GetCourseResponse: &courses.GetCourseResponse{
				Details: &courses.CourseDetails{
					Teaser: &courses.CourseTeaser{
						Slug:        course.Slug,
						Title:       course.Title,
						MainVideoId: course.MainVideoID,
						Description: course.Description,
					},
					PlaylistId: course.PlaylistID,
				},
			},
		},
	})
}
