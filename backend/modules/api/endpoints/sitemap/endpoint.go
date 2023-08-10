package sitemap

import (
	"backend/modules/api/config"
	"backend/modules/api/endpoints/content"
	category_models "backend/modules/api/endpoints/content/models"
	content_models "backend/modules/api/endpoints/content/models"
	course_models "backend/modules/api/endpoints/courses/models"
	"backend/x/web"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pahanini/go-sitemap-generator"
	"gorm.io/gorm"
)

const (
	sitemapDir = "/tmp/sitemaps/hatechnolog"
)

func hostPath(path string) string {
	return fmt.Sprintf("%s%s", config.Globals.Auth.RedirectHost, path)
}

func EndpointSitemap(dbConn *gorm.DB, c echo.Context) error {
	sitemapPath := fmt.Sprintf("%s/sitemap.xml", sitemapDir)
	os.Remove(sitemapPath)
	sm := sitemap.New(sitemap.Options{
		Dir:     sitemapDir,
		BaseURL: config.Globals.Auth.RedirectHost,
	})
	if err := sm.Open(); err != nil {
		c.Logger().Error("Error opening sitemap: ", err)
		return web.GenerateInternalServerError()
	}
	sm.Add(sitemap.URL{Loc: hostPath("/")})
	sm.Add(sitemap.URL{Loc: hostPath("/blog")})
	sm.Add(sitemap.URL{Loc: hostPath("/courses")})
	sm.Add(sitemap.URL{Loc: hostPath("/category")})

	var posts []content_models.ContentTeaser
	if err := content.QueryPosts(dbConn).Find(&posts).Error; err != nil {
		c.Logger().Error("Error fetching posts: ", err)
		return web.GenerateInternalServerError()
	}
	for _, post := range posts {
		sm.Add(sitemap.URL{
			Loc:     hostPath(fmt.Sprintf("/blog/%s", post.Slug)),
			LastMod: post.UpdatedAt.Format("2006-01-02"),
		})
	}
	var pages []content_models.ContentTeaser
	if err := content.QueryPages(dbConn).Find(&pages).Error; err != nil {
		c.Logger().Error("Error fetching pages: ", err)
		return web.GenerateInternalServerError()
	}
	for _, page := range pages {
		sm.Add(sitemap.URL{
			Loc:     hostPath(fmt.Sprintf("/%s", page.Slug)),
			LastMod: page.UpdatedAt.Format("2006-01-02"),
		})
	}
	var courses []course_models.CourseTeaser
	if err := dbConn.Model(&course_models.Course{}).Find(&courses).Error; err != nil {
		c.Logger().Error("Error fetching courses: ", err)
		return web.GenerateInternalServerError()
	}
	for _, course := range courses {
		sm.Add(sitemap.URL{
			Loc: hostPath(fmt.Sprintf("/courses/%s", course.Slug)),
		})
	}
	var categories []category_models.CategoryTeaser
	if err := dbConn.Model(&category_models.Category{}).Find(&categories).Error; err != nil {
		c.Logger().Error("Error fetching categories: ", err)
		return web.GenerateInternalServerError()
	}
	for _, category := range categories {
		sm.Add(sitemap.URL{
			Loc: hostPath(fmt.Sprintf("/category/%s", category.Slug)),
		})
	}
	if err := sm.Close(); err != nil {
		c.Logger().Error("Error saving sitemap: ", err)
		return web.GenerateInternalServerError()
	}
	content, err := ioutil.ReadFile(sitemapPath)
	if err != nil {
		c.Logger().Error("Error reading sitemap file: ", err)
		return web.GenerateInternalServerError()
	}
	content = bytes.Replace(content, []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?><urlset"), []byte("<urlset"), -1)
	return c.XMLBlob(http.StatusOK, content)
}
