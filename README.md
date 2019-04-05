# Pulsar - A Go Web framework for web artisans

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
- Database Configuration + ORM
- Emails

## How to use

You first need to get the package:

```
go get github.com/pulsar-go/pulsar
```

**IMPORTANT**: Check the official example here: <https://github.com/pulsar-go/example>

Then you'll need to create some server configuration (`server.toml` for example):

```toml
# Server stores all the settings releated
# to the HTTP / HTTP server itself.
[server]
    # Host determines the IP or name of the
    # HTTP server.
    host = ""
    # Port specifies the port that is going to
    # be used to run the HTTP server.
    port = "8080"
    # Development specifies if the server is going
    # to run in development mode, allowing some output
    # in the terminal while the application runs giving
    # help and insight of what's going on.
    development = true

# HTTPS stores all the settings releated
# to the TLS (SSL) settings used to ensure
# encrypted connections.
[https]
    # Enables or disables the HTTPs server.
    # Please notice that this should always be
    # enabled. HTTPs should always be used.
    # Only a HTTP or HTTPs server can be fired
    # so if enabled, no HTTP server will be provided.
    enabled = true
    # Certificate file that the HTTPs
    # server will use to encrypt connections.
    # Auto generated if no cert_file and key_file
    # exists in the path provided.
    cert_file = "./server.cert"
    # Key file that the HTTPs
    # server will use to encrypt connections.
    # Auto generated if no cert_file and key_file
    # exists in the path provided.
    key_file = "./server.key"

# Views stores all the settings releated
# to the views stored in the application
# used when delivering responses from
# the web server.
[views]
    # Path represents the relative or absolute
    # path where the views will come from.
    # It acts as a path prefix when returning views
    path = "./views"

# Database stores all the settings releated
# to the database connection that the ORM
# will use in order to provide it's functionality
[database]
    # Determines the driver to use
    # Possible options are:
    # 'mysql', 'postgres', 'sqlite3'.
    driver = "mysql"
    # Database represents the database name
    # to use or the database path in case of
    # using sqlite3 driver. If driver is sqlire3
    # the value ':memory:' can also be used to create
    # a temp database stored in memory.
    database = "sample"
    # Host represents the database host where
    # it will connect to. Unused if using
    # the sqlite3 driver.
    host = "localhost"
    # Port represents the database port where
    # it will connect to. Unused if using
    # the sqlite3 driver.
    port = "3306"
    # User represents the user used to establish
    # the database connection. Unused if using
    # the sqlite3 driver.
    user = "sample"
    # Password represents the password used to establish
    # the database connection. Unused if using
    # the sqlite3 driver.
    password = "secret"
    # Auto migrate is used to migrate the database shema
    # using the provided models. It  will ONLY create tables,
    # missing columns and missing indexes, and WON’T change
    # existing column’s type or delete unused columns
    # to protect your data.
    auto_migrate = true

# Mail stores all the information about
# SMTP mailing to send any form of email.
[mail]
    # Host determines the SMTP host that
    # is going to be used.
    host = "smtp.mailtrap.io"
    # Port determines the SMTP port that
    # is going to be used.
    port = "465"
    # Identity determines how the auth is
    # pretended to act as. Usually this
    # should be an empty string.
    identity = ""
    # Username used to authenticate when connecting
    # to the host. Part of the credentials.
    username = ""
    # Password used to authenticate when connecting
    # to the host. Part of the credentials.
    password = ""
    # From determines who the mail is going to be sent
    # from. This setting is the default from address used.
    from = "mail@example.com"
```

Then create a main file (`server.go` for example):
```go
package main

import (
	"log"

    "github.com/pulsar-go/pulsar"
    "github.com/pulsar-go/pulsar/config"
    "github.com/pulsar-go/pulsar/router"
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
)

type sample struct {
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
	return response.JSON(sample{Name: "Erik", Age: 22})
}

func sampleMiddleware(next router.Handler) router.Handler {
	return router.Handler(func(req *request.HTTP) response.HTTP {
		log.Println("Before route middleware")
		r := next(req)
		log.Println("After route middleware")
		return r
	})
}

func main() {
	// Get the settings from the configuration files.
	config.Set("./server.toml")
	// Set the application routes.
	router.Routes.
        Get("/", index).
        Get("/user/:id", user).
        Group(&router.Options{Prefix: "/sample", Middleware: sampleMiddleware}, func(routes *router.Router) {
            routes.Get("/about", about)
        })
	// Serve the HTTP server.
	log.Fatalln(pulsar.Serve())
}
```

## Documentation

- Pulsar: <https://godoc.org/github.com/pulsar-go/pulsar>
- Config: <https://godoc.org/github.com/pulsar-go/pulsar/config>
- Router: <https://godoc.org/github.com/pulsar-go/pulsar/router>
- Request: <https://godoc.org/github.com/pulsar-go/pulsar/request>
- Response: <https://godoc.org/github.com/pulsar-go/pulsar/response>
