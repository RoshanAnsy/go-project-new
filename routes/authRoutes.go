package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/roshanAnsy/go-project-new/controller"
)

func AuthRoutes(r *gin.Engine){
	r.POST("/signup",controllers.Signup);
	r.POST("/login",controllers.Login);
}