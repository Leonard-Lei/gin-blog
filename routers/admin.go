package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//管理后台首页
func GetAdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title": "GIN: 首页",
	})
}

//管理后台首页
func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login.html", gin.H{
		"title": "GIN: 登录界面",
	})
}
