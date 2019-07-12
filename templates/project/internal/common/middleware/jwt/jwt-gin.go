/**
 * Created by Wangwei on 2019-06-06 12:04.
 */

package jwt

import (
	"goadmin/internal/common"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	jwt = &JWT{}
)

func jwtTokenAbort(c *gin.Context, msg string) {
	common.ResponseWith(c, 501, msg)
	c.Abort()
}

// 请求拦截器中间件，验证token，后台管理专用
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if claims := getCustomClaims(c); claims != nil {
			c.Set("claims", claims)
			c.Next()
		}
	}
}

// 解析token中的信息
func getCustomClaims(c *gin.Context) *CustomClaims {
	clientId := c.Request.Header.Get("X-Client-Id")
	clientId = common.FormatClientId(clientId)
	if strings.TrimSpace(clientId) == "" {
		jwtTokenAbort(c, "client-id不能为空")
		return nil
	}

	authHeader := c.Request.Header.Get("X-Token")
	if authHeader == "" {
		jwtTokenAbort(c, "token不能为空")
		return nil
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		jwtTokenAbort(c, "token无效")
		return nil
	}

	claims, err := jwt.ParseToken(parts[1])
	if err != nil {
		jwtTokenAbort(c, err.Error())
		return nil
	}

	return claims
}
