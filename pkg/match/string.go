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
		description: "string equals",
	}, &definition{
		new:         newStringRegexMatcher,
		name:        "match-regex",
		short:       "M",
		description: "string regex matches",
	})
}

type stringMatcher struct {
	Field string
	Value string
}

func newStringMatcher(field, value string) (Matcher, error) {
	return &stringMatcher{
		Field: field,
		Value: value,
	}, nil
}

func (f *stringMatcher) Matches(e event.Event) bool {
	value, ok := extractString(e, f.Field)
	return ok && f.Value == value
}

type stringRegexMatcher struct {
	Field  string
	Regexp *regexp.Regexp
}

func newStringRegexMatcher(field, value string) (Matcher, error) {
	r, err := regexp.Compile(value)
	return &stringRegexMatcher{
		Field:  field,
		Regexp: r,
	}, err
}

func (f *stringRegexMatcher) Matches(e event.Event) bool {
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
