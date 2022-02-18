package services

import (
	"blog/internal/models"
	"blog/internal/repository"
)

type ArticleService interface {
	Save(models.Article) error
	Update(models.Article) error
	Delete(models.Article) error
	FindAll() []models.Article
}

type articleService struct {
	repository repository.ArticleRepository
}

func New(articleRepository repository.ArticleRepository) ArticleService {
	return &articleService{
		repository: articleRepository,
	}
}

func (services *articleService) Save(article models.Article) error {
	services.repository.Save(article)
	return nil
}

func (services *articleService) Update(article models.Article) error {
	services.repository.Update(article)
	return nil
}

func (services *articleService) Delete(article models.Article) error {
	services.repository.Delete(article)
	return nil
}

func (services *articleService) FindAll() []models.Article {
	return services.repository.FindAll()
}
