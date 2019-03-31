package router

import (
	"github.com/pulsar-go/pulsar/request"
	"github.com/pulsar-go/pulsar/response"
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
	return &Router{options: *options}
}

// Adds the route to the given router.
func addRoute(r *Router, uri string, handler Handler, method request.Type) {
	h := handler
	if r.options.Middleware != nil {
		h = r.options.Middleware(handler)
	}
	// Append the route to the list.
	r.Routes = append(
		r.Routes,
		Route{URI: r.options.Prefix + uri, Method: method, Handler: h},
	)
}

// Get creates a GET route.
func (r *Router) Get(uri string, handler Handler) {
	addRoute(r, uri, handler, request.GetRequest)
}

// Head creates a HEAD route.
func (r *Router) Head(uri string, handler Handler) {
	addRoute(r, uri, handler, request.HeadRequest)
}

// Post creates a POST route.
func (r *Router) Post(uri string, handler Handler) {
	addRoute(r, uri, handler, request.PostRequest)
}

// Put creates a PUT route.
func (r *Router) Put(uri string, handler Handler) {
	addRoute(r, uri, handler, request.PutRequest)
}

// Patch creates a PATCH route.
func (r *Router) Patch(uri string, handler Handler) {
	addRoute(r, uri, handler, request.PatchRequest)
}

// Delete creates a DELETE route.
func (r *Router) Delete(uri string, handler Handler) {
	addRoute(r, uri, handler, request.DeleteRequest)
}

// Group certain routes uner certain options.
func (r *Router) Group(options *Options, routes func(r *Router)) {
	router := CreateWithOptions(options)
	routes(router)
	r.Childs = append(r.Childs, router)
}
