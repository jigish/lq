package scanner

const (
	FormatAuto   = "auto"
	FormatJSON   = "json"
	FormatLogFmt = "logfmt"
)

type Options struct {
	Format string
}
