package printer

import "github.com/jigish/lq/pkg/match"

type Options struct {
	Quiet              bool
	PrintInvalidFormat bool
	TimestampFormat    string

	Matches []match.Matcher

	Includes []string
	includes map[string]struct{}
	Excludes []string
}
