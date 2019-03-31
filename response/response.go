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
	ViewResponse
	TemplateResponse
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

// View return a View response without templating data.
func View(name string) HTTP {
	path, err := filepath.Abs(filepath.Clean(config.Settings.Views.Path) + "/" + filepath.Clean(name))
	if err != nil {
		log.Fatalln(err)
	}
	return HTTP{Type: ViewResponse, TextData: path}
}

// Template return a View response with templating data.
func Template(name string, data interface{}) HTTP {
	path, err := filepath.Abs(filepath.Clean(config.Settings.Views.Path) + "/" + filepath.Clean(name))
	if err != nil {
		log.Fatalln(err)
	}
	return HTTP{Type: TemplateResponse, TextData: path, JSONData: data}
}

// Handle handles the HTTP request using a response writter.
func (res *HTTP) Handle(w http.ResponseWriter) {
	switch res.Type {
	case TextResponse:
		fmt.Fprint(w, res.TextData)
	case JSONResponse:
		w.Header().Set("Content-Type", "application/json")
		result, err := json.Marshal(res.JSONData)
		if err != nil {
			fmt.Fprint(w, "Error while marshaling JSON.")
		}
		fmt.Fprint(w, string(result))
	case ViewResponse:
		content, err := ioutil.ReadFile(res.TextData)
		if err != nil {
			log.Println("File " + res.TextData + " not found.")
		}
		fmt.Fprintf(w, string(content))
	case TemplateResponse:
		template.Must(template.ParseFiles(res.TextData)).Execute(w, res.JSONData)
	default:
		fmt.Fprint(w, "Invalid HTTP response type.")
	}
}
