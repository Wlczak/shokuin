package auth

import (
	"net/http"
	"wlczak/shokuin/logger"
	"wlczak/shokuin/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(lvl utils.AuthLevel) gin.HandlerFunc {
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
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				c.Abort()
				return
			}

			_, claims, err := utils.DecodeToken(token)

			if err != nil {
				zap.Error(err.Error())
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				c.Abort()
				return
			}

			if claims["auth_level"] == nil {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				c.Abort()
				return
			}

			authLevel := claims["auth_level"].(utils.AuthLevel)

			if authLevel < utils.AuthLevelUser {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
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
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				c.Abort()
				return
			}

			_, claims, err := utils.DecodeToken(token)

			if err != nil {
				zap.Error(err.Error())
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				c.Abort()
				return
			}

			if claims["auth_level"] == nil {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				c.Abort()
				return
			}

			authLevel := claims["auth_level"].(utils.AuthLevel)

			if authLevel < utils.AuthLevelAdmin {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				c.Abort()
				return
			}
		}

	default:
		return func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

	}
}
