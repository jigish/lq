package match

import (
	"strconv"

	"github.com/jigish/lq/pkg/event"
)

func init() {
	definitions = append(definitions, &definition{
		new:         newFloatMatcher,
		name:        "match-float",
		description: "float equals",
	}, &definition{
		new:         newFloatGreaterThanMatcher,
		name:        "match-float-greater",
		description: "float is greater than",
		separator:   ">",
	}, &definition{
		new:         newFloatGreaterThanOrEqualsMatcher,
		name:        "match-float-greater-or-equal",
		description: "float is greater than or equals",
		separator:   ">=",
	}, &definition{
		new:         newFloatLessThanMatcher,
		name:        "match-float-less",
		description: "float is less than",
		separator:   "<",
	}, &definition{
		new:         newFloatLessThanOrEqualsMatcher,
		name:        "match-float-less-or-equal",
		description: "float is less than or equals",
		separator:   "<=",
	})
}

type floatMatcher struct {
	Field string
	Value float64
}

func newFloatMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	return &floatMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *floatMatcher) Matches(e event.Event) bool {
	value, ok := extractFloat64(e, f.Field)
	return ok && f.Value == value
}

type floatGreaterThanMatcher struct {
	Field string
	Value float64
}

func newFloatGreaterThanMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	return &floatGreaterThanMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *floatGreaterThanMatcher) Matches(e event.Event) bool {
	value, ok := extractFloat64(e, f.Field)
	return ok && f.Value > value
}

type floatGreaterThanOrEqualsMatcher struct {
	Field string
	Value float64
}

func newFloatGreaterThanOrEqualsMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	return &floatGreaterThanOrEqualsMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *floatGreaterThanOrEqualsMatcher) Matches(e event.Event) bool {
	value, ok := extractFloat64(e, f.Field)
	return ok && f.Value >= value
}

type floatLessThanMatcher struct {
	Field string
	Value float64
}

func newFloatLessThanMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	return &floatLessThanMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *floatLessThanMatcher) Matches(e event.Event) bool {
	value, ok := extractFloat64(e, f.Field)
	return ok && f.Value < value
}

type floatLessThanOrEqualsMatcher struct {
	Field string
	Value float64
}

func newFloatLessThanOrEqualsMatcher(field, valueStr string) (Matcher, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	return &floatLessThanOrEqualsMatcher{
		Field: field,
		Value: value,
	}, err
}

func (f *floatLessThanOrEqualsMatcher) Matches(e event.Event) bool {
	value, ok := extractFloat64(e, f.Field)
	return ok && f.Value <= value
}

func extractFloat64(e event.Event, field string) (float64, bool) {
	untypedField, exists := extractField(e, field)
	if !exists {
		return 0, false
	}
	switch field := untypedField.(type) {
	case float32:
		return float64(field), true
	case float64:
		return field, true
	default:
		return 0, false
	}
}
