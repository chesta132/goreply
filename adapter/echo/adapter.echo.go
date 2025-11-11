package adapter

import (
	"io"
	"net/http"

	"github.com/chesta132/goreply/adapter"
	"github.com/labstack/echo/v4"
)

// echoAdapter wraps echo.Context to implement adapter.Adapter.
//
// Example:
//
//	Client.New(adapter.AdaptEcho(c)).Success(data).OkJSON()
type echoAdapter struct {
	ctx echo.Context
}

// Adapt converts echo.Context into an Adapter.
//
// Example:
//
//	rp := Client.New(adapter.AdaptEcho(c))
func AdaptEcho(c echo.Context) adapter.Adapter {
	return &echoAdapter{ctx: c}
}

// Header returns the response headers map.
func (e *echoAdapter) Header() http.Header {
	return e.ctx.Response().Header()
}

// Write writes raw bytes to the response.
func (e *echoAdapter) Write(b []byte) (int, error) {
	return e.ctx.Response().Writer.Write(b)
}

// Write status header
func (g *echoAdapter) SetStatus(statusCode int) {
	g.ctx.Response().Writer.WriteHeader(statusCode)
}

// JsonSender writes JSON response with given status.
//
// Please use reply to handle this sender.
//
// Example:
//
//	e.JsonSender(200, map[string]any{"msg": "ok"}) // -> {"msg":"ok"}
func (e *echoAdapter) JsonSender(statusCode int, data interface{}) error {
	return e.ctx.JSON(statusCode, data)
}

// XmlSender writes XML response with given status.
//
// Please use reply to handle this sender.
//
// Example:
//
//	e.XmlSender(200, User{ID: 1}) // -> <User><ID>1</ID></User>
func (e *echoAdapter) XmlSender(statusCode int, data interface{}) error {
	return e.ctx.XML(statusCode, data)
}

// BinarySender writes raw bytes as application/octet-stream.
//
// Please use reply to handle this sender.
//
// Example:
//
//	e.BinarySender(200, []byte{0xFF, 0xD8}) // JPEG SOI marker
func (e *echoAdapter) BinarySender(statusCode int, data []byte) error {
	return e.ctx.Blob(statusCode, "application/octet-stream", data)
}

// TextSender writes plain text response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	e.TextSender(200, "Hello!") // -> Hello!
func (e *echoAdapter) TextSender(statusCode int, text string) error {
	return e.ctx.String(statusCode, text)
}

// HtmlSender writes HTML string response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	e.HtmlSender(200, "<h1>Hi</h1>") // -> <h1>Hi</h1>
func (e *echoAdapter) HtmlSender(statusCode int, html string) error {
	return e.ctx.HTML(statusCode, html)
}

// StreamSender streams data with custom Content-Type.
//
// Please use reply to handle this sender.
//
// Example:
//
//	e.StreamSender(200, "video/mp4", fileReader) // streams MP4
func (e *echoAdapter) StreamSender(statusCode int, contentType string, reader io.Reader) error {
	e.ctx.Response().Header().Set("Content-Type", contentType)
	e.ctx.Response().WriteHeader(statusCode)
	_, err := io.Copy(e.ctx.Response().Writer, reader)
	return err
}

// RedirectSender sends a redirect response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.RedirectSender(http.StatusMovedPermanently, "https://github.com/chesta132")
func (a *echoAdapter) RedirectSender(statusCode int, url string) {
	a.ctx.Redirect(statusCode, url)
}
