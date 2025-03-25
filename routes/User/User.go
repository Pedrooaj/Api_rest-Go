package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Pedrooaj/Api_rest-Go/controllers/User"
)

func UserRoutes(router *gin.Engine){
	router.POST("/user", controllers.InsertUser)
	router.GET("/user/:id", controllers.GetUser)
	router.GET("/users", controllers.ListUsers)
	router.DELETE("/user/:id", controllers.DeleteUser)
}