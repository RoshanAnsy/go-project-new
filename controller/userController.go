package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/roshanAnsy/go-project-new/config"
	"github.com/roshanAnsy/go-project-new/models"
)


func Signup(c *gin.Context){
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	var exitUser models.User
	user, err := DB.Find(exitUser, "email = ?", req.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "some thing is issue while fetching the data form db",
		})
		return
	}

	if user!=nil {
		c.JSON(200, gin.H{
			"message" : "user all ready exit",
			"success":"false",
		})
		return
	}

	user := User{Name:req.Name, Email:req.Email, Password:req.Password}
	result, err := DB.Create(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "some thing is issue while creating the user",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"message" : "user created successfully",
		"success":"true",
		"user": result,
	})
	


}

func Login (c *gin.Context) {
	var req model.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message" : "user logged in successfully",
	})
}

func  GetUser(c* gin.Context){
	c.JSON(200, gin.H{
        "message" : "get user details",
    })
}