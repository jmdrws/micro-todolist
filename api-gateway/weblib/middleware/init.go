package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
//接收服务实例，	并存在gin.Key 中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		//将实例存在gin.Keys中
		context.Keys = make(map[string]interface{})
		context.Keys["userService"] = service[0]
		context.Next()
	}
}

//错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(200,gin.H{
					"code":404,
					"msg":fmt.Sprintf("%s",r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}