package jwt

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrorNoLogin = errors.New("请登录后操作! ")
)

type IStore interface {
	// IsBlackList 判断是否是黑名单
	IsBlackList(ctx context.Context, token string) bool
}

// Auth 授权中间件
func Auth(secret string, guard string, store IStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetJwtToken(c)

		claims, err := verify(guard, secret, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": err.Error()})
			c.Abort()
			return
		}

		// 这里还需要验证 token 黑名单
		if store.IsBlackList(c.Request.Context(), token) {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "请登录再试."})
			c.Abort()
			return
		}

		// 设置登录用户ID
		if uid, err := strconv.Atoi(claims.Id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "解析 jwt 失败."})
			c.Abort()
			return
		} else {
			c.Set("__UID__", uid)
		}

		// 记录 jwt 相关信息
		c.Set("jwt", map[string]string{
			"token":      token,
			"expires_at": strconv.Itoa(int(claims.ExpiresAt)),
		})

		c.Next()
	}
}

func verify(guard string, secret string, token string) (*AuthClaims, error) {

	if token == "" {
		return nil, ErrorNoLogin
	}

	claims, err := ParseToken(token, secret)
	if err != nil {
		return nil, err
	}

	// 判断权限认证守卫是否一致
	if claims.Guard != guard || claims.Valid() != nil {
		return nil, ErrorNoLogin
	}

	return claims, nil
}
