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

// Success sets reply status to "SUCCESS" and attaches data.
func (r *Reply) Success(data any) *Reply {
	r.setStatus("SUCCESS")
	r.setData(data)
	return r
}

// Error sets reply status to "ERROR" and attaches an error payload.
func (r *Reply) Error(code, message string, options ...ErrorOption) *Reply {
	r.setStatus("ERROR")
	payload := ErrorPayload{Code: code, Message: message}

	// set optional values
	for _, opt := range options {
		opt(&payload)
	}

	r.setData(payload)
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
		r.m.Meta.Debug = messages[0]
	} else if len(messages) > 0 {
		r.m.Meta.Debug = messages
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

// WithDetails returns ErrorOption to build error with details.
func WithDetails(details string) ErrorOption {
	return func(ep *ErrorPayload) {
		ep.Details = details
	}
}

// WithFields returns ErrorOption to build error with fields error.
func WithFields(fields FieldsError) ErrorOption {
	return func(ep *ErrorPayload) {
		ep.Fields = fields
	}
}
