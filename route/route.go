package route

import (
	. "../controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	structure := route.Group("/structure")
	{
		structure.GET("/", StructureHome)
		structure.POST("/create", StructureCreate)
		structure.GET("/show/:id", StructureShow)
		structure.PUT("/update", StructureUpdate)
		structure.DELETE("/delete/:id", StructureDelete)
		structure.POST("/delete", StructureDeleteMultiple)
	}

	return route
}
