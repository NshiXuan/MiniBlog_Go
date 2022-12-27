package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

// 首页
func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

// 注册
func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("username: %v\n", username)

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)

	// 重定向 不能使用200
	c.Redirect(301, "/")
}

// 登录
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", "")
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在!")
		fmt.Println("用户名不存在!")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}
}

// 博客控制器
// 博客列表
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

// 添加博客
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)
	c.Redirect(302, "/post_index")
}

// 跳转到添加博客
func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

// 博客详情
func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s) // 转成int类型
	p := dao.Mgr.GetPost(pid)
	content := blackfriday.Run([]byte(p.Content))
	fmt.Printf("content: %v\n", content)

	c.HTML(200, "postDetail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}

// 其它
func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}
