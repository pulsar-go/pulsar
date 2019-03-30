package pulsar

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
	"github.com/kabukky/httpscerts"
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
	"github.com/pulsar-go/pulsar/router"
)

// Settings represents the pulsar server settings structure.
type Settings struct {
	Server struct {
		Host        string `toml:"host"`
		Port        string `toml:"port"`
		Development bool   `toml:"development"`
	} `toml:"server"`
	HTTPS struct {
		Enabled  bool   `toml:"enabled"`
		CertFile string `toml:"cert_file"`
		KeyFile  string `toml:"key_file"`
	} `toml:"https"`
}

// fileExists determines if a file exists in a given path.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// GetConfig gets the configuration from a configuration file.
func GetConfig(path string) *Settings {
	settings := new(Settings)
	// Open the server configuration file.
	absPath, _ := filepath.Abs(filepath.Clean(path))
	if _, err := toml.DecodeFile(absPath, &settings); err != nil {
		log.Fatalln("There was an error decoding file " + absPath)
	}
	// Transform the relative paths into absolute.
	settings.HTTPS.CertFile, _ = filepath.Abs(filepath.Dir(path) + "/" + filepath.Clean(settings.HTTPS.CertFile))
	settings.HTTPS.KeyFile, _ = filepath.Abs(filepath.Dir(path) + "/" + filepath.Clean(settings.HTTPS.KeyFile))
	// Create and return the settings.
	return settings
}

// debugHandler is responsible for each http handler in debug mode.
func developmentHandler(route *router.Route) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.Printf("[PULSAR] Request %s \n", r.URL)
		res := route.Handler(&request.HTTP{Req: r, Params: ps})
		switch res.Type {
		case response.TextResponse:
			fmt.Fprint(w, res.TextData)
		case response.JSONResponse:
			w.Header().Set("Content-Type", "application/json")
			result, err := json.Marshal(res.JSONData)
			if err != nil {
				fmt.Fprint(w, "Error while marshaling JSON.")
			}
			fmt.Fprint(w, string(result))
		default:
			fmt.Fprint(w, "Invalid HTTP response type.")
		}
	}
}

// productionHandler is responsible for each http handler in debug mode.
func productionHandler(route *router.Route) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		res := route.Handler(&request.HTTP{Req: r, Params: ps})
		switch res.Type {
		case response.TextResponse:
			fmt.Fprint(w, res.TextData)
		case response.JSONResponse:
			w.Header().Set("Content-Type", "application/json")
			result, err := json.Marshal(res.JSONData)
			if err != nil {
				fmt.Fprint(w, "Error while marshaling JSON.")
			}
			fmt.Fprint(w, string(result))
		default:
			fmt.Fprint(w, "Invalid HTTP response type.")
		}
	}
}

// RegisterRoutes registers the routes.
func RegisterRoutes(settings *Settings, mux *httprouter.Router, r *router.Router) {
	// Register the routes.
	var handler func(*router.Route) func(http.ResponseWriter, *http.Request, httprouter.Params)
	if settings.Server.Development {
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
		RegisterRoutes(settings, mux, element)
	}
}

// Serve starts the server.
func Serve(router *router.Router, settings *Settings) error {
	mux := httprouter.New()
	// Register the application routes.
	RegisterRoutes(settings, mux, router)
	// Set the address of the server.
	address := settings.Server.Host + ":" + settings.Server.Port
	// Generate a SSL certificate if needed.
	if settings.HTTPS.Enabled {
		err := httpscerts.Check(settings.HTTPS.CertFile, settings.HTTPS.KeyFile)
		// If they are not available, generate new ones.
		if err != nil {
			err = httpscerts.Generate(settings.HTTPS.CertFile, settings.HTTPS.KeyFile, address)
			if err != nil {
				log.Fatal("Unable to create HTTP certificates.")
			}
		}
	}
	if settings.Server.Development {
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
	if settings.HTTPS.Enabled {
		if settings.Server.Development {
			fmt.Printf("Creating a HTTP/2 server with TLS on %s\n", address)
			fmt.Printf("Certificate: %s\nKey: %s\n\n", settings.HTTPS.CertFile, settings.HTTPS.KeyFile)
		}
		return http.ListenAndServeTLS(address, settings.HTTPS.CertFile, settings.HTTPS.KeyFile, mux)
	}
	if settings.Server.Development {
		fmt.Printf("Creating a HTTP/1.1 server on %s\n\n", address)
	}
	return http.ListenAndServe(address, mux)
}
