package main

import (
	"github.com/Pedrooaj/Api_rest-Go/routes/User"
	"github.com/gin-gonic/gin"
)


func main(){
	app:= gin.Default();
	
	Routes.UserRoutes(app)
	
	app.Run(":3000")
}