package routers

import (
	"blog/internal/api"
	"blog/internal/controller"
	"blog/internal/repository"
	"blog/internal/services"
	"io"
	"os"

	"blog/internal/middlewares"

	"blog/internal/docs" // Swagger generated files

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	articleRepository repository.ArticleRepository = repository.NewArticleRepository()
	articleService    services.ArticleService      = services.New(articleRepository)
	loginService      services.LoginService        = services.NewLoginService()
	jwtService        services.JWTService          = services.NewJWTService()
	articleController controller.ArticleController = controller.New(articleService)
	loginController   controller.LoginController   = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func ServerCore() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Ivankao Blog - Article API"
	docs.SwaggerInfo.Description = "Ivankao Blog - Blog Article API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:7528"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	defer articleRepository.CloseDB()

	setupLogOutput()
	router := gin.New()
	router.Use(
		gin.Recovery(),
		middlewares.Logger(),
		gindump.Dump(),
	)
	router.Static("/css", "../../internal/templates/css")

	router.LoadHTMLGlob("../../internal/templates/*.html")

	articleAPI := api.NewArticleAPI(loginController, articleController)
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", articleAPI.Authenticate)
		}

		articles := apiRoutes.Group("/articles", middlewares.AuthorizeJWT())
		{
			articles.GET("", articleAPI.GetArticles)
			articles.POST("", articleAPI.CreateArticle)
			articles.PUT(":id", articleAPI.UpdateArticle)
			articles.DELETE(":id", articleAPI.DeleteArticle)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Login
	// router.POST("/login", func(ctx *gin.Context) {
	// 	token := loginController.Login(ctx)
	// 	if token != "" {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"token": token,
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusUnauthorized, nil)
	// 	}
	// })
	// // API
	// apiRoutes := router.Group("/api", middlewares.AuthorizeJWT())
	// {
	// 	apiRoutes.POST("/posts", func(c *gin.Context) {
	// 		err := articleController.Save(c)
	// 		if err != nil {
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			c.JSON(http.StatusOK, gin.H{"message": "Article Input is valid"})
	// 		}
	// 	})
	// 	apiRoutes.GET("/test", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{"message": "OK!!"})
	// 	})

	// 	apiRoutes.GET("/posts", func(ctx *gin.Context) {
	// 		ctx.JSON(200, articleController.FindAll())
	// 	})

	// 	apiRoutes.PUT("/posts/:id", func(ctx *gin.Context) {
	// 		err := articleController.Update(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
	// 		}

	// 	})

	// 	apiRoutes.DELETE("/posts/:id", func(ctx *gin.Context) {
	// 		err := articleController.Delete(ctx)
	// 		if err != nil {
	// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		} else {
	// 			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
	// 		}

	// 	})

	// }

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
