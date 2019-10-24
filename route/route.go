package route

import (
	. "../controller"
	. "../middleware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	route.Use(ApiMiddleware(db))
	route.Use(CORS())
	structure := route.Group("/structure")
	{
		structure.GET("/", StructureHome)
		structure.POST("/create", StructureCreate)
		structure.GET("/show/:id", StructureShow)
		structure.PUT("/update/:id", StructureUpdate)
		structure.DELETE("/delete/:id", StructureDelete)
		structure.POST("/delete", StructureDeleteMultiple)
	}

	return route
}
