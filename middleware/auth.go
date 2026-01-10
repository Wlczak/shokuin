package middleware

import (
	"net/http"

	"github.com/wlczak/shokuin/logger"
	"github.com/wlczak/shokuin/utils"

	"github.com/gin-gonic/gin"
)

func Auth(lvl utils.AuthLevel) gin.HandlerFunc {
	return auth(lvl, false)
}

func ApiAuth(lvl utils.AuthLevel) gin.HandlerFunc {
	return auth(lvl, true)
}

func auth(lvl utils.AuthLevel, isApi bool) gin.HandlerFunc {
	switch lvl {
	case utils.AuthLevelNone:
		return func(c *gin.Context) {
			c.Next()
		}

	case utils.AuthLevelUser:
		return func(c *gin.Context) {
			token, err := c.Cookie("SHOKUIN_JWT")
			zap := logger.GetLogger()
			if err != nil {
				zap.Error(err.Error())
				redirect(c, isApi)
				c.Abort()
				return
			}

			_, claims, err := utils.DecodeToken(token)

			if err != nil {
				zap.Error(err.Error())
				redirect(c, isApi)
				c.Abort()
				return
			}

			if claims.Auth_level == 0 {
				redirect(c, isApi)
				c.Abort()
				return
			}

			authLevel := claims.Auth_level

			if authLevel < utils.AuthLevelUser {
				redirect(c, isApi)
				c.Abort()
				return
			}
		}

	case utils.AuthLevelAdmin:
		return func(c *gin.Context) {
			token, err := c.Cookie("SHOKUIN_JWT")
			zap := logger.GetLogger()
			if err != nil {
				zap.Error(err.Error())
				redirect(c, isApi)
				c.Abort()
				return
			}

			_, claims, err := utils.DecodeToken(token)

			if err != nil {
				zap.Error(err.Error())
				redirect(c, isApi)
				c.Abort()
				return
			}

			if claims.Auth_level == 0 {
				redirect(c, isApi)
				c.Abort()
				return
			}

			authLevel := claims.Auth_level

			if authLevel < utils.AuthLevelAdmin {
				redirect(c, isApi)
				c.Abort()
				return
			}
		}

	default:
		return func(c *gin.Context) {
			redirect(c, isApi)
		}

	}
}

func redirect(c *gin.Context, isApi bool) {
	if isApi {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	} else {

		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}
