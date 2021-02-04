package routers

import "github.com/gin-gonic/gin"

/* var clienteController = controller.NewClienteController(service.NewClienteService())

func (r Rotas) registrarRotasCliente(rota *gin.RouterGroup) {
	rota.GET("/clientes/:cliente_id", clienteController.GetByID)
	rota.POST("/clientes", clienteController.Create)
}
*/

//Rota objeto de configuração das rotas
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(ctx *gin.Context)
	RequerAutenticacao bool
}

//Configurar carregar lista de rotas
func Configurar(r *gin.Engine) *gin.Engine {
	rotas := rotasAccounts
	rotas = append(rotas, rotasCustomer...)
	//func(httpMethod string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	for _, rota := range rotas {
		r.Handle(rota.Metodo, rota.URI, rota.Funcao)
	}

	return r
}
