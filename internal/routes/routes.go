package routes

import (
	"github.com/gin-gonic/gin"
	animal "github.com/tvitcom/animalapi-example/internal/animal"
)

func GetRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("animal/", animal.IndexHandler)
	r.GET("animal/show/:id", animal.ShowHandler)
	r.GET("animal/new/", animal.NewHandler)
	r.POST("animal/", animal.CreateHandler)
	r.GET("animal/edit/:id", animal.EditHandler)
	r.POST("animal/update/:id", animal.UpdateHandler)
	r.GET("animal/delete/:id", animal.DeleteHandler)

	return r
}
