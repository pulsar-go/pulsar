package router

import (
	"github.com/ConsoleTVs/cervol/request"
	"github.com/ConsoleTVs/cervol/response"
)

// Handler represents a route handler.
type Handler func(req *request.HTTP) response.HTTP

// Middleware represents a route middleware.
type Middleware func(next Handler) Handler

// Options represents the route options.
type Options struct {
	Prefix     string
	Middleware Middleware
}

// Route is the definition of a route.
type Route struct {
	URI     string
	Method  request.Type
	Handler Handler
}

// Router determines how a route is.
type Router struct {
	Routes  []Route
	options Options
	Childs  []*Router
}

// Create creates a new router (routes).
func Create() *Router {
	r := new(Router)
	r.options = Options{}
	return r
}

// CreateWithOptions creates a new router (routes) given the options.
func CreateWithOptions(options *Options) *Router {
	r := new(Router)
	r.options = *options
	return r
}

// Get creates a GET route.
func (r *Router) Get(uri string, handler func(req *request.HTTP) response.HTTP) {
	h := handler
	if r.options.Middleware != nil {
		h = r.options.Middleware(handler)
	}
	// Append the route to the list.
	r.Routes = append(
		r.Routes,
		Route{URI: r.options.Prefix + uri, Handler: h},
	)
}

// Group certain routes uner certain options.
func (r *Router) Group(options *Options, routes func(r *Router)) {
	router := CreateWithOptions(options)
	routes(router)
	r.Childs = append(r.Childs, router)
}
