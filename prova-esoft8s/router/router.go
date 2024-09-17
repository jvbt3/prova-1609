package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prova-esoft8s/controller"
)

func NewRouter(viagemController *controller.ViagemController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	viagemRouter := router.Group("/viagem")
	viagemRouter.GET("", viagemController.FindAll)
	viagemRouter.GET("/:viagemId", viagemController.FindById)
	viagemRouter.POST("", viagemController.Create)
	viagemRouter.PATCH("/:viagemId", viagemController.Update)
	viagemRouter.DELETE("/:viagemId", viagemController.Delete)

	return service
}
