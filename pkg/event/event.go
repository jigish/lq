package event

import (
	"github.com/rs/zerolog"
)

type Event interface {
	zerolog.LogObjectMarshaler
}

type Map map[string]any

var _ Event = Map{}

func (e Map) MarshalZerologObject(ev *zerolog.Event) {
	for k, v := range e {
		ev.Any(k, v)
	}
}

type Error string

var _ Event = Error("")

func (e Error) MarshalZerologObject(ev *zerolog.Event) {
	ev.Str("error", string(e))
}

func (e Error) Error() string {
	return string(e)
}

type InvalidFormat string

var _ Event = InvalidFormat("")

func (e InvalidFormat) MarshalZerologObject(ev *zerolog.Event) {
	ev.Msg(string(e))
}
