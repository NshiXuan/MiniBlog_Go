package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	// e.GET("/index", controller.ListUser)
	// 注册
	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.Register)

	// 登录
	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	// 博客
	e.GET("/post_index", controller.GetPostIndex)
	e.POST("/post", controller.AddPost)
	e.GET("/post", controller.GoAddPost)
	e.GET("/post_detail", controller.PostDetail)

	// 首页
	e.GET("/", controller.Index)
	e.Run()
}
