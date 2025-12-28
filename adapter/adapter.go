// Package adapter provides a unified interface for sending HTTP responses
// across different web frameworks (Gin, Echo, Fiber) and standard net/http.
//
// This adapter pattern allows you to write framework-agnostic HTTP handlers
// that can work with any supported framework by simply wrapping the framework's
// context with the appropriate adapter.
//
// Example usage with net/http:
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//	    a := adapter.AdaptHttp(w, r)
//	    a.JsonSender(200, map[string]string{"status": "ok"})
//	}
package adapter

import (
	"io"
	"net/http"
)

// Adapter provides a unified interface for sending HTTP responses.
// It abstracts away the differences between various Go web frameworks,
// allowing you to write framework-agnostic code.
type Adapter interface {
	// JsonSender sends a JSON response with the given status code and data.
	// The data will be automatically marshaled to JSON.
	JsonSender(statusCode int, data interface{}) error

	// XmlSender sends an XML response with the given status code and data.
	// The data will be automatically marshaled to XML.
	XmlSender(statusCode int, data interface{}) error

	// BinarySender sends binary data with application/octet-stream content type.
	// Useful for sending files, images, or any binary content.
	BinarySender(statusCode int, data []byte) error

	// TextSender sends plain text response with UTF-8 encoding.
	TextSender(statusCode int, text string) error

	// HtmlSender sends HTML content with proper content type.
	// Note: This does not escape HTML entities. Use html.EscapeString if needed.
	HtmlSender(statusCode int, html string) error

	// StreamSender streams data from an io.Reader to the response.
	// Useful for streaming large files, video, or real-time data.
	StreamSender(statusCode int, contentType string, reader io.Reader) error

	// RedirectSender sends a redirect response.
	RedirectSender(statusCode int, url string)

	// Write implements io.Writer interface, allowing direct writes to the response.
	Write([]byte) (int, error)

	// Header returns the HTTP headers that will be sent with the response.
	// You can use this to set custom headers before sending the response.
	Header() http.Header

	// Set status code to header
	SetStatus(statusCode int)

	// Get contexted value
	Get(key any) (value any, ok bool)

	// Set value to request context
	Set(key, value any)
}
