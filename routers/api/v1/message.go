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
	"gin-blog/service/message_service"
)

// @Summary 获取一段留言
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/message/{id} [get]
func GetMessage(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	messageService := message_service.Message{ID: id}
	exists, err := messageService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MESSAGE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_MESSAGE, nil)
		return
	}

	message, err := messageService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, message)
}

// @Summary 获取多篇留言
// @Produce  json
// @Param tag_id body int false "ID"
// @Param category_id body int false "State"
// @Param state body int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/messages [get]
func GetMessages(c *gin.Context) {

	appG := app.Gin{C: c}

	valid := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
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

	messageService := message_service.Message{
		State:    state,
		PageNum:  util.GetPage(pageNum),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := messageService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_MESSAGE_FAIL, nil)
		return
	}

	messages, err := messageService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_MESSAGES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = messages
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddMessageForm struct {
	Content  string `form:"content" valid:"Required;MaxSize(102400)"`
	CreateBy int    `form:"create_by"`
	State    int    `form:"state" valid:"Range(0,1)"`
}

// @Summary 增加留言
// @Produce  json
// @Param article_id body int true "ArticleID"
// @Param content body string true "Content"
// @Param create_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddMessage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddMessageForm
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

	messageService := message_service.Message{
		Content:  form.Content,
		State:    form.State,
		CreateBy: form.CreateBy,
	}
	if err := messageService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditMessageForm struct {
	ID       int `form:"id" valid:"Required;Min(1)"`
	UpdateBy int `form:"update_by"`
	State    int `form:"state"`
}

// @Summary 修改文章留言
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Param article_id body string true "ArticleId"
// @Param state body int false "State"
// @Param update_by body string true "UpdateBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditMessage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditMessageForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	body := make([]byte, 1024)
	n, _ := c.Request.Body.Read(body)
	//fmt.Println(string(body[0:n]))
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

	messageService := message_service.Message{
		ID:       form.ID,
		UpdateBy: form.UpdateBy,
		State:    form.State,
	}

	exists, err := messageService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MESSAGE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = messageService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 删除留言
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/messages/{id} [delete]
func DeleteMessage(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	messageService := message_service.Message{ID: id}
	exists, err := messageService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MESSAGE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_MESSAGE, nil)
		return
	}

	if err := messageService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
