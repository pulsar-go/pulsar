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
)

// Type is the name of the response type.
type Type uint

// Indicate the available response types.
const (
	TextResponse Type = iota
	JSONResponse
	StaticResponse
	ViewResponse
)

// HTTP is the web server response.
type HTTP struct {
	Type     Type
	TextData string
	JSONData interface{}
}

// Text returns a HTTP response with plain text.
func Text(text string) HTTP {
	return HTTP{Type: TextResponse, TextData: text}
}

// JSON returns a HTTP response with the JSON headers.
func JSON(data interface{}) HTTP {
	return HTTP{Type: JSONResponse, JSONData: data}
}

// Static return a View response without templating data.
func Static(name string) HTTP {
	path, err := filepath.Abs(filepath.Clean(config.Settings.Views.Path) + "/" + filepath.Clean(name))
	if err != nil {
		log.Fatalln(err)
	}
	return HTTP{Type: StaticResponse, TextData: path}
}

// View return a View response with templating data.
func View(name string, data interface{}) HTTP {
	path, err := filepath.Abs(filepath.Clean(config.Settings.Views.Path) + "/" + filepath.Clean(name))
	if err != nil {
		log.Fatalln(err)
	}
	return HTTP{Type: ViewResponse, TextData: path, JSONData: data}
}

// Handle handles the HTTP request using a response writter.
func (response *HTTP) Handle(writer http.ResponseWriter) {
	switch response.Type {
	case TextResponse:
		fmt.Fprint(writer, response.TextData)
	case JSONResponse:
		writer.Header().Set("Content-Type", "application/json")
		result, err := json.Marshal(response.JSONData)
		if err != nil {
			fmt.Fprint(writer, "Error while marshaling JSON.")
		}
		fmt.Fprint(writer, string(result))
	case StaticResponse:
		content, err := ioutil.ReadFile(response.TextData)
		if err != nil {
			log.Println("File " + response.TextData + " not found.")
		}
		fmt.Fprintf(writer, string(content))
	case ViewResponse:
		template.Must(template.ParseFiles(response.TextData)).Execute(writer, response.JSONData)
	default:
		fmt.Fprint(writer, "Invalid HTTP response type.")
	}
}
