package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go-chat/app/helper"
	"go-chat/config"
)

// ApiAuth 授权中间件
// guard 授权守卫
func JwtAuth(conf *config.Config, guard string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getAuthToken(c)
		if token == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 40001,
				"msg":  "请登录后操作!",
			})
			return
		}

		jwt, err := helper.ParseJwtToken(conf, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 40002,
				"msg":  "Token 信息验证失败!",
			})
			return
		}

		// 判断权限认证守卫是否一致
		if jwt.Guard != guard {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 40002,
				"msg":  "Token 信息验证失败!",
			})
			return
		}

		// 黑名单判断

		c.Set("jwt", jwt)
		c.Set("user_id", jwt.UserId)

		c.Next()
	}
}

// getAuthToken 获取登录授权token
func getAuthToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	token = strings.TrimLeft(token, "Bearer")
	token = strings.TrimSpace(token)

	// Headers 中没有授权信息则读取 url 中的 token
	if len(token) == 0 {
		token = c.DefaultQuery("token", "")
	}

	if len(token) == 0 {
		token = c.DefaultPostForm("token", "")
	}

	return token
}