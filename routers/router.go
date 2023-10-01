package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/washington-shoji/gin-api/handlers"
)

func NewRouter(
	bookHandler *handlers.BookHandler,
	userAccountHandler *handlers.UserAccountHandler,
	loginHandler *handlers.LoginHandler,
	tableTopGameHandler *handlers.TableTopGameHandler,
	expHandler *handlers.ExpHandler,
) *gin.Engine {
	service := gin.Default()

	service.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	service.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "server up and running")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")

	loginRouter := router.Group("/login")
	loginRouter.POST("", loginHandler.Login)

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

	expRouter := router.Group("/exp")
	expRouter.POST("", expHandler.Exp)

	return service
}
