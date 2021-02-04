package routers

import (
	"net/http"

	"github.com/igson/banking/src/controller"
	"github.com/igson/banking/src/datasources/banking"
	"github.com/igson/banking/src/domain/repository"
	"github.com/igson/banking/src/domain/service"
)

var (
	customerRepo       = repository.NewCustomerRepository(banking.GetDbClient())
	customerService    = service.NewCustomerService(customerRepo)
	customerController = controller.NewCustomerController(customerService)
)

var rotasCustomer = []Rota{

	{
		URI:                "/customers",
		Metodo:             http.MethodGet,
		Funcao:             customerController.GetAllCustomers,
		RequerAutenticacao: false,
	},
	{
		URI:                "/customers/:customer_id",
		Metodo:             http.MethodGet,
		Funcao:             customerController.GetCustomer,
		RequerAutenticacao: false,
	},
}
