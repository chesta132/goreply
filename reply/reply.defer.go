package reply

// Defer registers one or more functions to be executed before the response is sent.
// These functions are executed in the order they were registered.
// Useful for cleanup operations, logging, or committing/rolling back transactions.
//
// Example:
//
//	rp.Defer(func() { tx.Rollback() })
//	rp.Defer(func() { log.Println("Response sent") })
//	rp.Success(data).OkJSON() // deferred funcs execute before sending
func (r *Reply) Defer(funcs ...func()) *Reply {
	r.defers = append(r.defers, funcs...)
	return r
}

// execDefer executes all deferred functions in the order they were registered.
func (r *Reply) execDefer() {
	for _, f := range r.defers {
		f()
	}
}
