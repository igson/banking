package app

import (
	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/app/routers"
)

//GerarRotas iniciar rotas
func GerarRotas() *gin.Engine {
	r := gin.Default()
	return routers.Configurar(r)
}
