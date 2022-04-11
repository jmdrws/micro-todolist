package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	userService := ginCtx.Keys["userService"].(service.UserService)
	userRes,err := userService.UserRegister(context.Background(),&userReq)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK,gin.H{"data":userRes})
}

func UserLogin(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	userService := ginCtx.Keys["userService"].(service.UserService)
	userRes,err := userService.UserLogin(context.Background(),&userReq)
	PanicIfUserError(err)
	token,err := utils.GenerateToken(uint(userRes.UserDetail.ID))
	ginCtx.JSON(http.StatusOK, gin.H{
		"code": userRes.Code,
		"msg":  "成功",
		"data": gin.H{
			"user":  userRes.UserDetail,
			"token": token,
		},
	})
}