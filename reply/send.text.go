package reply

import (
	"errors"
	"net/http"
)

// ReplyText sends a plain text response with the specified status code.
// Data must be a string; if not, logs an error and sends empty string.
func (r *Reply) replyText(code int) error {
	return r.send(func() error {
		d, ok := r.m.Data.(string)
		if !ok {
			logError(errors.New("DATA TYPE IS NOT STRING"), 2)
			d = ""
		}
		return r.a.TextSender(code, d)
	}, 2)
}

// ReplyText sends a plain text response with the specified status code.
// Data must be a string; if not, logs an error and sends empty string.
//
// Example:
//
//	rp.Success("Hello world!").ReplyText(http.StatusOK) // -> Hello world!
func (r *Reply) ReplyText(code int) error {
	return r.replyText(code)
}

// OkText is a shortcut for ReplyText with status 200 OK.
func (r *Reply) OkText() error {
	return r.replyText(http.StatusOK)
}

// CreatedText sends status 201 Created with plain text body.
func (r *Reply) CreatedText() error {
	return r.replyText(http.StatusCreated)
}
