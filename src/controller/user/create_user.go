package user

import (
	"crud-em-go_huncoding/src/configuration/rest_err"
	"crud-em-go_huncoding/src/controller/model/request"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateUser(ginctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ginctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("there are some incorrect filds: %s\n", err.Error()))
		ginctx.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
