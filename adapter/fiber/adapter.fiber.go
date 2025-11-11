package adapter

import (
	"bufio"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/chesta132/reply-go/adapter"
	"github.com/gofiber/fiber/v2"
)

// fiberAdapter wraps fiber.Ctx to implement adapter.Adapter.
//
// Example:
//
//	Client.New(adapter.AdaptFiber(c)).Success(data).OkJSON()
type fiberAdapter struct {
	ctx *fiber.Ctx
}

// Adapt converts fiber.Ctx into an Adapter.
//
// Example:
//
//	rp := Client.New(adapter.AdaptFiber(c))
func AdaptFiber(c *fiber.Ctx) adapter.Adapter {
	return &fiberAdapter{ctx: c}
}

// Header returns the response headers map.
func (f *fiberAdapter) Header() http.Header {
	header := make(http.Header)
	f.ctx.Response().Header.VisitAll(func(key, value []byte) {
		header.Add(string(key), string(value))
	})
	return header
}

// Write writes raw bytes to the response.
func (f *fiberAdapter) Write(b []byte) (int, error) {
	f.ctx.Write(b)
	return len(b), nil
}

// Write status header
func (g *fiberAdapter) SetStatus(statusCode int) {
	g.ctx.Status(statusCode)
}

// JsonSender writes JSON response with given status.
//
// Please use reply to handle this sender.
//
// Example:
//
//	f.JsonSender(200, map[string]any{"msg": "ok"}) // -> {"msg":"ok"}
func (f *fiberAdapter) JsonSender(statusCode int, data interface{}) error {
	f.ctx.Status(statusCode)
	return f.ctx.JSON(data)
}

// XmlSender writes XML response with given status.
//
// Please use reply to handle this sender.
//
// Example:
//
//	f.XmlSender(200, User{ID: 1}) // -> <User><ID>1</ID></User>
func (f *fiberAdapter) XmlSender(statusCode int, data interface{}) error {
	f.ctx.Status(statusCode)
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return err
	}
	f.ctx.Set("Content-Type", "application/xml")
	return f.ctx.Send(xmlData)
}

// BinarySender writes raw bytes as application/octet-stream.
//
// Please use reply to handle this sender.
//
// Example:
//
//	f.BinarySender(200, []byte{0xFF, 0xD8}) // JPEG SOI marker
func (f *fiberAdapter) BinarySender(statusCode int, data []byte) error {
	f.ctx.Status(statusCode)
	f.ctx.Set("Content-Type", "application/octet-stream")
	return f.ctx.Send(data)
}

// TextSender writes plain text response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	f.TextSender(200, "Hello!") // -> Hello!
func (f *fiberAdapter) TextSender(statusCode int, text string) error {
	f.ctx.Status(statusCode)
	return f.ctx.SendString(text)
}

// HtmlSender writes HTML string response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	f.HtmlSender(200, "<h1>Hi</h1>") // -> <h1>Hi</h1>
func (f *fiberAdapter) HtmlSender(statusCode int, html string) error {
	f.ctx.Status(statusCode)
	f.ctx.Set("Content-Type", "text/html; charset=utf-8")
	return f.ctx.SendString(html)
}

// StreamSender streams data with custom Content-Type.
//
// Please use reply to handle this sender.
//
// Example:
//
//	f.StreamSender(200, "video/mp4", fileReader) // streams MP4
func (f *fiberAdapter) StreamSender(statusCode int, contentType string, reader io.Reader) error {
	f.ctx.Status(statusCode)
	f.ctx.Set("Content-Type", contentType)
	f.ctx.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		io.Copy(w, reader)
		w.Flush()
	})
	return nil
}

// RedirectSender sends a redirect response.
//
// Please use reply to handle this sender.
//
// Example:
//
//	a.RedirectSender(http.StatusMovedPermanently, "https://github.com/chesta132")
func (a *fiberAdapter) RedirectSender(statusCode int, url string) {
	a.ctx.Redirect(url, statusCode)
}
