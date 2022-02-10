package controller

import (
	"blog/internal/models"
	"blog/internal/services"
	"blog/internal/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type ArticleController interface {
	FindAll() []models.Article
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	services services.ArticleService
}

var validate *validator.Validate

func New(services services.ArticleService) ArticleController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		services: services,
	}
}

func (c *controller) FindAll() []models.Article {
	return c.services.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var article models.Article
	err := ctx.BindJSON(&article)
	if err != nil {
		return err
	}
	err = validate.Struct(article)
	if err != nil {
		return err
	}
	c.services.Save(article)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	articles := c.services.FindAll()
	data := gin.H{
		"title":    "Article Page",
		"articles": articles,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
