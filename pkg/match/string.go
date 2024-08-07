package match

import (
	"regexp"

	"github.com/jigish/lq/pkg/event"
)

func init() {
	definitions = append(definitions, &definition{
		new:         newStringMatcher,
		name:        "match",
		short:       "m",
		description: "string regex matches",
	})
}

type stringMatcher struct {
	Field  string
	Regexp *regexp.Regexp
}

func newStringMatcher(field, value string) (Matcher, error) {
	r, err := regexp.Compile(value)
	return &stringMatcher{
		Field:  field,
		Regexp: r,
	}, err
}

func (f *stringMatcher) Matches(e event.Event) bool {
	value, ok := extractString(e, f.Field)
	return ok && f.Regexp.MatchString(value)
}

func extractString(e event.Event, field string) (string, bool) {
	untypedField, exists := extractField(e, field)
	if !exists {
		return "", false
	}
	field, ok := untypedField.(string)
	return field, ok
}
