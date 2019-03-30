package response

// Type is the name of the response type.
type Type uint

// Indicate the available response types.
const (
	TextResponse Type = iota
	JSONResponse
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
