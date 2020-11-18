<h1 align="center">Goilerplate</h1>

<p align="center">
  <a href="https://github.com/resotto/goilerplate/actions"><img src="https://github.com/resotto/goilerplate/workflows/test/badge.svg" /></a>
  <a href="https://goreportcard.com/report/github.com/resotto/goilerplate"><img src="https://goreportcard.com/badge/github.com/resotto/goilerplate" /></a>
  <a href="https://pkg.go.dev/github.com/resotto/goilerplate"><img src="https://pkg.go.dev/badge/github.com/resotto/goilerplate" /></a>
  <a href="https://github.com/resotto/goilerplate/issues/1"><img src="https://img.shields.io/badge/chat-on%20issue-yellow"></a>
  <a href="https://github.com/resotto/goilerplate/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-GPL%20v3.0-brightgreen.svg" /></a>
</p>

<p align="center">
  Clean Boilerplate of Go, Domain-Driven Design, Clean Architecture, Gin and GORM.  
</p>

<p align="center">
  <img src="https://user-images.githubusercontent.com/19743841/93784289-cbd84200-fc67-11ea-997b-b99af8affe17.png">
</p>

---

What is Goilerplate?

- **Good example of Go with Clean Architecture.**
- **Rocket start guide of Go, Domain-Driven Design, [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), [Gin](https://github.com/gin-gonic/gin), and [GORM](https://github.com/go-gorm/gorm)**.

Who is the main user of Goilerplate?

- All kinds of Gophers (newbie to professional).

Why Goilerplate?

- **Easy-applicable boilerplate in Go.**

Note

- Default application/test code is trivial because you will write cool logic.
- [Public API of bitbank](https://github.com/bitbankinc/bitbank-api-docs/blob/master/public-api.md#general-endpoints), which is bitcoin exchange located in Tokyo, is used for some endpoints by default.

---

## Table of Contents

- [Getting Started](#getting-started)
- [go get Goilerplate via SSH](#go-get-goilerplate-via-ssh)
- [Endpoints](#endpoints)
- [Package Structure](#package-structure)
- [How to Cross the Border of Those Layers](#how-to-cross-the-border-of-those-layers)
- [Dependency Injection](#dependency-injection)
- [How to Start with Goilerplate](#how-to-start-with-goilerplate)
- [Testing](#testing)
- [Naming Convention](#naming-convention)
- [With Gochk](#with-gochk)
- [With PostgreSQL](#with-postgresql)
- [Feedbacks](#feedbacks)
- [License](#license)
- [Author](#author)

## Getting Started

```zsh
go get -u github.com/resotto/goilerplate        # might take few minutes
cd ${GOPATH}/src/github.com/resotto/goilerplate
go run cmd/app/main.go                          # from root directory
open http://0.0.0.0:8080
```

## `go get` Goilerplate via SSH

`go get` fetches GitHub repository via HTTPS by default. So you might fail `go get`:

```zsh
~  > go get -u github.com/resotto/goilerplate
# cd .; git clone -- https://github.com/resotto/goilerplate /Users/resotto/go/src/github.com/resotto/goilerplate
Cloning into '/Users/resotto/go/src/github.com/resotto/goilerplate'...
fatal: could not read Username for 'https://github.com': terminal prompts disabled
package github.com/resotto/goilerplate: exit status 128
```

If you `go get` GitHub repository via SSH, please run following command:

```zsh
git config --global url.git@github.com:.insteadOf https://github.com/
```

And then, please try [Getting Started](#getting-started) again.

## Endpoints

- With Template
  - `GET /`
    - NOTICE: Following path is from CURRENT directory, so please run Gin from root directory.
      ```go
      r.LoadHTMLGlob("internal/app/adapter/view/*")
      ```
- With Public API of bitbank
  - `GET /ticker`
  - `GET /candlestick`
    - NOTICE: This works from 0AM ~ 3PM (UTC) due to its API constraints.
- With PostgreSQL
  - [NOTICE: Please run postgres container first with this step.](#with-postgresql)
    - `GET /parameter`
    - `GET /order`

## Package Structure

```zsh
.
├── LICENSE
├── README.md
├── build                                     # Packaging and Continuous Integration
│   ├── Dockerfile
│   └── init.sql
├── cmd                                       # Main Application
│   └── app
│       └── main.go
├── internal                                  # Private Codes
│   └── app
│       ├── adapter
│       │   ├── controller.go                 # Controller
│       │   ├── postgresql                    # Database
│       │   │   ├── conn.go
│       │   │   └── model                     # Database Model
│       │   │       ├── card.go
│       │   │       ├── cardBrand.go
│       │   │       ├── order.go
│       │   │       ├── parameter.go
│       │   │       ├── payment.go
│       │   │       └── person.go
│       │   ├── repository                    # Repository Implementation
│       │   │   ├── order.go
│       │   │   └── parameter.go
│       │   ├── service                       # Application Service Implementation
│       │   │   └── bitbank.go
│       │   └── view                          # Templates
│       │       └── index.tmpl
│       ├── application
│       │   ├── service                       # Application Service Interface
│       │   │   └── exchange.go
│       │   └── usecase                       # Usecase
│       │       ├── addNewCardAndEatCheese.go
│       │       ├── ohlc.go
│       │       ├── parameter.go
│       │       ├── ticker.go
│       │       └── ticker_test.go
│       └── domain
│           ├── factory                       # Factory
│           │   └── order.go
│           ├── order.go                      # Entity
│           ├── parameter.go
│           ├── parameter_test.go
│           ├── person.go
│           ├── repository                    # Repository Interface
│           │   ├── order.go
│           │   └── parameter.go
│           └── valueobject                   # ValueObject
│               ├── candlestick.go
│               ├── card.go
│               ├── cardbrand.go
│               ├── pair.go
│               ├── payment.go
│               ├── ticker.go
│               └── timeunit.go
└── testdata                                  # Test Data
    └── exchange_mock.go
```

### ![#fffacd](https://via.placeholder.com/15/fffacd/000000?text=+) Domain Layer

- The core of Clean Architecture. It says "Entities".

### ![#f08080](https://via.placeholder.com/15/f08080/000000?text=+) Application Layer

- The second layer from the core. It says "Use Cases".

### ![#98fb98](https://via.placeholder.com/15/98fb98/000000?text=+) Adapter Layer

- The third layer from the core. It says "Controllers / Gateways / Presenters".

### ![#87cefa](https://via.placeholder.com/15/87cefa/000000?text=+) External Layer

- The fourth layer from the core. It says "Devices / DB / External Interfaces / UI / Web".
  - **We DON'T write much codes in this layer.**

<p align="center">
  <img src="https://user-images.githubusercontent.com/19743841/93830264-afa9c480-fcaa-11ea-9589-7c5308c291f4.jpg">
</p>
<p align="center">
  <a href="https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html">The Clean Architecture</a>
</p>

## How to Cross the Border of Those Layers

In Clean Architecture, there is [The Dependency Rule](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html#the-dependency-rule):

> This rule says that source code dependencies can only point inwards. Nothing in an inner circle can know anything at all about something in an outer circle.

In other words, **Dependency Injection** is required to follow this rule.

Therefore, please follow the next four steps:

1. Define Interface
1. Take Argument as Interface and Call Functions of It
1. Implement It
1. Inject Dependency

Here, I pick up the example of Repository.

### Repository

```zsh
.
└── internal
    └── app
        ├── adapter
        │   ├── controller.go    # 4. Dependency Injection
        │   └── repository
        │       └── parameter.go # 3. Implementation
        ├── application
        │   └── usecase
        │       └── parameter.go # 2. Interface Function Call
        └── domain
            ├── parameter.go
            └── repository
                └── parameter.go # 1. Interface
```

1. Interface at Domain Layer:

```go
package repository

import "github.com/resotto/goilerplate/internal/app/domain"

// IParameter is interface of parameter repository
type IParameter interface {
	Get() domain.Parameter
}
```

2. Usecase at Application Layer:

```go
package usecase

// NOTICE: This usecase DON'T depend on Adapter layer
import (
	"github.com/resotto/goilerplate/internal/app/domain"
	"github.com/resotto/goilerplate/internal/app/domain/repository"
)

// Parameter is the usecase of getting parameter
func Parameter(r repository.IParameter) domain.Parameter {
	return r.Get()
}
```

3. Implementation at Adapter Layer:

```go
package repository

// Parameter is the repository of domain.Parameter
type Parameter struct{}

// Get gets parameter
func (r Parameter) Get() domain.Parameter {
	db := postgresql.Connection()
	var param model.Parameter
	result := db.First(&param, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	return domain.Parameter{
		Funds: param.Funds,
		Btc:   param.Btc,
	}
}
```

4. Dependency Injection at Controller of Adapter Layer:

```go
package adapter

// NOTICE: Controller depends on INNER CIRCLE so it points inward (The Dependency Rule)
import (
	"github.com/gin-gonic/gin"
	"github.com/resotto/goilerplate/internal/app/adapter/repository"
	"github.com/resotto/goilerplate/internal/app/application/usecase"
)

var (
	parameterRepository = repository.Parameter{}
)

func (ctrl Controller) parameter(c *gin.Context) {
	parameter := usecase.Parameter(parameterRepository) // Dependency Injection
	c.JSON(200, parameter)
}
```

Implementation of Application Service is also the same.

## Dependency Injection

**In Goilerplate, dependencies are injected manually.**

- NOTICE: If other DI tool in Go doesn't become some kind of application framework, it will also be acceptable.

There are two ways of passing dependencies:

- with positional arguments
- with keyword arguments

### With Positional Arguments

First, define usecase with arguments of interface type.

```go
package usecase

func Parameter(r repository.IParameter) domain.Parameter { // Take Argument as Interface
	return r.Get()
}
```

Second, initialize implementation and give it to the usecase.

```go
package adapter

var (
	parameterRepository = repository.Parameter{}        // Initialize Implementation
)

func (ctrl Controller) parameter(c *gin.Context) {
	parameter := usecase.Parameter(parameterRepository) // Inject Implementation to Usecase
	c.JSON(200, parameter)
}
```

### With Keyword Arguments

First, define argument struct and usecase taking it.

```go
package usecase

// OhlcArgs are arguments of Ohlc usecase
type OhlcArgs struct {
	E service.IExchange                       // Interface
	P valueobject.Pair
	T valueobject.Timeunit
}

func Ohlc(a OhlcArgs) []valueobject.CandleStick { // Take Argument as OhlcArgs
	return a.E.Ohlc(a.P, a.T)
}
```

And then, initialize the struct with keyword arguments and give it to the usecase.

```go
package adapter

var (
	bitbank             = service.Bitbank{}      // Implementation
)

func (ctrl Controller) candlestick(c *gin.Context) {
	args := usecase.OhlcArgs{                    // Initialize Struct with Keyword Arguments
		E: bitbank,                          // Passing the implementation
		P: valueobject.BtcJpy,
		T: valueobject.OneMin,
	}
	candlestick := usecase.Ohlc(args)            // Give Arguments to Usecase
	c.JSON(200, candlestick)
}
```

### Global Injecter Variable

In manual DI, implementation initialization cost will be expensive.  
So, let's use global injecter variable in order to initialize them only once.

```go
package adapter

var (
	bitbank             = service.Bitbank{}      // Injecter Variable
	parameterRepository = repository.Parameter{}
	orderRepository     = repository.Order{}
)

func (ctrl Controller) ticker(c *gin.Context) {
	pair := valueobject.BtcJpy
	ticker := usecase.Ticker(bitbank, pair)      // DI by passing bitbank
	c.JSON(200, ticker)
}
```

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

That's it!

## Testing

```zsh
~/go/src/github.com/resotto/goilerplate (master) > go test ./internal/app/...
?       github.com/resotto/goilerplate/internal/app/adapter     [no test files]
?       github.com/resotto/goilerplate/internal/app/adapter/postgresql  [no test files]
?       github.com/resotto/goilerplate/internal/app/adapter/postgresql/model    [no test files]
?       github.com/resotto/goilerplate/internal/app/adapter/repository  [no test files]
?       github.com/resotto/goilerplate/internal/app/adapter/service     [no test files]
?       github.com/resotto/goilerplate/internal/app/application/service [no test files]
ok      github.com/resotto/goilerplate/internal/app/application/usecase 0.204s
ok      github.com/resotto/goilerplate/internal/app/domain      0.273s
?       github.com/resotto/goilerplate/internal/app/domain/factory      [no test files]
?       github.com/resotto/goilerplate/internal/app/domain/repository   [no test files]
?       github.com/resotto/goilerplate/internal/app/domain/valueobject  [no test files]
```

There are two rules:

- Name of the package where test code included is `xxx_test`.
- Place mocks on `testdata` package.

### Test Package Structure

```zsh
.
├── internal
│   └── app
│       ├── application
│       │   └── usecase
│       │       ├── ticker.go      # Usecase
│       │       └── ticker_test.go # Usecase Test
│       └── domain
│           ├── parameter.go       # Entity
│           └── parameter_test.go  # Entity Test
└── testdata
    └── exchange_mock.go           # Mock if needed
```

### Entity

Please write tests in the same directory as where the entity located.

```zsh
.
└── internal
    └── app
        └── domain
            ├── parameter.go      # Target Entity
            └── parameter_test.go # Test
```

```go
// parameter_test.go
package domain_test

import (
	"testing"

	"github.com/resotto/goilerplate/internal/app/domain"
)

func TestParameter(t *testing.T) {
	tests := []struct {
		name                       string
		funds, btc                 int
		expectedfunds, expectedbtc int
	}{
		{"more funds than btc", 1000, 0, 1000, 0},
		{"same amount", 100, 100, 100, 100},
		{"much more funds than btc", 100000, 20, 100000, 20},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			parameter := domain.Parameter{
				Funds: tt.funds,
				Btc:   tt.btc,
			}
			if parameter.Funds != tt.expectedfunds {
				t.Errorf("got %q, want %q", parameter.Funds, tt.expectedfunds)
			}
			if parameter.Btc != tt.expectedbtc {
				t.Errorf("got %q, want %q", parameter.Btc, tt.expectedbtc)
			}
		})
	}
}

```

### Usecase

Please prepare mock on `testdata` package (if needed) and write tests in the same directory as the usecase.

```zsh
.
├── internal
│   └── app
│       └── application
│           ├── service
│           │   └── exchange.go    # Application Service Interface
│           └── usecase
│               ├── ticker.go      # Target Usecase
│               └── ticker_test.go # Test
└── testdata
    └── exchange_mock.go           # Mock of Application Service Interface
```

```go
// exchange_mock.go
package testdata

import "github.com/resotto/goilerplate/internal/app/domain/valueobject"

// MExchange is mock of service.IExchange
type MExchange struct{}

// Ticker is mock implementation of service.IExchange.Ticker()
func (e MExchange) Ticker(p valueobject.Pair) valueobject.Ticker {
	return valueobject.Ticker{
		Sell:      "1000",
		Buy:       "1000",
		High:      "2000",
		Low:       "500",
		Last:      "1200",
		Vol:       "20",
		Timestamp: "1600769562",
	}
}

// Ohlc is mock implementation of service.IExchange.Ohlc()
func (e MExchange) Ohlc(p valueobject.Pair, t valueobject.Timeunit) []valueobject.CandleStick {
	cs := make([]valueobject.CandleStick, 0)
	return append(cs, valueobject.CandleStick{
		Open:      "1000",
		High:      "2000",
		Low:       "500",
		Close:     "1500",
		Volume:    "30",
		Timestamp: "1600769562",
	})
}
```

```go
// ticker_test.go
package usecase_test

import (
	"testing"

	"github.com/resotto/goilerplate/internal/app/application/usecase"
	"github.com/resotto/goilerplate/internal/app/domain/valueobject"
	"github.com/resotto/goilerplate/testdata"
)

func TestTicker(t *testing.T) {
	tests := []struct {
		name              string
		pair              valueobject.Pair
		expectedsell      string
		expectedbuy       string
		expectedhigh      string
		expectedlow       string
		expectedlast      string
		expectedvol       string
		expectedtimestamp string
	}{
		{"btcjpy", valueobject.BtcJpy, "1000", "1000", "2000", "500", "1200", "20", "1600769562"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mexchange := testdata.MExchange{} // using Mock
			result := usecase.Ticker(mexchange, tt.pair)
			if result.Sell != tt.expectedsell {
				t.Errorf("got %q, want %q", result.Sell, tt.expectedsell)
			}
			if result.Buy != tt.expectedbuy {
				t.Errorf("got %q, want %q", result.Buy, tt.expectedbuy)
			}
			if result.High != tt.expectedhigh {
				t.Errorf("got %q, want %q", result.High, tt.expectedhigh)
			}
			if result.Low != tt.expectedlow {
				t.Errorf("got %q, want %q", result.Low, tt.expectedlow)
			}
			if result.Last != tt.expectedlast {
				t.Errorf("got %q, want %q", result.Last, tt.expectedlast)
			}
			if result.Vol != tt.expectedvol {
				t.Errorf("got %q, want %q", result.Vol, tt.expectedvol)
			}
			if result.Timestamp != tt.expectedtimestamp {
				t.Errorf("got %q, want %q", result.Timestamp, tt.expectedtimestamp)
			}
		})
	}
}
```

## Naming Convention

### Interface

- Add prefix `I` like `IExchange`.
  - NOTICE: If you can distinguish interface from implementation, any naming convention will be acceptable.

### Mock

- Add prefix `M` like `MExchange`.
  - NOTICE: If you can distinguish mock from production, any naming convention will be acceptable.

### File

- File names can be duplicated.
- For test, add suffix `_test` like `parameter_test.go`.
- For mock, add suffix `_mock` like `exchange_mock.go`.

### Package

- For package name, please check following posts:

  - [Package names](https://blog.golang.org/package-names)
  - [Names](https://golang.org/doc/effective_go.html#names)

- For package layout, please check:
  - [Project Layout](https://github.com/golang-standards/project-layout)

## With Gochk

[Gochk, static dependency analysis tool for go files,](https://github.com/resotto/gochk) empowers Goilerplate so much!

**[Gochk](https://github.com/resotto/gochk) confirms that codebase follows [Clean Architecture The Dependency Rule](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html#the-dependency-rule).**

Let's merge Gochk into CI process.

```yml
name: test

on:
  push:
    branches:
      - master
    paths-ignore:
      - "**/*.md"
  pull_request:
    branches:
      - master

jobs:
  gochk-goilerplate:
    runs-on: ubuntu-latest
    container:
      image: docker://ghcr.io/resotto/gochk:latest
    steps:
      - name: Clone Goilerplate
        uses: actions/checkout@v2
        with:
          repository: {{ github.repository }}
      - name: Run Gochk
        run: |
          /go/bin/gochk -c=/go/src/github.com/resotto/gochk/configs/config.json
```

And then, [its result is](https://github.com/resotto/goilerplate/runs/1367461573):

![Gochk Result in GitHub Actions](https://user-images.githubusercontent.com/19743841/98438959-6f56b680-2131-11eb-8b6e-d835e56239e0.png)

## With PostgreSQL

First, you pull the docker image `ghcr.io/resotto/goilerplate-pg` from GitHub Container Registry and run container with following command:

```zsh
docker run -d -it --name pg -p 5432:5432 -e POSTGRES_PASSWORD=postgres ghcr.io/resotto/goilerplate-pg:latest
```

Then, let's check it out:

```zsh
open http://0.0.0.0:8080/parameter
open http://0.0.0.0:8080/order
```

### Building Image

If you fail pulling image from GitHub Container Registry, you also can build Docker image from Dockerfile.

```zsh
cd build
docker build -t goilerplate-pg:latest .
docker run -d -it --name pg -p 5432:5432 -e POSTGRES_PASSWORD=postgres goilerplate-pg:latest
```

### Docker Image

The image you pulled from GitHub Container Registry is built from the simple Dockerfile and init.sql.

```Dockerfile
FROM postgres

EXPOSE 5432

COPY ./init.sql /docker-entrypoint-initdb.d/
```

```sql
create table parameters (
    id integer primary key,
    funds integer,
    btc integer
);

insert into parameters values (1, 10000, 10);

create table persons (
    person_id uuid primary key,
    name text not null,
    weight integer
);

create table card_brands (
    brand text primary key
);

create table cards (
    card_id uuid primary key,
    brand text references card_brands(brand) on update cascade
);

create table orders (
    order_id uuid primary key,
    person_id uuid references persons(person_id)
);

create table payments (
    order_id uuid primary key references orders(order_id),
    card_id uuid references cards(card_id)
);

insert into persons values ('f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b', 'Jerry', 1);

insert into card_brands values ('VISA'), ('AMEX');

insert into cards values ('3224ebc0-0a6e-4e22-9ce8-c6564a1bb6a1', 'VISA');

insert into orders values ('722b694c-984c-4208-bddd-796553cf83e1', 'f3bf75a9-ea4c-4f57-9161-cfa8f96e2d0b');

insert into payments values ('722b694c-984c-4208-bddd-796553cf83e1', '3224ebc0-0a6e-4e22-9ce8-c6564a1bb6a1');
```

## Feedbacks

[Feel free to write your thoughts](https://github.com/resotto/goilerplate/issues/1)

## License

[GNU General Public License v3.0](https://github.com/resotto/goilerplate/blob/master/LICENSE).

## Author

Resotto

<a href="https://github.com/resotto"><img src="https://user-images.githubusercontent.com/19743841/97778118-4629a980-1bb8-11eb-97ed-76dcdbe50406.png" /></a>
<a href="https://twitter.com/resotto3"><img src="https://user-images.githubusercontent.com/19743841/97777698-52f8ce00-1bb5-11eb-93c9-b06e0c48b693.png" /></a>
<a href="https://github.com/sponsors/resotto"><img src="https://img.shields.io/badge/Sponsor-ffffff?logo=github&logoColor=pink&style=flat-square" /></a>
