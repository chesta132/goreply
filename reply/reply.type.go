package reply

import (
	"io"

	"github.com/chesta132/goreply/adapter"
)

// Transformer defines a function to transform the payload before sending.
type Transformer func(client *Reply) any

// Finalizer defines a function to run before sending the response.
type Finalizer func(client *Reply)

// CodeAliases maps error codes to HTTP status codes.
type CodeAliases map[string]int

// DefaultHeaders defines default headers to include in every response.
type DefaultHeaders map[string]string

// PaginationType defines the pagination strategy.
type PaginationType string

const (
	// PaginationPage indicates page-based pagination.
	PaginationPage PaginationType = "page"
	// PaginationOffset indicates offset-based pagination.
	PaginationOffset PaginationType = "offset"
)

// Tokens holds authentication or session tokens.
type Tokens map[string]string

// Pagination holds pagination metadata, embedded in Meta when needed.
//
// Example:
//
//	Pagination{Next: 20, HasNext: true, Current: 0}
type Pagination struct {
	Next    int  `json:"next" xml:"next"`                       // Next page/offset
	HasNext bool `json:"hasNext" xml:"hasNext"`                 // True if more results exist
	Current int  `json:"current" xml:"current"`                 // Current page/offset
	Total   int  `json:"total,omitempty" xml:"total,omitempty"` // Total data, available if use total paginate
}

// Meta contains reply metadata.
//
// Example:
//
//	Meta{Status: "SUCCESS", Info: "Fetched 10 items"}
type Meta struct {
	Status     string      `json:"status" xml:"status"`                               // "SUCCESS" or "ERROR"
	Info       string      `json:"information,omitempty" xml:"information,omitempty"` // Optional info message
	Pagination *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`   // Pagination info if applicable
	Tokens     Tokens      `json:"tokens,omitempty" xml:"tokens,omitempty"`           // Optional tokens (e.g. auth)
}

// ReplyEnvelope is the standard API response envelope.
//
// Example:
//
//	ReplyEnvelope{Meta: Meta{Status: "SUCCESS"}, Data: user}
type ReplyEnvelope struct {
	Meta Meta `json:"meta" xml:"meta"` // Metadata section
	Data any  `json:"data" xml:"data"` // Payload data
}

// ErrorPayload defines the error response body.
//
// Example:
//
//	ErrorPayload{Code: "NOT_FOUND", Message: "User not found"}
type ErrorPayload struct {
	Code    string `json:"code" xml:"code"`                           // Machine-readable error code
	Message string `json:"message" xml:"message"`                     // Human-readable message
	Details string `json:"details,omitempty" xml:"details,omitempty"` // Optional debug details
	Field   string `json:"field,omitempty" xml:"field,omitempty"`     // Field causing the error (if any)
}

// OptErrorPayload holds optional error fields for partial errors.
//
// Example:
//
//	OptErrorPayload{Details: "Invalid email format", Field: "email"}
type OptErrorPayload struct {
	Details string `json:"details,omitempty" xml:"details,omitempty"` // Optional debug details
	Field   string `json:"field,omitempty" xml:"field,omitempty"`     // Field causing the error (if any)
}

// Reply is the main HTTP response helper with chained methods.
type Reply struct {
	Payload any // Transformed payload. Only available after reply

	m      *ReplyEnvelope  // Internal payload
	a      adapter.Adapter // Response adapter (e.g. gin, echo)
	c      *Client         // Client config
	sent   bool            // True after response is sent
	defers []func()        // Functions to execute before sending
}

// Client holds global config for Reply instances.
type Client struct {
	Finalizer      Finalizer      // Runs before sending
	Transformer    Transformer    // Transforms payload
	CodeAliases    CodeAliases    // Maps error codes to HTTP status
	DefaultHeaders DefaultHeaders // Default response headers
	PaginationType PaginationType // "page" or "offset". Default: "offset"
}

// Stream enables streaming responses (files, SSE, etc.).
//
// Example:
//
//	Stream{Data: fileReader, ContentType: "video/mp4"}
type Stream struct {
	Data        io.Reader // Stream source
	ContentType string    // MIME type of the stream
}
