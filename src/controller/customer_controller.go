package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/igson/banking/src/errors"
	"github.com/igson/banking/src/interfaces"
)

type CustomerController interface {
	GetAllCustomers(ctx *gin.Context)
	GetCustomer(ctx *gin.Context)
}

type customerController struct {
	service interfaces.ICustomerService
}

//NewCustomerController construtor pra injeção das dependências
func NewCustomerController(customerService interfaces.ICustomerService) CustomerController {
	return &customerController{
		service: customerService,
	}
}

func (c *customerController) GetAllCustomers(ctx *gin.Context) {

	status := ctx.Param("status")

	if status == "" {
		restErr := errors.NewBadRequestError("Status inválido")
		ctx.JSON(restErr.StatusCode, restErr)
		return
	}

	customers, err := c.service.GetAllCustomer(status)

	if err != nil {
		ctx.JSON(err.StatusCode, err)
	} else {
		ctx.JSON(http.StatusOK, customers)
	}

}

func (c *customerController) GetCustomer(ctx *gin.Context) {

	id, error := strconv.ParseInt(ctx.Param("customer_id"), 10, 64)

	if error != nil {
		restErr := errors.NewBadRequestError("Customer ID deve ser número")
		ctx.JSON(restErr.StatusCode, restErr)
		return
	}

	customer, err := c.service.GetCustomer(id)

	if err != nil {
		ctx.JSON(err.StatusCode, err)
	} else {
		ctx.JSON(http.StatusOK, customer)
	}

}
