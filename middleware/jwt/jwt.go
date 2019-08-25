package jwt

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		//token := c.Query("token")
		uri := c.Request.RequestURI
		method := c.Request.Method
		fmt.Println(uri)
		if method == "GET" {
			if strings.Contains(uri, "/api/v1/articles") || strings.Contains(uri, "/api/v1/tags") {
				//if uri == "/api/v1/articles" || uri == "/api/v1/tags" {
				c.Next()
				return
			}
		}

		token := c.Request.Header.Get("token")

		// body := make([]byte, 1024)
		// n, _ := c.Request.Body.Read(body)
		// fmt.Println(string(body[0:n]))

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
