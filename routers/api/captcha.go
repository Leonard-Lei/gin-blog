package api

import (
	"net/http"

	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/util"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	ImageUrl  string `json:"imageUrl"`
}

var GetCaptcha = func(context *gin.Context) {

	appG := app.Gin{C: context}

	captchaRes := CaptchaResponse{}
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	if d.CaptchaId != "" {
		captchaRes.CaptchaId = d.CaptchaId
		captchaRes.ImageUrl = "/show/" + d.CaptchaId + ".png"
	} else {
		appG.Response(http.StatusInternalServerError, e.ERROR_CAPTCHA_QUERY_PARAM_ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"CaptchaId": captchaRes.CaptchaId,
		"ImageUrl":  captchaRes.ImageUrl,
	})
}

var VerifyCaptcha = func(context *gin.Context) {

	appG := app.Gin{C: context}

	captchaId := context.Request.FormValue("captchaId")
	value := context.Request.FormValue("value")
	if captchaId == "" || value == "" {
		appG.Response(http.StatusOK, e.ERROR_CAPTCHA_QUERY_PARAM_ERROR, nil)
		return
	} else {
		if captcha.VerifyString(captchaId, value) {
			appG.Response(http.StatusOK, e.SUCCESS, nil)
			return
		} else {
			appG.Response(http.StatusOK, e.ERROR, nil)
			return
		}
	}
}

var GetCaptchaPng = func(context *gin.Context) {
	source := context.Param("source")
	logging.Info("GetCaptchaPng : " + source)
	util.CaptchaServeHTTP(context.Writer, context.Request)
}
