package reply

import "errors"

var (
	ErrAlreadySent = errors.New("reply: can not send more data, response already sent")
)
