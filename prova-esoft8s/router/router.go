package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"prova-esoft8s/controller"
	"time"
)

func NewRouter(viagemController *controller.ViagemController) *gin.Engine {
	service := gin.Default()

	service.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
