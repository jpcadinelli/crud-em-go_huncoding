package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", user.FindUserById)
	r.GET("/user/:userEmail", user.FindUserByEmail)
	r.POST("/user/", user.CreateUser)
	r.PUT("/user/:userId", user.UpdateUser)
	r.DELETE("/user/:userId", user.DeleteUser)
}
