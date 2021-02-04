package middleware

import "github.com/gin-gonic/gin"

//AuthMiddleware filtros de acesso
type AuthMiddleware struct {
}

func (a AuthMiddleware) authorizationHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentRote := ctx.
	}
}
