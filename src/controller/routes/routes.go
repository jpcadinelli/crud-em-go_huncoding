package routes

import (
	"crud-em-go_huncoding/src/controller/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", user.FindUserById)
	r.GET("/user/:userEmail", user.FindUserByEmail)
	r.POST("/user/", user.CreateUser)
	r.PUT("/user/:userId", user.UpdateUser)
	r.DELETE("/user/:userId", user.DeleteUser)
}
