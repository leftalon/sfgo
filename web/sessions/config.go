package sessions

import "sfgo/core/config"

var C struct {
	Name   string
	Key    string
	MaxAge int
}

func init() {
	C.Name = "_DEFAULT_NAME_"
	C.MaxAge = 3153600000

	config.Register("tcgo.session", &C)
}
