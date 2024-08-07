package match

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/jigish/lq/pkg/event"
)

const defaultSeparator = "="

var definitions = []*definition{}

type Matcher interface {
	Matches(event.Event) bool
}

type definition struct {
	new func(string, string) (Matcher, error)

	name        string
	short       string
	description string
	separator   string

	unparsed []string
}

func AddFlags(cmd *cobra.Command) {
	for _, definition := range definitions {
		if len(definition.separator) == 0 {
			definition.separator = defaultSeparator
		}
		cmd.Flags().StringSliceVarP(&definition.unparsed, definition.name, definition.short, []string{},
			"print log lines only when the given field "+definition.description+" the given value. format: field"+definition.separator+"value")
	}
}

func Parse() ([]Matcher, error) {
	matchers := []Matcher{}
	for _, definition := range definitions {
		for _, toParse := range definition.unparsed {
			parts := strings.Split(toParse, definition.separator)
			if len(parts) != 2 {
				return nil, errors.New("invalid match: " + definition.description + ": " + toParse)
			}
			matcher, err := definition.new(parts[0], parts[1])
			if err != nil {
				return nil, errors.Wrap(err, "invalid match: "+definition.description)
			}
			matchers = append(matchers, matcher)
		}
	}

	return matchers, nil
}

func extractField(e event.Event, f string) (any, bool) {
	typedE, ok := e.(event.Map)
	if !ok {
		return nil, false
	}

	untypedField, exists := typedE[f]
	if exists {
		return untypedField, true
	}

	// try nested
	parts := strings.Split(f, ".")
	var m map[string]any
	untypedField = e
	for _, part := range parts {
		m, ok = untypedField.(event.Map)
		if !ok {
			m, ok = untypedField.(map[string]any)
			if !ok {
				return nil, false
			}
		}
		untypedField, exists = m[part]
		if !exists {
			return nil, false
		}
	}
	return untypedField, true
}
