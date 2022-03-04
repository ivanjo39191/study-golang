package services

import (
	"blog/internal/models"
	"blog/internal/repository"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	TITLE       = "Article Title2"
	DESCRIPTION = "Article Description"
	URL         = "https://youtu.be/JgW-i2QjgHQ"
	FIRSTNAME   = "Ivan3"
	LASTNAME    = "Kao"
	EMAIL       = "ivankao3@mail.com"
)

var testArticle models.Article = models.Article{
	Title:       TITLE,
	Description: DESCRIPTION,
	URL:         URL,
	Author: models.Person{
		FirstName: FIRSTNAME,
		LastName:  LASTNAME,
		Email:     EMAIL,
	},
}

var _ = Describe("Article Service", func() {

	var (
		articleRepository repository.ArticleRepository
		articleService    ArticleService
	)

	BeforeSuite(func() {
		articleRepository = repository.NewArticleRepository()
		articleService = New(articleRepository)
	})

	Describe("Fetching all existing articles", func() {

		Context("If there is a article in the database", func() {

			BeforeEach(func() {
				articleService.Save(testArticle)
			})

			It("should return at least one element", func() {
				articleList := articleService.FindAll()

				Ω(articleList).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				firstArticle := articleService.FindAll()[0]

				Ω(firstArticle.Title).Should(Equal(TITLE))
				Ω(firstArticle.Description).Should(Equal(DESCRIPTION))
				Ω(firstArticle.URL).Should(Equal(URL))
				Ω(firstArticle.Author.FirstName).Should(Equal(FIRSTNAME))
				Ω(firstArticle.Author.LastName).Should(Equal(LASTNAME))
				Ω(firstArticle.Author.Email).Should(Equal(EMAIL))
			})

			AfterEach(func() {
				article := articleService.FindAll()[0]
				articleService.Delete(article)
			})

		})

		Context("If there are no articles in the database", func() {

			It("should return an empty list", func() {
				articles := articleService.FindAll()

				Ω(articles).Should(BeEmpty())
			})

		})
	})

	AfterSuite(func() {
		articleRepository.CloseDB()
	})
})
