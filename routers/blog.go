package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//管理后台首页
func GetBlogIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/index.html", gin.H{
		"title": "GIN: 测试加载HTML模板",
	})
}
