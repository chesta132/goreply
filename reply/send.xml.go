package reply

import "net/http"

// ReplyXML sends an XML-formatted response with the specified status code.
// Payload will be marshaled to XML automatically.
func (r *Reply) replyXML(code int) error {
	return r.send(func() error {
		return r.a.XmlSender(code, r.Payload)
	}, 2)
}

// ReplyXML sends an XML-formatted response with the specified status code.
// Payload will be marshaled to XML automatically.
//
// Example:
//		rp.Success(User{ID: 1, Name: "Chesta"}).ReplyXML(http.StatusOK)
//		// -> <User><ID>1</ID><Name>Chesta</Name></User>
func (r *Reply) ReplyXML(code int) error {
	return r.replyXML(code)
}

// OkXML is a shortcut for ReplyXML with status 200 OK.
func (r *Reply) OkXML() error {
	return r.replyXML(http.StatusOK)
}

// CreatedXML sends status 201 Created with XML body.
func (r *Reply) CreatedXML() error {
	return r.replyXML(http.StatusCreated)
}

// FailXML sends a XML response with an error status.
// If code is provided, use it; otherwise, retrieve from codeAliases
// or default to 500.
func (r *Reply) FailXML(code ...int) error {
	c, _ := r.retrieveStatusCode(code...)
	return r.replyXML(c)
}
