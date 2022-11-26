package router

import (
	"github.com/gin-gonic/gin"
	"github.com/heqingbao/ginchat/docs"
	"github.com/heqingbao/ginchat/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	v1.GET("/helloworld", service.Helloworld)
	v1.GET("/user/getUserList", service.GetUserList)
	v1.GET("/user/createUser", service.CreateUser)
	v1.GET("/user/deleteUser", service.DeleteUser)
	v1.POST("/user/updateUser", service.UpdateUser)
	v1.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	// 发送消息
	v1.GET("/user/sendMsg", service.SendMsg)
	return r
}
