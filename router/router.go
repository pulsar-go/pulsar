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
func addRoute(router *Router, uri string, handler Handler, method request.Type) *Router {
	h := handler
	if router.options.Middleware != nil {
		h = router.options.Middleware(handler)
	}
	// Append the route to the list.
	router.Routes = append(
		router.Routes,
		Route{URI: router.options.Prefix + uri, Method: method, Handler: h},
	)
	return router
}

// Get creates a GET route.
func (router *Router) Get(uri string, handler Handler) *Router {
	return addRoute(router, uri, handler, request.GetRequest)
}

// Head creates a HEAD route.
func (router *Router) Head(uri string, handler Handler) *Router {
	return addRoute(router, uri, handler, request.HeadRequest)
}

// Post creates a POST route.
func (router *Router) Post(uri string, handler Handler) *Router {
	return addRoute(router, uri, handler, request.PostRequest)
}

// Put creates a PUT route.
func (router *Router) Put(uri string, handler Handler) *Router {
	return addRoute(router, uri, handler, request.PutRequest)
}

// Patch creates a PATCH route.
func (router *Router) Patch(uri string, handler Handler) *Router {
	return addRoute(router, uri, handler, request.PatchRequest)
}

// Delete creates a DELETE route.
func (router *Router) Delete(uri string, handler Handler) *Router {
	return addRoute(router, uri, handler, request.DeleteRequest)
}

// Group certain routes uner certain options.
func (router *Router) Group(options *Options, routes func(router *Router)) *Router {
	routerGroup := &Router{options: *options}
	routes(routerGroup)
	router.Childs = append(router.Childs, routerGroup)
	return router
}
