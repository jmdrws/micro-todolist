package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(),middleware.InitMiddleware(service),middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession",store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("pong",func(c *gin.Context){
			c.JSON(200,"success")
		})
		v1.POST("/user/register",handlers.UserRegister)
		v1.POST("/user/login",handlers.UserLogin)
	}
	return ginRouter
}