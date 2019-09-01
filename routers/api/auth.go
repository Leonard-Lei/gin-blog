package api

import (
	//"gin-blog/pkg/logging"
	//"log"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Unknwon/com"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gin-blog/pkg/setting"

	"gin-blog/pkg/app"
	//"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"gin-blog/service/auth_service"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary 登陆
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth [get]
func GetLogin(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// @Summary 获取用户
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/auth/{id} [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{ID: id}
	exists, err := authService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MESSAGE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_MESSAGE, nil)
		return
	}

	auth, err := authService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, auth)
}

// @Summary 获取多篇用户
// @Produce  json
// @Param tag_id body int false "ID"
// @Param category_id body int false "State"
// @Param state body int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/auths [get]
func GetAuths(c *gin.Context) {

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

	authService := auth_service.Auth{
		State:    state,
		PageNum:  util.GetPage(pageNum),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := authService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_MESSAGE_FAIL, nil)
		return
	}

	auths, err := authService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_MESSAGES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = auths
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddAuthForm struct {
	Username string `form:"username" valid:"Required;MaxSize(100)"`
	Password string `form:"password" valid:"Required;MaxSize(255)"`
	CreateBy int    `form:"create_by"`
	State    int    `form:"state" valid:"Range(0,1)"`
}

// @Summary 增加用户
// @Produce  json
// @Param article_id body int true "ArticleID"
// @Param content body string true "Content"
// @Param create_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func AddAuth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddAuthForm
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

	authService := auth_service.Auth{
		Username: form.Username,
		Password: form.Password,
		State:    form.State,
		CreateBy: form.CreateBy,
	}
	if err := authService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditAuthForm struct {
	ID       int `form:"id" valid:"Required;Min(1)"`
	UpdateBy int `form:"update_by"`
	State    int `form:"state"`
}

// @Summary 修改文章用户
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Param article_id body string true "ArticleId"
// @Param state body int false "State"
// @Param update_by body string true "UpdateBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditAuth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditAuthForm{ID: com.StrTo(c.Param("id")).MustInt()}
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

	authService := auth_service.Auth{
		ID:       form.ID,
		UpdateBy: form.UpdateBy,
		State:    form.State,
	}

	exists, err := authService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MESSAGE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = authService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 删除用户
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/auths/{id} [delete]
func DeleteAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
	}

	authService := auth_service.Auth{ID: id}
	exists, err := authService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_MESSAGE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_MESSAGE, nil)
		return
	}

	if err := authService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_MESSAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 验证码
// @Accept  json
// @Produce  json
// @Router /vertify [GET]
func Verify(c *gin.Context) {
	util.LoadVerify(c)
}

// //鉴权接口
// func Auth() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		//鉴权的一般思路
// 		//1、判断当前的角色id roleid
// 		//2、获取该角色所有的权限roleauth,获取系统全部权限allauth
// 		//3、获取当前的uri := ctx.Request.RequestURI
// 		//4、判断uri是否在allauth中,如果不在,则表面无需鉴权,通过,
// 		//   否则检测uri是否在roleauth中,如果是则通过,否则鉴权失败
// 		uri := ctx.Request.RequestURI
// 		isapi := !strings.Contains(uri, ".shtml")
// 		ispage := !isapi
// 		auths := AllAuth()
// 		var exist bool = true

// 		_, exist = auths[uri]
// 		//如果不存在,说明这个是不需要做权限校验的
// 		if !exist {
// 			ctx.Next()
// 			return
// 		}

// 		iroleId := LoadRoleId(ctx)
// 		roleId := 0
// 		if nil != iroleId {
// 			roleId = iroleId.(int)
// 		} else {
// 			//这里设置一下默认的id
// 		}

// 		//没有登陆则直接返回
// 		if roleId == 0 {
// 			if ispage {
// 				ctx.HTML(http.StatusOK, "public/error.html", gin.H{"msg": "你没有权限进行该操作"})
// 				ctx.Abort()
// 			} else {

// 				ResultFail(ctx, "鉴权失败")
// 				ctx.Abort()
// 			}
// 			return
// 		}
// 		//获取角色map
// 		roleAuth := RoleAuth(roleId)
// 		_, exist = roleAuth[uri]

// 		//如果不存在,说明这个没有权限
// 		if exist {
// 			ctx.Next()
// 			return
// 		}

// 		if ispage {
// 			ctx.HTML(http.StatusOK, "public/error.html", gin.H{"msg": "你没有权限进行该操作"})
// 			ctx.Abort()
// 		} else {
// 			ResultFail(ctx, "鉴权失败")
// 			ctx.Abort()
// 		}
// 		return
// 	}
// }

// //这个参数在ResController初始化时处理
// var allAuth map[string]int64 = make(map[string]int64)

// //将auth存储起来
// func AllAuth(auth ...map[string]int64) map[string]int64 {
// 	if len(auth) > 0 {
// 		allAuth = auth[0]
// 		return nil
// 	} else {
// 		return allAuth
// 	}
// }

// //将auth存储起来
// var mapRoleAuth map[int]map[string]int64 = make(map[int]map[string]int64)

// //这个参数在RoleController初始化时处理
// func RoleAuth(roleId int, auth ...map[string]int64) map[string]int64 {
// 	if len(auth) > 0 {
// 		mapRoleAuth[roleId] = auth[0]
// 		return nil
// 	} else {
// 		return mapRoleAuth[roleId]
// 	}

// }
