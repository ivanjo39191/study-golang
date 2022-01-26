package services

import (
	"blog/internal/models"
)

type ArticleService interface {
	Save(models.Article) models.Article
	FindAll() []models.Article
}

type articleService struct {
	articles []models.Article
}

func New() ArticleService {
	return &articleService{}
}

func (service *articleService) Save(article models.Article) models.Article {
	service.articles = append(service.articles, article)
	return article
}

func (service *articleService) FindAll() []models.Article {
	return service.articles
}
