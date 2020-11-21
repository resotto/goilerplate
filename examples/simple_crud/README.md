# WIP
(from https://github.com/resotto/goilerplate/issues/2#issuecomment-729662469, by @resotto)

## How to start with Goilerplate

With Goilerplate, you can start your project smoothly.

For explanation, let's create simple "CR" part of CRUD of following specifications with Goilerplate.

Specifications:
- There are three entities such as Customer, Product, and Order.
- Order aggregates Customer and Product (Order is Aggregate Root).
- There is only one usecase to create an order.

NOTICE:
- For convenience, the minimum codes are shown here.
- For convenience, there are no test codes in this explanation.

First of all, please prepare .go files with following package layout.

### Package Layout
```zsh
.
└── internal
    └── app
        ├── adapter
        │   ├── controller.go                 # Controller
        │   └── repository                    # Repository Implementation
        │       ├── customer.go
        │       ├── product.go
        │       └── order.go
        ├── application
        │   └── usecase                       # Usecase
        │       └── createOrder.go
        └── domain
            ├── customer.go                   # Entity
            ├── product.go                    # Entity
            ├── order.go                      # Entity
            └── repository                    # Repository Interface
                ├── customer.go
                ├── product.go
                └── order.go
```

### Define Entities

Secondly, let's create entities, Customer, Product, and Order.

```go
// customer.go
package domain

type Customer struct {
	ID string
	Name string
}
```

```go
// product.go
package domain

type Product struct {
	ID string
	Price int
}
```

```go
// order.go
package domain

type Order struct {
	ID string
	Customer Customer
	Product Product
}
```

### Define Repository Interfaces

After defining entities, let's prepare their repositories in `domain` package.

```go
// customer.go
package repository

type ICustomer interface {
	Get(id string) domain.Customer
}
```

```go
// product.go
package repository

type IProduct interface {
	Get(id string) domain.Product
}
```

```go
// order.go
package repository

type IOrder interface {
	Save(order Order)
}
```

### Define Usecase

And then, let's prepare the usecase of creating order.

```go
// createOrder.go
package usecase

import (
	"domain"            // simplified for convenience
	"domain/repository" // simplified for convenience
)

type CreateOrderArgs struct {
	CustomerID         string
	ProductID          string
	CustomerRepository repository.ICustomer
	ProductRepository  repository.IProduct
	OrderRepository    repository.IOrder
}

func CreateOrder(args CreateOrderArgs) domain.Order {
	customer := args.CustomerRepository.Get(args.CustomerID)
	product := args.ProductRepository.Get(args.ProductID)
	order := domain.Order{
		ID: "123",
		Customer: customer,
		Product: product,
	}
	args.OrderRepository.Save(order)
	return order
}
```

### Define Repository Implementations

After preparing the usecase, let's implement repository interfaces in `adapter` package.

However, this part is omitted here for convenience.

```go
// order.go
package repository

import (
	"domain" // simplified for convenience
)

type Order struct{}

func (o Order) Save(order domain.Order) {
	// omitted here for convenience
}
```

### Define Controller

Finally, let's define controller to call the usecase of creating an order.

```go
// controller.go
package adapter

import (
	"repository" // simplified for convenience
	"usecase"    // simplified for convenience

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
```
