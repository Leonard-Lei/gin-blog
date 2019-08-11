package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//后台首页
func GetAdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title": "GIN: 首页",
	})
}

//后台登陆页面
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.html", gin.H{
		"title": "GIN: 登录界面",
	})
}

//后台博客列表
func GetAdminBlogList(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/list.html", gin.H{
		"title": "GIN: 博客列表",
	})
}

//写博客
func GetWriteBlog(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/writeBlog.html", gin.H{
		"title": "GIN: 写博客页面",
	})
}