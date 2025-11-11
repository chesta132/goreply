package adapter

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/chesta132/goreply/adapter"
)

// netHttpAdapter wraps net/http ResponseWriter to implement adapter.Adapter.
//
// Example:
//
//	Client.New(nethttpadapter.AdaptHttp(w, r)).Success(data).OkJSON()
type netHttpAdapter struct {
	w http.ResponseWriter
	r *http.Request
}

// Adapt converts http.ResponseWriter into an Adapter.
//
// Example:
//
//	rp := Client.New(nethttpadapter.AdaptHttp(w, r))
func AdaptHttp(w http.ResponseWriter, r *http.Request) adapter.Adapter {
	return &netHttpAdapter{w: w, r: r}
}

// Header returns the response headers map.
func (a *netHttpAdapter) Header() http.Header {
	return a.w.Header()
}

// Write writes raw bytes to the response.
func (a *netHttpAdapter) Write(b []byte) (int, error) {
	return a.w.Write(b)
}

// Write status header
func (a *netHttpAdapter) SetStatus(statusCode int) {
	a.w.WriteHeader(statusCode)
}

// JsonSender writes JSON response with given status.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.JsonSender(200, map[string]any{"msg": "ok"}) // -> {"msg":"ok"}
func (a *netHttpAdapter) JsonSender(statusCode int, data interface{}) error {
	a.w.Header().Set("Content-Type", "application/json")
	a.w.WriteHeader(statusCode)
	return json.NewEncoder(a.w).Encode(data)
}

// XmlSender writes XML response with given status.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.XmlSender(200, User{ID: 1}) // -> <User><ID>1</ID></User>
func (a *netHttpAdapter) XmlSender(statusCode int, data interface{}) error {
	a.w.Header().Set("Content-Type", "application/xml")
	a.w.WriteHeader(statusCode)
	return xml.NewEncoder(a.w).Encode(data)
}

// BinarySender writes raw bytes as application/octet-stream.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.BinarySender(200, []byte{0xFF, 0xD8}) // JPEG SOI marker
func (a *netHttpAdapter) BinarySender(statusCode int, data []byte) error {
	a.w.Header().Set("Content-Type", "application/octet-stream")
	a.w.WriteHeader(statusCode)
	_, err := a.w.Write(data)
	return err
}

// TextSender writes plain text response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.TextSender(200, "Hello!") // -> Hello!
func (a *netHttpAdapter) TextSender(statusCode int, text string) error {
	a.w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	a.w.WriteHeader(statusCode)
	_, err := a.w.Write([]byte(text))
	return err
}

// HtmlSender writes HTML string response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.HtmlSender(200, "<h1>Hi</h1>") // -> <h1>Hi</h1>
func (a *netHttpAdapter) HtmlSender(statusCode int, html string) error {
	a.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	a.w.WriteHeader(statusCode)
	_, err := a.w.Write([]byte(html))
	return err
}

// StreamSender streams data with custom Content-Type.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.StreamSender(200, "video/mp4", fileReader) // streams MP4
func (a *netHttpAdapter) StreamSender(statusCode int, contentType string, reader io.Reader) error {
	a.w.Header().Set("Content-Type", contentType)
	a.w.WriteHeader(statusCode)
	_, err := io.Copy(a.w, reader)
	return err
}

// RedirectSender sends a redirect response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.RedirectSender(http.StatusMovedPermanently, "https://github.com/chesta132")
func (a *netHttpAdapter) RedirectSender(statusCode int, url string) {
	http.Redirect(a.w, a.r, url, statusCode)
}
