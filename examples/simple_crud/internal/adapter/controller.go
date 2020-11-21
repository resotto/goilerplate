package adapter

import (
	"examples/internal/adapter/repository"
	"examples/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

var (
	customerRepository = repository.Customer{}
	productRepository  = repository.Product{}
	orderRepository    = repository.Order{}
)

type Controller struct{}

func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	r.POST("/order", ctrl.createOrder)
	return r
}

func (ctrl Controller) createOrder(c *gin.Context) {
	customerID := c.Query("customerId")
	productID := c.Query("productId")
	args := usecase.CreateOrderArgs{
		CustomerID:         customerID,
		ProductID:          productID,
		CustomerRepository: customerRepository,
		ProductRepository:  productRepository,
		OrderRepository:    orderRepository,
	}
	order := usecase.CreateOrder(args)
	c.JSON(200, order)
}