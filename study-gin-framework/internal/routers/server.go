package routers

import (
	"blog/internal/controller"
	"blog/internal/services"
	"io"
	"net/http"
	"os"

	"blog/internal/middlewares"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	articleService    services.ArticleService      = services.New()
	articleController controller.ArticleController = controller.New(articleService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func ServerCore() {

	setupLogOutput()
	router := gin.New()
	router.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		gindump.Dump(),
	)
	router.Static("/css", "../../internal/templates/css")

	router.LoadHTMLGlob("../../internal/templates/*.html")

	apiRoutes := router.Group("/api")
	{
		apiRoutes.POST("/posts", func(c *gin.Context) {
			err := articleController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Article Input is valid"})
			}
		})
		apiRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "OK!!"})
		})
	}

	viewRoutes := router.Group("/view")
	{
		viewRoutes.GET("/posts", func(c *gin.Context) {
			c.JSON(200, articleController.FindAll())
		})
		viewRoutes.GET("/articles", articleController.ShowAll)
	}

	router.Run(":7528")
	err := router.Run()
	if err != nil {
		panic(err)
	}

}
