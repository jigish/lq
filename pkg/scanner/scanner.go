package scanner

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/kr/logfmt"

	"github.com/jigish/lq/pkg/event"
)

type Scanner struct {
	scanner *bufio.Scanner
	Options

	c chan<- event.Event // internal channel can only be sent to
	C <-chan event.Event // exported channel can only be read from
}

func New(r io.Reader, o Options) *Scanner {
	c := make(chan event.Event)
	return &Scanner{
		scanner: bufio.NewScanner(r),
		Options: o,
		c:       c,
		C:       c,
	}
}

func (s *Scanner) Scan(ctx context.Context) {
	go s.doScan(ctx)
}

func (s *Scanner) sendLine(line []byte) {
	switch s.Format {
	case FormatAuto:
		if isJSON(string(line)) {
			s.sendJSON(line)
		} else {
			s.sendLogFmt(line)
		}
	case FormatLogFmt:
		s.sendLogFmt(line)
	case FormatJSON:
		s.sendJSON(line)
	default:
		s.c <- event.Error("unimplemented format: " + s.Options.Format)
		return
	}
}

func isJSON(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "{")
}

func (s *Scanner) sendJSON(line []byte) {
	var e event.Map
	if err := json.Unmarshal(line, &e); err != nil {
		s.c <- event.InvalidFormat(line)
		return
	}
	s.c <- e
}

func unmarshalLogFmtSafe(line []byte) (e event.Map, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("runtime panic: %v", x)
		}
	}()

	err = logfmt.Unmarshal(line, &e)
	return
}

func (s *Scanner) sendLogFmt(line []byte) {
	e, err := unmarshalLogFmtSafe(line)
	if err != nil {
		s.c <- event.InvalidFormat(line)
		return
	}
	s.c <- e
}

func (s *Scanner) doScan(ctx context.Context) {
	defer close(s.c)

	for s.scanner.Scan() {
		s.sendLine(s.scanner.Bytes())

		select {
		case <-ctx.Done():
			break
		default:
		}
	}
	if err := s.scanner.Err(); err != nil {
		s.c <- event.Error(err.Error())
	}
}
