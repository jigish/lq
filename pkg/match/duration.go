package match

import (
	"time"

	"github.com/jigish/lq/pkg/event"
)

func init() {
	definitions = append(definitions, &definition{
		new:         newDurationMatcher,
		name:        "match-duration",
		description: "duration equals",
	}, &definition{
		new:         newDurationGreaterThanMatcher,
		name:        "match-duration-greater",
		description: "duration is greater than",
		separator:   ">",
	}, &definition{
		new:         newDurationGreaterThanOrEqualsMatcher,
		name:        "match-duration-greater-or-equal",
		description: "duration is greater than or equals",
		separator:   ">=",
	}, &definition{
		new:         newDurationLessThanMatcher,
		name:        "match-duration-less",
		description: "duration is less than",
		separator:   "<",
	}, &definition{
		new:         newDurationLessThanOrEqualsMatcher,
		name:        "match-duration-less-or-equal",
		description: "duration is less than or equals",
		separator:   "<=",
	})
}

type durationMatcher struct {
	Field string
	Value time.Duration
}

func newDurationMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.ParseDuration(valueStr)
	return &durationMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *durationMatcher) Matches(e event.Event) bool {
	value, ok := extractDuration(e, f.Field)
	return ok && f.Value == value
}

type durationGreaterThanMatcher struct {
	Field string
	Value time.Duration
}

func newDurationGreaterThanMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.ParseDuration(valueStr)
	return &durationGreaterThanMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *durationGreaterThanMatcher) Matches(e event.Event) bool {
	value, ok := extractDuration(e, f.Field)
	return ok && f.Value > value
}

type durationGreaterThanOrEqualsMatcher struct {
	Field string
	Value time.Duration
}

func newDurationGreaterThanOrEqualsMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.ParseDuration(valueStr)
	return &durationGreaterThanOrEqualsMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *durationGreaterThanOrEqualsMatcher) Matches(e event.Event) bool {
	value, ok := extractDuration(e, f.Field)
	return ok && f.Value >= value
}

type durationLessThanMatcher struct {
	Field string
	Value time.Duration
}

func newDurationLessThanMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.ParseDuration(valueStr)
	return &durationLessThanMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *durationLessThanMatcher) Matches(e event.Event) bool {
	value, ok := extractDuration(e, f.Field)
	return ok && f.Value < value
}

type durationLessThanOrEqualsMatcher struct {
	Field string
	Value time.Duration
}

func newDurationLessThanOrEqualsMatcher(field, valueStr string) (Matcher, error) {
	value, err := time.ParseDuration(valueStr)
	return &durationLessThanOrEqualsMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *durationLessThanOrEqualsMatcher) Matches(e event.Event) bool {
	value, ok := extractDuration(e, f.Field)
	return ok && f.Value <= value
}

func extractDuration(e event.Event, field string) (time.Duration, bool) {
	untypedField, exists := extractField(e, field)
	if !exists {
		return 0, false
	}
	switch field := untypedField.(type) {
	case int:
		return time.Duration(field), true
	case int8:
		return time.Duration(field), true
	case int16:
		return time.Duration(field), true
	case int32:
		return time.Duration(field), true
	case int64:
		return time.Duration(field), true
	case string:
		fieldDuration, err := time.ParseDuration(field)
		return fieldDuration, err != nil
	case time.Duration:
		return field, true
	default:
		return 0, false
	}
}
