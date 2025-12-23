package reply

import "net/http"

// ReplyJSON sends a JSON-formatted response with the specified status code.
// The Payload in *Reply will be automatically encoded.
func (r *Reply) replyJSON(code int) error {
	return r.send(func() error {
		return r.a.JsonSender(code, r.Payload)
	}, 2)
}

// ReplyJSON sends a JSON-formatted response with the specified status code.
// The Payload in *Reply will be automatically encoded.
//
// Example:
//		rp.Success(Data{Msg: "ok"}).ReplyJSON(http.StatusOK) // -> {"msg":"ok"}
func (r *Reply) ReplyJSON(code int) error {
	return r.replyJSON(code)
}

// OkJSON is a shortcut for ReplyJSON with status 200 OK.
func (r *Reply) OkJSON() error {
	return r.replyJSON(http.StatusOK)
}

// CreatedJSON sends status 201 Created with JSON body.
func (r *Reply) CreatedJSON() error {
	return r.replyJSON(http.StatusCreated)
}

// FailJSON sends a JSON response with an error status.
// If code is provided, use it; otherwise, retrieve from codeAliases
// or default to 500.
func (r *Reply) FailJSON(code ...int) error {
	c, _ := r.retrieveStatusCode(code...)
	return r.replyJSON(c)
}
