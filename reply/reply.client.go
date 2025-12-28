package reply

import (
	"os"

	"github.com/chesta132/goreply/adapter"
)

// Create a new client with given configuration.
//
// Example:
//
//	var Client = reply.NewClient(reply.Client{
//		Finalizer: finalizer,
//		Transformer: transformer,
//		CodeAliases    map[string]int{
//			"SERVER_ERROR": 500,
//			"BAD_REQUEST": 400,
//		},
//		DefaultHeaders map[string]string{
//			"Content-Type": "application/json",
//		},
//	})
func NewClient(config Client) *Client {
	return &config
}

// Create a new reply.
//
// Example:
//
//	rp := Client.New(nethttpadapter.Adapt(w))
//	// ...
//	rp.Success(datas).OkJSON()
func (c *Client) New(adapter adapter.Adapter) *Reply {
	rp := &Reply{a: adapter, c: c, m: &ReplyEnvelope{}}
	if c.DebugMode && os.Getenv("GO_ENV") == "production" {
		logGoReply("WARNING: DebugMode enabled in production")
	}
	if c.DefaultHeaders != nil {
		for k, v := range c.DefaultHeaders {
			rp.a.Header().Set(k, v)
		}
	}
	return rp
}
