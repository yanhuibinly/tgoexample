package route

import (
	"controller"
	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine) {
	//内部接口
	router.GET("/hello", controller.Hello)
}
