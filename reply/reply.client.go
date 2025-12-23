package reply

import "github.com/chesta132/goreply/adapter"

// Create a new client with given configuration.
//
// Example:
//		var Client = reply.NewClient(reply.Client{
//			finalizer: finalizer,
// 			transformer: transformer,
// 			codeAliases: map[string]int{
//				"SERVER_ERROR": 500,
//				"BAD_REQUEST": 400,
// 			},
// 			defaultHeaders map[string]string{
// 				"Content-Type": "application/json",
// 			},
//		})
func NewClient(config Client) *Client {
	return &config
}

// Create a new reply.
//
// Example:
//		rp := Client.New(nethttpadapter.Adapt(w))
// 		// ...
// 		rp.Success(datas).OkJSON()
func (c *Client) New(adapter adapter.Adapter) *Reply {
	rp := &Reply{a: adapter, c: c, m: &ReplyEnvelope{}}
	if c.defaultHeaders != nil {
		for k, v := range c.defaultHeaders {
			rp.a.Header().Set(k, v)
		}
	}
	return rp
}
