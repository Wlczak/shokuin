package auth

import (
	"net/http"
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

			if err != nil {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				return
			}

			_, claims, _ := utils.DecodeToken(token)

			authLevel := claims["auth_level"].(utils.AuthLevel)

			if authLevel < utils.AuthLevelUser {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
				return
			}
		}

	case utils.AuthLevelAdmin:
		return func(c *gin.Context) {

		}

	default:
		return func(c *gin.Context) {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

	}
}
