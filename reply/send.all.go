package reply

import "net/http"

// NoContent sends no content response with status 204
//
// Example:
// 	rp.NoContent()
func (r *Reply) NoContent() {
	r.send(func() error {
		r.a.Header().Del("Content-Type")
		r.a.SetStatus(http.StatusNoContent)
		return nil
	}, 1)
}

// Redirect sends a redirect response with given status code.
//
// Example:
// 	rp.Redirect(http.StatusMovedPermanently, "https://github.com/chesta132")
func (r *Reply) Redirect(statusCode int, url string) {
	r.send(func() error {
		r.a.RedirectSender(statusCode, url)
		return nil
	}, 1)
}
