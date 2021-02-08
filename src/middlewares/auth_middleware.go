package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/domain/repository/rest"
	"github.com/igson/banking/src/errors"
)

var (
	authRepo = rest.NewRemoteAuthRepository()
)

func Authorization(routerName string) gin.HandlerFunc {

	return func(c *gin.Context) {
		fmt.Println("-------> Authorization()")
		urlParams := make(map[string]string)

		authHeader := c.Request.Header.Get("Authorization")
		username := c.Request.Header.Get("username")

		fmt.Println(username)

		for k := range c.Request.URL.Query() {
			urlParams[k] = c.Request.URL.Query().Get(k)
			fmt.Println("Parameters ---> ", urlParams[k])
		}

		if authHeader != "" {

			token := getTokenFromHeader(authHeader)

			isAuthorized := authRepo.IsAuthorized(token, routerName, urlParams)

			if isAuthorized {
				c.Next()
			} else {
				appError := errors.NewStatusForbiddenError()
				c.JSON(appError.StatusCode, appError.Message)
				c.Abort()
				return
			}

		} else {
			errToken := errors.NewValidationError("Token ausente")
			c.JSON(errToken.StatusCode, errToken.Message)
			c.Abort()
			return
		}
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
