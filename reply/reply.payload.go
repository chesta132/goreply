package reply

// setStatus sets the "meta.status" field value.
// Returns the Reply for chaining.
func (r *Reply) setStatus(status string) *Reply {
	r.m.Meta.Status = status
	return r
}

// setData assigns data to the "data" field.
// Returns the Reply for chaining.
func (r *Reply) setData(data any) *Reply {
	r.m.Data = data
	return r
}

// Success marks the reply as successful and attaches data.
func (r *Reply) Success(data any) *Reply {
	r.setStatus("SUCCESS")
	r.setData(data)
	return r
}

// Error sets reply status to "ERROR" and attaches an error payload.
func (r *Reply) Error(code, message string, optional ...OptErrorPayload) *Reply {
	r.setStatus("ERROR")
	o := OptErrorPayload{}
	if len(optional) > 0 {
		o = optional[0]
	}
	r.setData(ErrorPayload{code, message, o.Details, o.Fields})
	return r
}

// Info sets info to reply meta information.
func (r *Reply) Info(information string) *Reply {
	r.m.Meta.Info = information
	return r
}

// Tokens sets tokens to reply meta tokens.
func (r *Reply) Tokens(tokens Tokens) *Reply {
	r.m.Meta.Tokens = tokens
	return r
}

// Debug sets debug messages to envelope.
func (r *Reply) Debug(messages ...any) *Reply {
	if len(messages) == 1 {
		r.m.Debug = messages[0]
	} else if len(messages) > 0 {
		r.m.Debug = messages
	}
	return r
}

// Envelope returns a copy of the internal envelope.
// Modifying the returned value does not affect the original.
func (r *Reply) Envelope() ReplyEnvelope {
	return *r.m
}

// Data returns a copy of the data from internal envelope.
// Modifying the returned value does not affect the original.
func (r *Reply) Data() any {
	return r.Envelope().Data
}

// Meta returns a copy of the meta from internal envelope.
// Modifying the returned value does not affect the original.
func (r *Reply) Meta() Meta {
	return r.Envelope().Meta
}

// GetDebug returns a copy of the debug from internal envelope.
// Modifying the returned value does not affect the original.
func (r *Reply) GetDebug() any {
	return r.Envelope().Debug
}
