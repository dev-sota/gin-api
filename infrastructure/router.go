package infrastructure

import (
	"github.com/dev-sota/gin-api/controllers"
	"github.com/gin-gonic/gin"
)

// RouterInit is initialize routes
func RouterInit() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.Index)
	return r
}
