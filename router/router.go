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

// Routes representrs the global application routes.
var Routes Router

// Adds the route to the given router.
func addRoute(r *Router, uri string, handler Handler, method request.Type) *Router {
	h := handler
	if r.options.Middleware != nil {
		h = r.options.Middleware(handler)
	}
	// Append the route to the list.
	r.Routes = append(
		r.Routes,
		Route{URI: r.options.Prefix + uri, Method: method, Handler: h},
	)
	return r
}

// Get creates a GET route.
func (r *Router) Get(uri string, handler Handler) *Router {
	return addRoute(r, uri, handler, request.GetRequest)
}

// Head creates a HEAD route.
func (r *Router) Head(uri string, handler Handler) *Router {
	return addRoute(r, uri, handler, request.HeadRequest)
}

// Post creates a POST route.
func (r *Router) Post(uri string, handler Handler) *Router {
	return addRoute(r, uri, handler, request.PostRequest)
}

// Put creates a PUT route.
func (r *Router) Put(uri string, handler Handler) *Router {
	return addRoute(r, uri, handler, request.PutRequest)
}

// Patch creates a PATCH route.
func (r *Router) Patch(uri string, handler Handler) *Router {
	return addRoute(r, uri, handler, request.PatchRequest)
}

// Delete creates a DELETE route.
func (r *Router) Delete(uri string, handler Handler) *Router {
	return addRoute(r, uri, handler, request.DeleteRequest)
}

// Group certain routes uner certain options.
func (r *Router) Group(options *Options, routes func(r *Router)) *Router {
	router := &Router{options: *options}
	routes(router)
	r.Childs = append(r.Childs, router)
	return r
}
