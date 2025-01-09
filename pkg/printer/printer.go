package printer

import (
	"io"

	"github.com/rs/zerolog"

	"github.com/jigish/lq/pkg/event"
)

var (
	timestampFields = []string{
		"ts",
		"time",
		"timestamp",
		"@timestamp",
	}
	levelFields = []string{
		"level",
		"log.level",
	}
	messageFields = []string{
		"msg",
		"message",
	}
)

type Printer struct {
	logger              zerolog.Logger
	alwaysIncludeFields map[string]struct{}
	Options
}

func New(w io.Writer, o Options) *Printer {
	output := zerolog.ConsoleWriter{Out: w, TimeFormat: o.TimestampFormat}
	o.includes = map[string]struct{}{}
	for _, include := range o.Includes {
		o.includes[include] = struct{}{}
	}
	return &Printer{
		logger:  zerolog.New(output),
		Options: o,
	}
}

func (p *Printer) Print(e event.Event) {

	if err, ok := e.(event.Error); ok {
		if !p.Options.Quiet {
			p.logger.Err(err).Send()
		}
	} else if typed, ok := e.(event.InvalidFormat); ok {
		if !p.Options.Quiet && p.Options.PrintInvalidFormat {
			p.logger.Info().Msg(string(typed))
		}
	} else {
		for _, match := range p.Matches {
			if !match.Matches(e) {
				return
			}
		}
		if eMap, ok := e.(event.Map); ok {
			p.setAlwaysIncludeFields(eMap)
			if len(p.Options.includes) > 0 {
				toDelete := []string{}
				for k := range eMap {
					if _, exists := p.Options.includes[k]; !exists {
						toDelete = append(toDelete, k)
					}
				}
				for _, exclude := range toDelete {
					if _, exists := p.alwaysIncludeFields[exclude]; exists {
						continue
					}
					delete(eMap, exclude)
				}
			}
			for _, exclude := range p.Options.Excludes {
				if _, exists := p.alwaysIncludeFields[exclude]; exists {
					continue
				}
				delete(eMap, exclude)
			}
		}
		p.logger.Log().EmbedObject(e).Send()
	}
}

func (p *Printer) setAlwaysIncludeFields(e event.Map) {
	if len(p.alwaysIncludeFields) != 0 {
		return
	}

	zerolog.TimestampFieldName = getTimestampField(e)
	zerolog.LevelFieldName = getLevelField(e)
	zerolog.MessageFieldName = getMessageField(e)
	p.alwaysIncludeFields = map[string]struct{}{
		zerolog.TimestampFieldName: {},
		zerolog.LevelFieldName:     {},
		zerolog.MessageFieldName:   {},
	}
}

func getTimestampField(e event.Map) string {
	if zerolog.TimestampFieldName != "auto" {
		return zerolog.TimestampFieldName
	}

	for _, tsFieldName := range timestampFields {
		if _, ok := e[tsFieldName]; ok {
			return tsFieldName
		}
	}

	return timestampFields[0]
}

func getLevelField(e event.Map) string {
	if zerolog.LevelFieldName != "auto" {
		return zerolog.LevelFieldName
	}

	for _, tsFieldName := range levelFields {
		if _, ok := e[tsFieldName]; ok {
			return tsFieldName
		}
	}

	return levelFields[0]
}

func getMessageField(e event.Map) string {
	if zerolog.MessageFieldName != "auto" {
		return zerolog.MessageFieldName
	}

	for _, tsFieldName := range messageFields {
		if _, ok := e[tsFieldName]; ok {
			return tsFieldName
		}
	}

	return messageFields[0]
}
