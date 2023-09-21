package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/washington-shoji/gin-api/handlers"
)

func NewRouter(bookHandler *handlers.BookHandler) *gin.Engine {
	service := gin.Default()

	service.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "server up and running")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	bookRouter := router.Group("/book")
	bookRouter.GET("", bookHandler.GetAllBooks)
	bookRouter.GET("/:id", bookHandler.GetBookByID)
	bookRouter.POST("", bookHandler.Create)
	bookRouter.PATCH("/:id", bookHandler.Update)
	bookRouter.DELETE("/:id", bookHandler.Delete)

	return service
}
