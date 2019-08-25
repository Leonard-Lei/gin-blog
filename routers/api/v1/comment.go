package v1

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"encoding/json"
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"gin-blog/service/article_service"
	"gin-blog/service/comment_service"
)

// @Summary 获取一段评论
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/comment/{id} [get]
func GetComment(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	commentService := comment_service.Comment{ID: id}
	exists, err := commentService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_COMMENT_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_COMMENT, nil)
		return
	}

	comment, err := commentService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_COMMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, comment)
}

// @Summary 获取多篇评论
// @Produce  json
// @Param tag_id body int false "ID"
// @Param category_id body int false "State"
// @Param state body int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/comments [get]
func GetComments(c *gin.Context) {

	appG := app.Gin{C: c}

	valid := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	articleId := -1
	if arg := c.Query("article_id"); arg != "" {
		articleId = com.StrTo(arg).MustInt()
		valid.Min(articleId, -1, "article_id")
	}

	pageNum := -1
	if arg := c.Query("page_num"); arg != "" {
		pageNum = com.StrTo(arg).MustInt()
		valid.Min(pageNum, -1, "page_num")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	commentService := comment_service.Comment{
		ArticleID: articleId,
		State:     state,
		PageNum:   util.GetPage(pageNum),
		PageSize:  setting.AppSetting.PageSize,
	}

	total, err := commentService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_COMMENT_FAIL, nil)
		return
	}

	comments, err := commentService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_COMMENTS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = comments
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddCommentForm struct {
	ArticleID int    `form:"article_id" valid:"Required;Min(1)"`
	Content   string `form:"content" valid:"Required;MaxSize(102400)"`
	CreateBy  int    `form:"create_by"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary 增加评论
// @Produce  json
// @Param article_id body int true "ArticleID"
// @Param content body string true "Content"
// @Param create_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddComment(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddCommentForm
	)

	body := make([]byte, 1024)
	n, _ := c.Request.Body.Read(body)
	fmt.Println(string(body[0:n]))
	//string 转json 再转 form
	err := json.Unmarshal([]byte(string(body[0:n])), &form)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_JSON_PARAMS, nil)
		return
	}
	httpCode, errCode := app.BindAndValid(c, form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	articleService := article_service.Article{ID: form.ArticleID}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	commentService := comment_service.Comment{
		ArticleID: form.ArticleID,
		Content:   form.Content,
		State:     form.State,
		CreateBy:  form.CreateBy,
	}
	if err := commentService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_COMMENT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
