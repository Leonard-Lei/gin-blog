package util

import (
	"gin-blog/pkg/setting"
	//"github.com/Unknwon/com"
	//"github.com/gin-gonic/gin"
)

// GetPage get page parameters
//func GetPage(c *gin.Context) int {
func GetPage(page int) int {
	result := 0
	//page := com.StrTo(c.Query("page")).MustInt()

	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	} else {
		result = -1
	}

	return result
}
