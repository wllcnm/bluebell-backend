package routers

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置为发布者模式
	}

	r := gin.New()

	//加载日志中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("templates/index.html") //加载html
	r.Static("/static", "./static")         //加载静态文件

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	//注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", controller.LoginHandler)             //登录路由
		v1.POST("/signup", controller.SignUpHandler)           //注册路由
		v1.GET("/refer_token", controller.RefreshTokenHandler) //token刷新

		//v1.GET("/posts", controller.CreatePostHandler)
	}

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.POST("/post", controller.CreatePostHandler) //创建帖子
	}
	return r
}
