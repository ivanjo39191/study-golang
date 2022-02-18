package repository

import (
	"blog/internal/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ArticleRepository interface {
	Save(article models.Article)
	Update(article models.Article)
	Delete(article models.Article)
	FindAll() []models.Article
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewArticleRepository() ArticleRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&models.Article{}, &models.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(article models.Article) {
	db.connection.Create(&article)
}

func (db *database) Update(article models.Article) {
	db.connection.Save(&article)
}

func (db *database) Delete(article models.Article) {
	db.connection.Delete(&article)
}

func (db *database) FindAll() []models.Article {
	var articles []models.Article
	db.connection.Set("gorm:auto_preload", true).Find(&articles)
	return articles
}
