package main

/*

pedrooaj
878710

*/

import (

	"github.com/gin-gonic/gin"
	"github.com/Pedrooaj/Api_rest-Go/controllers/User"

)


func main(){
	app:= gin.Default();


	
	app.POST("/user", controllers.InsertUser)
	app.GET("/user/:id", controllers.GetUser)
	app.GET("/users", controllers.ListUsers)
	app.DELETE("/user/:id", controllers.DeleteUser)

	app.Run(":3000")
}