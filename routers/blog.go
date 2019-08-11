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

//博客详情
func GetBlogLink(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/link.html", gin.H{
		"title": "GIN: 博客友情链接",
	})
}

//博客留言
func GetBlogGustbook(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/guestbook.html", gin.H{
		"title": "GIN: 博客友情链接",
	})
}

//博客归档
func GetBlogArchives(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/archives.html", gin.H{
		"title": "GIN: 博客归档",
	})
}

//博客里程碑
func GetBlogRoadmap(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/roadmap.html", gin.H{
		"title": "GIN: 博客里程碑",
	})
}

//博客搜索
func GetBlogSearch(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/search.html", gin.H{
		"title": "GIN: 博客搜索",
	})
}

//博客搜索
func GetBlogAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/about.html", gin.H{
		"title": "GIN: 关于我们",
	})
}


