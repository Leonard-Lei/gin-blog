package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//管理后台首页
func GetAdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title": "GIN: 测试加载HTML模板",
	})
}
