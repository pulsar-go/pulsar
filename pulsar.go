package pulsar

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/kabukky/httpscerts"
	"github.com/pulsar-go/pulsar/config"
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/router"
)

// fileExists determines if a file exists in a given path.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// debugHandler is responsible for each http handler in debug mode.
func developmentHandler(route *router.Route) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("[PULSAR] Request %s \n", r.URL)
		res := route.Handler(&request.HTTP{Req: r, Params: ps})
		res.Handle(w)
	}
}

// productionHandler is responsible for each http handler in debug mode.
func productionHandler(route *router.Route) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := route.Handler(&request.HTTP{Req: r, Params: ps})
		res.Handle(w)
	}
}

// RegisterRoutes registers the routes.
func RegisterRoutes(mux *httprouter.Router, r *router.Router) {
	// Register the routes.
	var handler func(*router.Route) func(http.ResponseWriter, *http.Request, httprouter.Params)
	if config.Settings.Server.Development {
		handler = developmentHandler
	} else {
		handler = productionHandler
	}
	for _, element := range r.Routes {
		route := element
		switch route.Method {
		case request.GetRequest:
			mux.GET(route.URI, handler(&route))
		case request.HeadRequest:
			mux.HEAD(route.URI, handler(&route))
		case request.PostRequest:
			mux.POST(route.URI, handler(&route))
		case request.PutRequest:
			mux.PUT(route.URI, handler(&route))
		case request.PatchRequest:
			mux.PATCH(route.URI, handler(&route))
		case request.DeleteRequest:
			mux.DELETE(route.URI, handler(&route))
		}
	}
	// Register his childs.
	for _, element := range r.Childs {
		RegisterRoutes(mux, element)
	}
}

// Serve starts the server.
func Serve(router *router.Router) error {
	mux := httprouter.New()
	// Register the application routes.
	RegisterRoutes(mux, router)
	// Set the address of the server.
	address := config.Settings.Server.Host + ":" + config.Settings.Server.Port
	// Generate a SSL certificate if needed.
	if config.Settings.HTTPS.Enabled {
		err := httpscerts.Check(config.Settings.HTTPS.CertFile, config.Settings.HTTPS.KeyFile)
		// If they are not available, generate new ones.
		if err != nil {
			err = httpscerts.Generate(config.Settings.HTTPS.CertFile, config.Settings.HTTPS.KeyFile, address)
			if err != nil {
				log.Fatal("Unable to create HTTP certificates.")
			}
		}
	}
	if config.Settings.Server.Development {
		fmt.Println("-----------------------------------------------------")
		fmt.Println("|                                                   |")
		fmt.Println("|  P U L S A R                                      |")
		fmt.Println("|  Go Web Micro-framework                           |")
		fmt.Println("|                                                   |")
		fmt.Println("|  Erik Campobadal <soc@erik.cat>                   |")
		fmt.Println("|  Krishan König <krishan.koenig@googlemail.com>    |")
		fmt.Println("|                                                   |")
		fmt.Println("-----------------------------------------------------")
		fmt.Println()
	}
	if config.Settings.HTTPS.Enabled {
		if config.Settings.Server.Development {
			fmt.Printf("Creating a HTTP/2 server with TLS on %s\n", address)
			fmt.Printf("Certificate: %s\nKey: %s\n\n", config.Settings.HTTPS.CertFile, config.Settings.HTTPS.KeyFile)
		}
		return http.ListenAndServeTLS(address, config.Settings.HTTPS.CertFile, config.Settings.HTTPS.KeyFile, mux)
	}
	if config.Settings.Server.Development {
		fmt.Printf("Creating a HTTP/1.1 server on %s\n\n", address)
	}
	return http.ListenAndServe(address, mux)
}
