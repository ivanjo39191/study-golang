package api

import (
	"net/http"

	"blog/internal/controller"
	"blog/internal/models"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
	loginController   controller.LoginController
	articleController controller.ArticleController
}

func NewArticleAPI(loginController controller.LoginController,
	articleController controller.ArticleController) *ArticleApi {
	return &ArticleApi{
		loginController:   loginController,
		articleController: articleController,
	}
}

// Paths Information

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} models.JWT
// @Failure 401 {object} models.Response
// @Router /auth/token [post]
func (api *ArticleApi) Authenticate(ctx *gin.Context) {
	token := api.loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &models.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, &models.Response{
			Message: "Not Authorized",
		})
	}
}

// GetArticles godoc
// @Security bearerAuth
// @Summary List existing articles
// @Description Get all the existing articles
// @Tags articles,list
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Article
// @Failure 401 {object} models.Response
// @Router /articles [get]
func (api *ArticleApi) GetArticles(ctx *gin.Context) {
	ctx.JSON(200, api.articleController.FindAll())
}

// CreateArticle godoc
// @Security bearerAuth
// @Summary Create new articles
// @Description Create a new article
// @Tags articles,create
// @Accept  json
// @Produce  json
// @Param article body entity.Article true "Create article"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /articles [post]
func (api *ArticleApi) CreateArticle(ctx *gin.Context) {
	err := api.articleController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &models.Response{
			Message: "Success!",
		})
	}
}

// UpdateArticle godoc
// @Security bearerAuth
// @Summary Update articles
// @Description Update a single article
// @Security bearerAuth
// @Tags articles
// @Accept  json
// @Produce  json
// @Param  id path int true "Article ID"
// @Param article body entity.Article true "Update article"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /articles/{id} [put]
func (api *ArticleApi) UpdateArticle(ctx *gin.Context) {
	err := api.articleController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &models.Response{
			Message: "Success!",
		})
	}
}

// DeleteArticle godoc
// @Security bearerAuth
// @Summary Remove articles
// @Description Delete a single article
// @Security bearerAuth
// @Tags articles
// @Accept  json
// @Produce  json
// @Param  id path int true "Article ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /articles/{id} [delete]
func (api *ArticleApi) DeleteArticle(ctx *gin.Context) {
	err := api.articleController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &models.Response{
			Message: "Success!",
		})
	}
}
