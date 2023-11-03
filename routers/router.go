package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/washington-shoji/gin-api/handlers"
	"github.com/washington-shoji/gin-api/middleware"
)

func NewRouter(
	bookHandler *handlers.BookHandler,
	userAccountHandler *handlers.UserAccountHandler,
	loginHandler *handlers.LoginHandler,
	tableTopGameHandler *handlers.TableTopGameHandler,
	dynamicData *handlers.DynamicDataHandler,
	eventHandler *handlers.EventHandler,
	expHandler *handlers.ExpHandler,

) *gin.Engine {
	service := gin.Default()

	service.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	service.LoadHTMLGlob("templates/**/*")
	service.Static("/static", "./static")

	service.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "server up and running")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")

	loginRouter := router.Group("/auth")
	loginRouter.GET("", loginHandler.LoginRenderForm)
	loginRouter.POST("", loginHandler.Login)
	loginRouter.POST("/login", loginHandler.LoginRenderAuth)
	loginRouter.POST("/logout", loginHandler.LogOutAuth)

	bookRouter := router.Group("/book")
	//bookRouter.Use(middleware.JwtAuthMiddleware())
	bookRouter.GET("", bookHandler.GetAllBooks)
	bookRouter.GET("/:id", bookHandler.GetBookByID)
	bookRouter.POST("", bookHandler.Create)
	bookRouter.PATCH("/:id", bookHandler.Update)
	bookRouter.DELETE("/:id", bookHandler.Delete)

	userAccRouter := router.Group("/account")
	userAccRouter.GET("", userAccountHandler.GetAllUserAccounts)
	userAccRouter.GET("/:id", userAccountHandler.GetUserAccountByID)
	userAccRouter.POST("", userAccountHandler.Create)
	userAccRouter.PATCH("/:id", userAccountHandler.Update)
	userAccRouter.DELETE("/:id", userAccountHandler.Delete)

	tableTopGameRouter := router.Group("/table-top")
	tableTopGameRouter.GET("", tableTopGameHandler.FindAll)
	tableTopGameRouter.GET("/:id", tableTopGameHandler.FindByID)
	tableTopGameRouter.POST("", tableTopGameHandler.Create)
	tableTopGameRouter.PATCH("/:id", tableTopGameHandler.Update)
	tableTopGameRouter.DELETE("/:id", tableTopGameHandler.Delete)

	eventRouter := router.Group("/event")
	eventRouter.GET("", eventHandler.GetAllEvents)
	eventRouter.GET("/:id", eventHandler.GetEventByID)
	eventRouter.POST("", eventHandler.Create)
	eventRouter.PATCH("/:id", eventHandler.Update)
	eventRouter.DELETE("/:id", eventHandler.Delete)

	dynamicDataRouter := router.Group("/dynamic")
	dynamicDataRouter.GET("", dynamicData.FindAll)
	dynamicDataRouter.GET("/:id", dynamicData.FindByID)
	dynamicDataRouter.POST("", dynamicData.Create)
	dynamicDataRouter.PATCH("/:id", dynamicData.Update)
	dynamicDataRouter.DELETE("/:id", dynamicData.Delete)

	expRouter := router.Group("/exp")
	expRouter.POST("", expHandler.ExpCreate)
	expRouter.GET("", expHandler.ExpGetAll)

	htmlRouter := router.Group("/exp-html")
	htmlRouter.Use(middleware.JwtAuthMiddlewareCookie())

	htmlRouter.GET("/book", bookHandler.RenderBookForm)
	htmlRouter.POST("/book", bookHandler.CreateBookForm)

	htmlRouter.GET("/base", bookHandler.RenderPartials)
	htmlRouter.GET("/content", bookHandler.RenderContent)
	htmlRouter.GET("/home", bookHandler.RenderHomepage)
	htmlRouter.GET("/book-update/:id", bookHandler.RenderUpdateBookForm)
	htmlRouter.PUT("/book-update/:id", bookHandler.RenderUpdateBook)
	htmlRouter.DELETE("/home/:id", bookHandler.RenderDeleteBook)

	return service
}
