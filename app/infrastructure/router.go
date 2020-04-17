package router

import (
	users "github.com/dev-sota/gin-api/app/adapter/controllers"
	"github.com/gin-gonic/gin"
)

// Init is initialize routes
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	u := r.Group("/users")
	{
		ctrl := users.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}
	return r
}
