package reply

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

// Retrieve optional status code or take code alias if status code is empty.
func (r *Reply) retrieveStatusCode(code ...int) (c int, ok bool) {
	if len(code) > 0 {
		return code[0], true
	}
	d, ok := r.m.Data.(ErrorPayload)
	if r.c.CodeAliases != nil && ok {
		if code, exists := r.c.CodeAliases[d.Code]; exists {
			return code, true
		}
	}
	return http.StatusInternalServerError, false
}

// Validate is reply has already sent.
// Finalize reply.
// Send data and make sure it won't send more data.
func (r *Reply) send(sender func() error, callerSkip int) error {
	if r.sent {
		return ErrAlreadySent
	}
	r.finalize()
	r.execDefer()

	if err := sender(); err != nil {
		logError(fmt.Errorf("ERROR WHILE REPLYING\n%v", err), callerSkip+1)
		return err
	}

	r.sent = true
	return nil
}

// Finalize reply by execute finalizer config and transform the payload
func (r *Reply) finalize() {
	if r.c.Finalizer != nil {
		r.c.Finalizer(r)
	}

	if r.c.Transformer != nil {
		r.Payload = r.c.Transformer(r)
	} else {
		envelope := &ReplyEnvelope{Meta: r.m.Meta, Data: r.m.Data}
		if r.c.DebugMode {
			envelope.Debug = r.m.Debug
		}
		r.Payload = envelope
	}
}

// Log an error
func logError(err error, callerSkip int) {
	_, file, line, _ := runtime.Caller(callerSkip)
	log.Printf("\n\n\n\n-------------------\n%s:%d\nREPLY-ERROR: %v\n-------------------\n\n\n\n", file, line, err.Error())
}
