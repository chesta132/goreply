package reply

import (
	"os"

	"github.com/chesta132/goreply/adapter"
)

type clientKeyType struct{}

var clientKey clientKeyType

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
// 		DebugMode: os.GetEnv("GO_ENV") != "production"
//	})
func NewClient(config Client) *Client {
	return &config
}

// Create a new reply.
//
// Example:
//
//	rp := Client.New(nethttpadapter.Adapt(w, r))
//	// ...
//	rp.Success(datas).OkJSON()
func (c *Client) New(adapter adapter.Adapter) *Reply {
	// create reply instance
	rp := &Reply{a: adapter, c: c, m: &ReplyEnvelope{}}

	// warn env
	if c.DebugMode && os.Getenv("GO_ENV") == "production" {
		logGoReply("WARNING: DebugMode enabled in production")
	}

	// set headers
	if c.DefaultHeaders != nil {
		for k, v := range c.DefaultHeaders {
			rp.a.Header().Set(k, v)
		}
	}

	return rp
}

// Reuse instance or create new Reply instance.
//
// Example:
//
//	rp := Client.Use(nethttpadapter.Adapt(w, r))
//	// ...
//	rp.Success(datas).OkJSON()
func (c *Client) Use(adapter adapter.Adapter) *Reply {
	// get or create reply instance
	var rp *Reply
	if val, ok := adapter.Get(clientKey); ok {
		if rp, ok := val.(*Reply); ok {
			return rp
		}
	}

	// create and set reply instance
	rp = c.New(adapter)
	adapter.Set(clientKey, rp)

	return rp
}
