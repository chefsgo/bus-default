package bus

import "github.com/chefsgo/bus"

func Driver() bus.Driver {
	return &defaultDriver{}
}

func init() {
	bus.Register("default", Driver())
}
