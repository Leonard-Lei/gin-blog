package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "gin-blog/docs"
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/upload"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	//新增获取token的方法
	r.GET("/auth", api.GetAuth)

	// 加载static文件夹下所有的文件
	r.LoadHTMLGlob("views/**/*")
	//swagger接口
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//图片上传
	r.POST("/upload", api.UploadImage)

	admin := r.Group("/admin")
	{
		//后台首页
		admin.GET("/index", GetAdminIndex)
		//后台登陆页面
		admin.GET("/login", GetLogin)
		//后台博客列表
		admin.GET("/list", GetAdminBlogList)
		//写博客页面
		admin.GET("/writeBlog", GetWriteBlog)
	}

	blog := r.Group("/blog")
	{
		//博客首页
		blog.GET("/index", GetBlogIndex)
		//博客首页
		blog.GET("/detail", GetBlogDetail)
		//博客友情链接
		blog.GET("/link", GetBlogLink)
		//博客留言
		blog.GET("/gustbook", GetBlogGustbook)
		//博客归档
		blog.GET("/archives", GetBlogArchives)
		//博客归档
		blog.GET("/hardware", GetBlogArchives) //博客归档
		blog.GET("/software", GetBlogArchives) //博客归档
		blog.GET("/life", GetBlogArchives)
		//博客里程碑
		blog.GET("/roadmap", GetBlogRoadmap)
		//博客搜索
		blog.GET("/search", GetBlogSearch)
		//关于我们
		blog.GET("/about", GetBlogAbout)

	}

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
