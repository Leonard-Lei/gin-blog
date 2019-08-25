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
	"gin-blog/service/category_service"
)

// @Summary 获取分类
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/category/{id} [get]
func GetCategory(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	categoryService := category_service.Category{ID: id}
	exists, err := categoryService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CATEGORY_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_CATEGORY, nil)
		return
	}

	category, err := categoryService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, category)
}

// @Summary 获取多个分类
// @Produce  json
// @Param category_id body int false "State"
// @Param state body int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/categorys [get]
func GetCategorys(c *gin.Context) {

	appG := app.Gin{C: c}

	valid := validation.Validation{}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state")
	}

	name := c.Query("name")

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

	categoryService := category_service.Category{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(pageNum),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := categoryService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_CATEGORY_FAIL, nil)
		return
	}

	categorys, err := categoryService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_CATEGORYS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = categorys
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddcategoryForm struct {
	Name     string `form:"name" valid:"Required;MaxSize(102400)"`
	CreateBy int    `form:"create_by"`
	State    int    `form:"state" valid:"Range(0,1)"`
}

// @Summary 增加评论
// @Produce  json
// @Param create_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddcategoryForm
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

	categoryService := category_service.Category{
		Name:     form.Name,
		State:    form.State,
		CreateBy: form.CreateBy,
	}

	// exists, err := categoryService.ExistByName()
	// if err != nil {
	// 	appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_CATEGORY_FAIL, nil)
	// 	return
	// }

	// if exists {
	// 	appG.Response(http.StatusOK, e.ERROR_ALREADY_EXIST_CATEGORY, nil)
	// 	return
	// }

	if err := categoryService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_CATEGORY_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
