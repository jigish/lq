package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/jigish/lq/pkg/match"
	"github.com/jigish/lq/pkg/printer"
	"github.com/jigish/lq/pkg/scanner"
)

const rfc3339Milli = "2006-01-02T15:04:05.000Z07:00"

var rootCmd = &cobra.Command{
	Use:   "lq",
	Short: "lq is a simple log pretty printer and filter",
	Long:  "lq is a simple log pretty printer and filter. reads from stdin.",
	Run:   root,
}

var (
	sOpts = scanner.Options{}
	pOpts = printer.Options{}
)

func init() {
	// scanner options
	rootCmd.Flags().StringVarP(&sOpts.Format, "format", "f", scanner.FormatAuto,
		"incoming log line format: "+scanner.FormatJSON+", "+scanner.FormatLogFmt+", or "+scanner.FormatAuto)

	// printer options
	rootCmd.Flags().BoolVarP(&pOpts.Quiet, "quiet", "q", false, "quiet all internal error logging")
	rootCmd.Flags().BoolVar(&pOpts.PrintInvalidFormat, "print-invalid-format", false, "print lines with invalid format")
	rootCmd.Flags().StringSliceVarP(&pOpts.Includes, "include", "i", nil, "only print these extra fields")
	rootCmd.Flags().StringSliceVarP(&pOpts.Excludes, "exclude", "e", nil, "do not print these extra fields")
	rootCmd.Flags().StringVar(&pOpts.TimestampFormat, "timestamp-format", rfc3339Milli,
		"timestamp format (see https://pkg.go.dev/time#pkg-constants)")
	rootCmd.Flags().StringVar(&zerolog.TimestampFieldName, "timestamp-field", "auto", "timestamp field name")
	rootCmd.Flags().StringVar(&zerolog.LevelFieldName, "level-field", "auto", "level field name")
	rootCmd.Flags().StringVar(&zerolog.MessageFieldName, "message-field", "auto", "message field name")

	// match options
	rootCmd.Flags().StringVar(&match.TimeFormat, "match-time-format", rfc3339Milli,
		"time format for --match-time* (see https://pkg.go.dev/time#pkg-constants)")

	match.AddFlags(rootCmd)
}

func root(cmd *cobra.Command, args []string) {
	var err error
	pOpts.Matches, err = match.Parse()
	if err != nil {
		panic(err)
	}

	s := scanner.New(os.Stdin, sOpts)
	p := printer.New(os.Stdout, pOpts)

	s.Scan(cmd.Context())
	for e := range s.C {
		p.Print(e)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
