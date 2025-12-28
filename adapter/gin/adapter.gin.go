package adapter

import (
	"io"
	"net/http"

	"github.com/chesta132/goreply/adapter"
	"github.com/gin-gonic/gin"
)

// ginAdapter wraps gin.Context to implement adapter.Adapter.
//
// Example:
//
//	Client.New(adapter.AdaptGin(c)).Success(data).OkJSON()
type ginAdapter struct {
	ctx *gin.Context
}

// Adapt converts gin.Context into an Adapter.
//
// Example:
//
//	rp := Client.New(adapter.AdaptGin(c))
func AdaptGin(c *gin.Context) adapter.Adapter {
	return &ginAdapter{ctx: c}
}

// Header returns the response headers map.
func (g *ginAdapter) Header() http.Header {
	return g.ctx.Writer.Header()
}

// Write writes raw bytes to the response.
func (g *ginAdapter) Write(b []byte) (int, error) {
	return g.ctx.Writer.Write(b)
}

// Write status header
func (g *ginAdapter) SetStatus(statusCode int) {
	g.ctx.Status(statusCode)
}

// JsonSender writes JSON response with given status.
//
// Please use reply to handle this sender.
//
// Example:
//
//	g.JsonSender(200, map[string]any{"msg": "ok"}) // -> {"msg":"ok"}
func (g *ginAdapter) JsonSender(statusCode int, data interface{}) error {
	g.ctx.JSON(statusCode, data)
	return nil
}

// XmlSender writes XML response with given status.
//
// Please use reply to handle this sender.
//
// Example:
//
//	g.XmlSender(200, User{ID: 1}) // -> <User><ID>1</ID></User>
func (g *ginAdapter) XmlSender(statusCode int, data interface{}) error {
	g.ctx.XML(statusCode, data)
	return nil
}

// BinarySender writes raw bytes as application/octet-stream.
//
// Please use reply to handle this sender.
//
// Example:
//
//	g.BinarySender(200, []byte{0xFF, 0xD8}) // JPEG SOI marker
func (g *ginAdapter) BinarySender(statusCode int, data []byte) error {
	g.ctx.Data(statusCode, "application/octet-stream", data)
	return nil
}

// TextSender writes plain text response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	g.TextSender(200, "Hello!") // -> Hello!
func (g *ginAdapter) TextSender(statusCode int, text string) error {
	g.ctx.String(statusCode, text)
	return nil
}

// HtmlSender writes HTML string response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	g.HtmlSender(200, "<h1>Hi</h1>") // -> <h1>Hi</h1>
func (g *ginAdapter) HtmlSender(statusCode int, html string) error {
	g.ctx.Data(statusCode, "text/html; charset=utf-8", []byte(html))
	return nil
}

// StreamSender streams data with custom Content-Type.
//
// Please use reply to handle this sender.
//
// Example:
//
//	g.StreamSender(200, "video/mp4", fileReader) // streams MP4
func (g *ginAdapter) StreamSender(statusCode int, contentType string, reader io.Reader) error {
	g.ctx.Status(statusCode)
	g.ctx.Header("Content-Type", contentType)
	_, err := io.Copy(g.ctx.Writer, reader)
	return err
}

// RedirectSender sends a redirect response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.RedirectSender(http.StatusMovedPermanently, "https://github.com/chesta132")
func (a *ginAdapter) RedirectSender(statusCode int, url string) {
	a.ctx.Redirect(statusCode, url)
}

// Get read context value.
//
// Please use reply to handle this sender.
//
// Example:
//
//	type instance struct{}
//
//	var replyInstance instance
//	a.Get(replyInstance)
func (a *ginAdapter) Get(key any) (any, bool) {
	return a.ctx.Get(key)
}

// Set sets value to request context.
//
// Please use reply to handle this sender.
//
// Example:
//
//	type instance struct{}
//
//	var replyInstance instance
//	a.Set(replyInstance, *reply)
func (a *ginAdapter) Set(key, value any) {
	a.ctx.Set(key, value)
}
