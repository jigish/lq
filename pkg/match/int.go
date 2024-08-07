package match

import (
	"strconv"

	"github.com/jigish/lq/pkg/event"
)

func init() {
	definitions = append(definitions, &definition{
		new:         newIntMatcher,
		name:        "match-int",
		description: "int equals",
	}, &definition{
		new:         newIntGreaterThanMatcher,
		name:        "match-int-greater",
		description: "int is greater than",
		separator:   ">",
	}, &definition{
		new:         newIntGreaterThanOrEqualsMatcher,
		name:        "match-int-greater-or-equal",
		description: "int is greater than or equals",
		separator:   ">=",
	}, &definition{
		new:         newIntLessThanMatcher,
		name:        "match-int-less",
		description: "int is less than",
		separator:   "<",
	}, &definition{
		new:         newIntLessThanOrEqualsMatcher,
		name:        "match-int-less-or-equal",
		description: "int is less than or equals",
		separator:   "<=",
	})
}

type intMatcher struct {
	Field string
	Value int64
}

func newIntMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseInt(valueStr, 10, 64)
	return &intMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *intMatcher) Matches(e event.Event) bool {
	value, ok := extractInt64(e, f.Field)
	return ok && f.Value == value
}

type intGreaterThanMatcher struct {
	Field string
	Value int64
}

func newIntGreaterThanMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseInt(valueStr, 10, 64)
	return &intGreaterThanMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *intGreaterThanMatcher) Matches(e event.Event) bool {
	value, ok := extractInt64(e, f.Field)
	return ok && f.Value > value
}

type intGreaterThanOrEqualsMatcher struct {
	Field string
	Value int64
}

func newIntGreaterThanOrEqualsMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseInt(valueStr, 10, 64)
	return &intGreaterThanOrEqualsMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *intGreaterThanOrEqualsMatcher) Matches(e event.Event) bool {
	value, ok := extractInt64(e, f.Field)
	return ok && f.Value >= value
}

type intLessThanMatcher struct {
	Field string
	Value int64
}

func newIntLessThanMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseInt(valueStr, 10, 64)
	return &intLessThanMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *intLessThanMatcher) Matches(e event.Event) bool {
	value, ok := extractInt64(e, f.Field)
	return ok && f.Value < value
}

type intLessThanOrEqualsMatcher struct {
	Field string
	Value int64
}

func newIntLessThanOrEqualsMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseInt(valueStr, 10, 64)
	return &intLessThanOrEqualsMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *intLessThanOrEqualsMatcher) Matches(e event.Event) bool {
	value, ok := extractInt64(e, f.Field)
	return ok && f.Value <= value
}

func extractInt64(e event.Event, field string) (int64, bool) {
	untypedField, exists := extractField(e, field)
	if !exists {
		return 0, false
	}
	switch field := untypedField.(type) {
	case int:
		return int64(field), true
	case int8:
		return int64(field), true
	case int16:
		return int64(field), true
	case int32:
		return int64(field), true
	case int64:
		return field, true
	case uint:
		return int64(field), true
	case uint8:
		return int64(field), true
	case uint16:
		return int64(field), true
	case uint32:
		return int64(field), true
	default:
		return 0, false
	}
}
