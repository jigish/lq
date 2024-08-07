package scanner

import (
	"bufio"
	"context"
	"encoding/json"
	"io"

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
	var e event.Map

	switch s.Format {
	case FormatJSON:
		if err := json.Unmarshal(line, &e); err != nil {
			s.c <- event.InvalidFormat(line)
			return
		}
	default:
		s.c <- event.Error("unimplemented format: " + s.Options.Format)
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
