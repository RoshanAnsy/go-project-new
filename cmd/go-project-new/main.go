package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/roshanAnsy/go-project-new/config"
	"github.com/roshanAnsy/go-project-new/routes"
)



func main(){
	//db connected to database
	config.ConnectDB();
	r := gin.Default();
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "hello from server guys",
		})
	})
	
	//userRoutes(r)
	routes.userRoutes(r);

	r.Run(":5000")
	fmt.Println("hello from server guys")
	
}