package reply

import (
	"errors"
	"html"
	"net/http"
)

// ReplyHTML sends an HTML string response with the specified status code.
// The Data in *Reply must be a string; if not, it will be treated as an empty string
// and an error will be logged. The string is automatically escaped before sending.
func (r *Reply) replyHTML(code int) {
	r.send(func() error {
		d, ok := r.m.Data.(string)
		if !ok {
			logError(errors.New("DATA TYPE IS NOT HTML STRING"), 2)
			d = ""
		}
		escaped := html.EscapeString(d)
		return r.a.HtmlSender(code, escaped)
	}, 2)
}

// ReplyHTML sends an HTML string response with the specified status code.
// The Data in *Reply must be a string; if not, it will be treated as an empty string
// and an error will be logged. The string is automatically escaped before sending.
//
// Example:
//
//	rp.Success("<p>Hello & welcome!</p>").ReplyHTML(http.StatusOK) // -> <p>Hello &amp; welcome!</p>
func (r *Reply) ReplyHTML(code int) {
	r.replyHTML(code)
}

// OkHTML is a shortcut for ReplyHTML with status 200 OK.
func (r *Reply) OkHTML() {
	r.replyHTML(http.StatusOK)
}

// CreatedHTML sends status 201 Created with HTML body.
func (r *Reply) CreatedHTML() {
	r.replyHTML(http.StatusCreated)
}
