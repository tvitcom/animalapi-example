package routes

import (
	"github.com/gin-gonic/gin"
	animal "github.com/tvitcom/animalapi-example/internal/animal"
)

func GetRoutes(r *gin.Engine) *gin.Engine {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": "API CRUD of manage the animal records - v.0.0.1",
			"data": "{POST|GET} /animal/{show,update,delete,new}([/:id]|[?page=0-9])",
			"ok": true,
		})
	})

	// r.GET("animal/", animal.IndexHandler)
	// r.GET("animal/show/:id", animal.ShowHandler)
	// r.GET("animal/new/", animal.NewHandler)
	// r.POST("animal/", animal.CreateHandler)
	// r.GET("animal/edit/:id", animal.EditHandler)
	// r.POST("animal/update/:id", animal.UpdateHandler)
	// r.GET("animal/delete/:id", animal.DeleteHandler)

	// Rest API only:
	r.GET("animal/", animal.ApiIndexHandler) //info about api
	r.GET("animal/new/", animal.ApiNewHandler) // return json-template
	r.GET("animal/show/:id", animal.ApiShowHandler) // retunt json about Animal
	r.POST("animal/", animal.ApiCreateHandler) //
	r.POST("animal/update/:id", animal.ApiUpdateHandler)
	r.GET("animal/delete/:id", animal.ApiDeleteHandler)

	return r
}
