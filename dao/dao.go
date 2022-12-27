package dao

import (
	"blog/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	Register(user *model.User)
	Login(username string) model.User

	// 博客DAO
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:admin123@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

// 注册
func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

// 登录
func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

/*   博客操作  */
// 添加博客
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}

// 获取全部博客
func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}

// 通过id获取博客
func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}
