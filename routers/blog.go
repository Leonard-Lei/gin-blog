package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//博客首页
func GetBlogIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/index.html", gin.H{
		"title": "GIN: 博客首页",
	})
}

//博客列表
func GetBlogList(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/list.html", gin.H{
		"title": "GIN: 博客列表",
	})
}

//博客详情
func GetBlogDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/detail.html", gin.H{
		"title": "GIN: 博客详情",
	})
}
