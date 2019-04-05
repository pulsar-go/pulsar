package request

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Type is the name of the response type.
type Type uint

// Indicate the available response types.
const (
	GetRequest Type = iota
	HeadRequest
	PostRequest
	PutRequest
	PatchRequest
	DeleteRequest
)

// HTTP represents the web server request.
type HTTP struct {
	Request *http.Request
	Params  httprouter.Params
}
