<h1 align="center">Goilerplate</h1>

<p align="center">
  <a href="https://github.com/resotto/goilerplate/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-GPL%20v3.0-brightgreen.svg" /></a>
</p>

<p align="center">
  Clean Boilerplate of Go + Domain-Driven Design + Clean Architecture + Gin + GORM.  
</p>

<p align="center">
  <img src="https://user-images.githubusercontent.com/19743841/93784289-cbd84200-fc67-11ea-997b-b99af8affe17.png">
</p>

---

Note:

- Default application code is trivial because you will implement cool logic.
- Public API of bitbank, which is bitcoin exchange located in Tokyo, is used for some endpoints by default.

Requirements:

- [Go](https://golang.org/doc/install)

---

## Getting Started

```zsh
go get github.com/resotto/goilerplate
cd ${GOPATH}/src/github.com/resotto/goilerplate
go run main.go # please run main.go from root directory
open http://0.0.0.0:8080
```

## Can't `go get` this package via SSH ?

`go get` GitHub repository via HTTPS by default.  
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
    r.LoadHTMLGlob("cmd/app/adapter/view/*")
    ```
- With Public API of bitbank
  - `GET /ticker`
  - `GET /candlestick`
    - NOTICE: This works from 0AM ~ 3PM (UTC) due to its API constraints.
- With PostgreSQL
  - `GET /parameter`
    - [NOTICE: Please run postgres container first with this step.](#with-postgresql)

## Package Structure

```
.
├── cmd
│   └── app
│       ├── adapter
│       │   ├── controller.go
│       │   ├── postgresql
│       │   │   ├── conn.go
│       │   │   └── model
│       │   │       └── parameter.go
│       │   ├── repository
│       │   │   └── parameter.go
│       │   ├── service
│       │   │   └── bitbank.go
│       │   └── view
│       │       └── index.tmpl
│       ├── application
│       │   ├── service
│       │   │   └── exchange.go
│       │   └── usecase
│       │       ├── ohlc.go
│       │       ├── parameter.go
│       │       └── ticker.go
│       └── domain
│           ├── parameter.go
│           ├── repository
│           │   └── parameter.go
│           └── valueobject
│               ├── candlestick.go
│               ├── pair.go
│               ├── ticker.go
│               └── timeunit.go
└── main.go
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

## How to cross the border of those layers

Please follow next four tasks:

1. Define Interface
1. Take Argument as Interface and Call Functions of It
1. Implement It
1. Inject Dependency

```
.
├── adapter
│   ├── controller.go    // 4. Dependency Injection
│   └── repository
│       └── parameter.go // 3. Implementation
├── application
│   └── usecase
│       └── parameter.go // 2. Interface Function Call
└── domain
    ├── parameter.go
    └── repository
        └── parameter.go // 1. Interface
```

Here, I pick up example of Repository whose import statement is omitted.

### Repository

1. Interface at Domain Layer:

```go
package repository

// IParameter is interface of parameter repository
type IParameter interface {
	Get() domain.Parameter
	Save(domain.Parameter)
}
```

2. Usecase at Application Layer:

```go
package usecase

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

// Save saves parameter
func (r Parameter) Save(p domain.Parameter) {
	// TODO
}
```

4. Dependency Injection at Controller of Adapter Layer:

```go
package adapter

func (ctrl Controller) parameter(c *gin.Context) {
	repository := repository.Parameter{}
	parameter := usecase.Parameter(repository) // Dependency Injection
	c.JSON(200, parameter)
}
```

Implementation of Application Service is also the same.

## Naming Convention

### Interface

- Add prefix `I` like `IParameter`.
  - NOTICE: If you can distinguish interface from implementation, any naming convention will be acceptable.

### File

- File names can be duplicated.

### Package

- For package, please check following posts:
  - [Package names](https://blog.golang.org/package-names)
  - [Names](https://golang.org/doc/effective_go.html#names)

## With PostgreSQL

First, you pull docker image from GitHub Container Registry and run container with following command:

```bash
docker run -d -it --name pg -p 5432:5432 -e POSTGRES_PASSWORD=postgres ghcr.io/resotto/goilerplate-pg:latest
```

Then, let's check it out:

```bash
open http://0.0.0.0:8080/parameter
```

### Docker Image

The image you pulled from GitHub Container Registry is built from simple Dockerfile and init.sql.

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
```

## Feedbacks

[Feel free to write your thoughts](https://github.com/resotto/goilerplate/issues/1)

## License

[GNU General Public License v3.0](https://github.com/resotto/goilerplate/blob/master/LICENSE).
