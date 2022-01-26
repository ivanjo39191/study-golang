package routers

import (
	"blog/internal/controller"
	"blog/internal/services"
	"io"
	"os"

	"blog/internal/middlewares"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	articleService    services.ArticleService      = services.New()
	ArticleController controller.ArticleController = controller.New(articleService)
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
	router.GET("/posts", func(c *gin.Context) {
		c.JSON(200, ArticleController.FindAll())
	})
	router.POST("/posts", func(c *gin.Context) {
		c.JSON(200, ArticleController.Save(c))
	})
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK!!"})
	})

	router.Run(":7528")
	err := router.Run()
	if err != nil {
		panic(err)
	}

}
