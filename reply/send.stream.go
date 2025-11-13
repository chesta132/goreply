package reply

import (
	"errors"
	"net/http"
)

// ReplyStream sends a streaming response with the specified status code.
// Data must be of type Stream; otherwise, logs an error and sends an empty stream.
func (r *Reply) replyStream(code int) error {
	return r.send(func() error {
		d, ok := r.m.Data.(Stream)
		if !ok {
			logError(errors.New("DATA TYPE IS NOT STREAM"), 2)
			return r.a.StreamSender(code, "", nil)
		}
		return r.a.StreamSender(code, d.ContentType, d.Data)
	}, 2)
}

// ReplyStream sends a streaming response with the specified status code.
// Data must be of type Stream; otherwise, logs an error and sends an empty stream.
//
// Example:
//
//	rp.Success(reply.Stream{Data: file, ContentType: "image/png"}).
//	ReplyStream(http.StatusOK) // streams the file as PNG
func (r *Reply) ReplyStream(code int) error {
	return r.replyStream(code)
}

// OkStream is a shortcut for ReplyStream with status 200 OK.
func (r *Reply) OkStream() error {
	return r.replyStream(http.StatusOK)
}

// CreatedStream sends status 201 Created with stream body.
func (r *Reply) CreatedStream() error {
	return r.replyStream(http.StatusCreated)
}
