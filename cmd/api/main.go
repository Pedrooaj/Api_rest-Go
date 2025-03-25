package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	

	app:= gin.Default();
	app.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"Message": "Hello world",
		})
	})
	app.Run(":3000")
}