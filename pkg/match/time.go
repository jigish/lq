package match

import (
	"time"

	"github.com/jigish/lq/pkg/event"
)

var TimeFormat string

func init() {
	definitions = append(definitions, &definition{
		new:         newTimeMatcher,
		name:        "match-time",
		description: "time equals",
	}, &definition{
		new:         newTimeGreaterThanMatcher,
		name:        "match-time-after",
		description: "time is after",
		separator:   ">",
	}, &definition{
		new:         newTimeGreaterThanOrEqualsMatcher,
		name:        "match-time-after-or-equal",
		description: "time is after or equals",
		separator:   ">=",
	}, &definition{
		new:         newTimeLessThanMatcher,
		name:        "match-time-before",
		description: "time is before",
		separator:   "<",
	}, &definition{
		new:         newTimeLessThanOrEqualsMatcher,
		name:        "match-time-before-or-equal",
		description: "time is before or equals",
		separator:   "<=",
	})
}

type timeMatcher struct {
	Field string
	Value time.Time
}

func newTimeMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.Parse(TimeFormat, valueStr)
	return &timeMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *timeMatcher) Matches(e event.Event) bool {
	value, ok := extractTime(e, f.Field)
	return ok && f.Value.Equal(value)
}

type timeGreaterThanMatcher struct {
	Field string
	Value time.Time
}

func newTimeGreaterThanMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.Parse(TimeFormat, valueStr)
	return &timeGreaterThanMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *timeGreaterThanMatcher) Matches(e event.Event) bool {
	value, ok := extractTime(e, f.Field)
	return ok && f.Value.After(value)
}

type timeGreaterThanOrEqualsMatcher struct {
	Field string
	Value time.Time
}

func newTimeGreaterThanOrEqualsMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.Parse(TimeFormat, valueStr)
	return &timeGreaterThanOrEqualsMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *timeGreaterThanOrEqualsMatcher) Matches(e event.Event) bool {
	value, ok := extractTime(e, f.Field)
	return ok && f.Value.Equal(value) || f.Value.After(value)
}

type timeLessThanMatcher struct {
	Field string
	Value time.Time
}

func newTimeLessThanMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.Parse(TimeFormat, valueStr)
	return &timeLessThanMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *timeLessThanMatcher) Matches(e event.Event) bool {
	value, ok := extractTime(e, f.Field)
	return ok && f.Value.Before(value)
}

type timeLessThanOrEqualsMatcher struct {
	Field string
	Value time.Time
}

func newTimeLessThanOrEqualsMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.Parse(TimeFormat, valueStr)
	return &timeLessThanOrEqualsMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *timeLessThanOrEqualsMatcher) Matches(e event.Event) bool {
	value, ok := extractTime(e, f.Field)
	return ok && f.Value.Equal(value) || f.Value.Before(value)
}

func extractTime(e event.Event, field string) (time.Time, bool) {
	untypedField, exists := extractField(e, field)
	if !exists {
		return time.Time{}, false
	}
	switch field := untypedField.(type) {
	case string:
		fieldTime, err := time.Parse(TimeFormat, field)
		return fieldTime, err != nil
	case time.Time:
		return field, true
	default:
		return time.Time{}, false
	}
}
