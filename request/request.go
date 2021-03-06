package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ErrorNoJSONHeader determines that the current request have no JSON headers.
var ErrorNoJSONHeader = errors.New("a")

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
	Request     *http.Request
	Body        string
	Writer      http.ResponseWriter
	Params      httprouter.Params
	Additionals map[string]interface{}
}

// JSON transforms the input body that's formatted in
func (req *HTTP) JSON(data interface{}) error {
	return json.NewDecoder(bytes.NewBufferString(req.Body)).Decode(data)
}
