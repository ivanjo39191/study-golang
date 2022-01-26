package controller

import (
	"blog/internal/models"
	"blog/internal/services"

	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	FindAll() []models.Article
	Save(ctx *gin.Context) models.Article
}

type controller struct {
	services services.ArticleService
}

func New(services services.ArticleService) ArticleController {
	return &controller{
		services: services,
	}
}

func (c *controller) FindAll() []models.Article {
	return c.services.FindAll()
}

func (c *controller) Save(ctx *gin.Context) models.Article {
	var article models.Article
	ctx.BindJSON(&article)
	c.services.Save(article)
	return article
}
