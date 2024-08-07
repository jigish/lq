# lq
simple command line tool to pretty print and filter logs

```
lq is a simple log pretty printer and filter

Usage:
  lq [flags]

Flags:
  -e, --exclude strings                           do not print these extra fields
  -f, --format string                             incoming log line format: only 'json' is supported as of now (default "json")
  -h, --help                                      help for lq
  -i, --include strings                           only print these extra fields
      --level-field string                        level field name (default "level")
  -m, --match strings                             print log lines only when the given field string regex matches the given value. format: field=value
      --match-duration strings                    print log lines only when the given field duration equals the given value. format: field=value
      --match-duration-greater strings            print log lines only when the given field duration is greater than the given value. format: field>value
      --match-duration-greater-or-equal strings   print log lines only when the given field duration is greater than or equals the given value. format: field>=value
      --match-duration-less strings               print log lines only when the given field duration is less than the given value. format: field<value
      --match-duration-less-or-equal strings      print log lines only when the given field duration is less than or equals the given value. format: field<=value
      --match-float strings                       print log lines only when the given field float equals the given value. format: field=value
      --match-float-greater strings               print log lines only when the given field float is greater than the given value. format: field>value
      --match-float-greater-or-equal strings      print log lines only when the given field float is greater than or equals the given value. format: field>=value
      --match-float-less strings                  print log lines only when the given field float is less than the given value. format: field<value
      --match-float-less-or-equal strings         print log lines only when the given field float is less than or equals the given value. format: field<=value
      --match-int strings                         print log lines only when the given field int equals the given value. format: field=value
      --match-int-greater strings                 print log lines only when the given field int is greater than the given value. format: field>value
      --match-int-greater-or-equal strings        print log lines only when the given field int is greater than or equals the given value. format: field>=value
      --match-int-less strings                    print log lines only when the given field int is less than the given value. format: field<value
      --match-int-less-or-equal strings           print log lines only when the given field int is less than or equals the given value. format: field<=value
      --match-time strings                        print log lines only when the given field time equals the given value. format: field=value
      --match-time-after strings                  print log lines only when the given field time is after the given value. format: field>value
      --match-time-after-or-equal strings         print log lines only when the given field time is after or equals the given value. format: field>=value
      --match-time-before strings                 print log lines only when the given field time is before the given value. format: field<value
      --match-time-before-or-equal strings        print log lines only when the given field time is before or equals the given value. format: field<=value
      --match-time-format string                  time format for --match-time* (see https://pkg.go.dev/time#pkg-constants) (default "2006-01-02T15:04:05.000Z07:00")
      --message-field string                      message field name (default "msg")
      --print-invalid-format                      print lines with invalid format
  -q, --quiet                                     quiet all internal error logging
      --timestamp-field string                    timestamp field name (default "ts")
      --timestamp-format string                   timestamp format (see https://pkg.go.dev/time#pkg-constants) (default "2006-01-02T15:04:05.000Z07:00")
```


### Credits

`lq` is inspired by [`humanlog`](https://github.com/humanlogio/humanlog) and leverages [`zerolog`](https://github.com/rs/zerolog) for pretty printing.
