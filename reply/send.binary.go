package reply

import (
	"errors"
	"net/http"
)

// ReplyBinary sends a binary response using the adapter's BinarySender.
// If Data is not []byte, it logs an error and sends an empty body.
func (r *Reply) replyBinary(code int) {
	r.send(func() error {
		d, ok := r.m.Data.([]byte)
		if !ok {
			logError(errors.New("DATA TYPE IS NOT BYTE"), 2)
			return r.a.BinarySender(code, []byte{})
		}
		return r.a.BinarySender(code, d)
	}, 2)
}

// ReplyBinary sends a binary response using the adapter's BinarySender.
// If Data is not []byte, it logs an error and sends an empty body.
//
// Example:
//
//	rp.Success(imageData).OkBinary()
func (r *Reply) ReplyBinary(code int) {
	r.replyBinary(code)
}

// OkBinary is a shortcut for ReplyBinary with http.StatusOK (200).
func (r *Reply) OkBinary() {
	r.replyBinary(http.StatusOK)
}

// CreatedBinary is a shortcut for ReplyBinary with http.StatusCreated (201).
func (r *Reply) CreatedBinary() {
	r.replyBinary(http.StatusCreated)
}
