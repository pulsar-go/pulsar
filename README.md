# Cervol - A Go Web Micro-framework for web artisans

This is a go web microframework for web artisans willing to create amazing go web applications with ease.

## Features

- Server configuration in a single file
- A blazing fast router ([github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter))
- Simplified request / response API
- MVC architecture
- Middlewares
- Automatic headers for different responses
- Automatic TLS (SSL) certificate using openssl cli
- Automatic server creation using HTTP/1.1 or HTTP/2

## How to use

You first need to get the package:

```
go get github.com/ConsoleTVs/cervol
```

Then you'll need to create some server configuration (`server.toml` for example):

```toml
[server]
    host = ""
    port = "8080"
    development = true

[https]
    enabled = true
    auto_generate_certificate = true
    cert_file = "./server.cert"
    key_file = "./server.key"
```

Then create a main file (`server.go` for example):
```go
package main

import (
    "log"

    "github.com/ConsoleTVs/cervol"
    "github.com/ConsoleTVs/cervol/router"
    "github.com/ConsoleTVs/cervol/request"
    "github.com/ConsoleTVs/cervol/response"
)

type Sample struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func index(req *request.HTTP) response.HTTP {
	return response.Text("Sample index response")
}

func user(req *request.HTTP) response.HTTP {
	return response.Text("User id: " + req.Params.ByName("id"))
}

func about(req *request.HTTP) response.HTTP {
	return response.JSON(Sample{Name: "Erik", Age: 22})
}

func middle(next router.Handler) router.Handler {
	return router.Handler(func(req *request.HTTP) response.HTTP {
		log.Println("Before route middleware")
		r := next(req)
		log.Println("After route middleware")
		return r
	})
}

func main() {
	// Get the settings from the configuration files.
	settings := server.GetConfig("./server.toml")
	// Set the application routes.
	routes := router.Create()
	routes.Get("/", index)
	routes.Get("/user/:id", user)
	routes.Group(&router.Options{Prefix: "/sample", Middleware: middle}, func(routes *router.Router) {
		routes.Get("/about", about)
	})
	// Serve the HTTP server.
	log.Fatalln(server.Serve(routes, settings))
}
```
