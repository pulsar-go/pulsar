package response

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/pulsar-go/pulsar/config"
	"github.com/pulsar-go/pulsar/request"
)

// Type is the name of the response type.
type Type uint

// Indicate the available response types.
const (
	TextResponse Type = iota
	JSONResponse
	StaticResponse
	AssetResponse
	ViewResponse
)

// HTTP is the web server response.
type HTTP struct {
	StatusCode int
	Type       Type
	TextData   string
	JSONData   interface{}
}

// Text returns a HTTP response with plain text.
func Text(text string) HTTP {
	return HTTP{StatusCode: http.StatusOK, Type: TextResponse, TextData: text}
}

// TextWithCode is a Text response with additional status code.
func TextWithCode(text string, code int) HTTP {
	res := Text(text)
	res.StatusCode = code
	return res
}

// JSON returns a HTTP response with the JSON headers.
func JSON(data interface{}) HTTP {
	return HTTP{StatusCode: http.StatusOK, Type: JSONResponse, JSONData: data}
}

// JSONWithCode is a JSON response with additional status code.
func JSONWithCode(data interface{}, code int) HTTP {
	res := JSON(data)
	res.StatusCode = code
	return res
}

// Static return a View response without templating data.
func Static(name string) HTTP {
	path, err := filepath.Abs(filepath.Clean(config.Settings.Views.Path) + "/" + filepath.Clean(name+".html"))
	if err != nil {
		log.Println(err)
	}
	return HTTP{StatusCode: http.StatusOK, Type: StaticResponse, TextData: path}
}

// StaticWithCode is a Static response with additional status code.
func StaticWithCode(name string, code int) HTTP {
	res := Static(name)
	res.StatusCode = code
	return res
}

// Asset return an asset response (css files, js files, images, etc.).
func Asset(name string) HTTP {
	path, err := filepath.Abs(filepath.Clean(config.Settings.Views.Path) + "/" + filepath.Clean(name))
	if err != nil {
		log.Println(err)
	}
	return HTTP{StatusCode: http.StatusOK, Type: AssetResponse, TextData: path}
}

// View return a View response with templating data.
func View(name string, data interface{}) HTTP {
	path, err := filepath.Abs(filepath.Clean(config.Settings.Views.Path) + "/" + filepath.Clean(name+".gohtml"))
	if err != nil {
		log.Println(err)
	}
	return HTTP{StatusCode: http.StatusOK, Type: ViewResponse, TextData: path, JSONData: data}
}

// ViewWithCode is a View response with additional code.
func ViewWithCode(name string, data interface{}, code int) HTTP {
	res := View(name, data)
	res.StatusCode = code
	return res
}

// Handle handles the HTTP request using a response writter.
func (response *HTTP) Handle(req *request.HTTP) {
	writer := req.Writer
	switch response.Type {
	case TextResponse:
		writer.WriteHeader(response.StatusCode)
		fmt.Fprint(writer, response.TextData)
	case JSONResponse:
		writer.Header().Set("Content-Type", "application/json")
		result, err := json.Marshal(response.JSONData)
		if err != nil {
			fmt.Fprint(writer, "Error while marshaling JSON.")
		}
		writer.WriteHeader(response.StatusCode)
		fmt.Fprint(writer, string(result))
	case StaticResponse, AssetResponse:
		content, err := ioutil.ReadFile(response.TextData)
		if err != nil {
			log.Println("File " + response.TextData + " not found.")
		}
		writer.WriteHeader(response.StatusCode)
		fmt.Fprintf(writer, string(content))
	case ViewResponse:
		writer.WriteHeader(response.StatusCode)
		template.Must(template.ParseFiles(response.TextData)).Execute(writer, response.JSONData)
	default:
		writer.WriteHeader(response.StatusCode)
		fmt.Fprint(writer, "Invalid HTTP response type.")
	}
}
