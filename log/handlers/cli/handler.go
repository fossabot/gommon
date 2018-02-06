package cli

import (
	"io"
	"strconv"
	"time"

	"github.com/dyweb/gommon/log"
	"github.com/dyweb/gommon/util/color"
)

const (
	defaultTimeStampFormat = time.RFC3339
)

type Handler struct {
	w     io.Writer
	start time.Time
}

func New(w io.Writer) *Handler {
	return &Handler{
		w:     w,
		start: time.Now(),
	}
}

func (h *Handler) HandleLog(level log.Level, time time.Time, msg string) {
	b := formatHead(level, time, msg)
	b = append(b, '\n')
	h.w.Write(b)
}

func (h *Handler) HandleLogWithSource(level log.Level, time time.Time, msg string, source string) {
	b := formatHeadWithSource(level, time, msg, source)
	b = append(b, '\n')
	h.w.Write(b)
}

func (h *Handler) HandleLogWithFields(level log.Level, time time.Time, msg string, fields log.Fields) {
	// we use raw slice instead of bytes buffer because we need to use strconv.Append*, which requires raw slice
	b := formatHead(level, time, msg)
	b = append(b, ' ')
	b = formatFields(b, fields)
	b[len(b)-1] = '\n'
	h.w.Write(b)
}

func (h *Handler) HandleLogWithSourceFields(level log.Level, time time.Time, msg string, source string, fields log.Fields) {
	b := formatHeadWithSource(level, time, msg, source)
	b = append(b, ' ')
	b = formatFields(b, fields)
	b[len(b)-1] = '\n'
	h.w.Write(b)
}

func (h *Handler) Flush() {
	if s, ok := h.w.(log.Syncer); ok {
		s.Sync()
	}
}

// NOTE: most of format functions are same as (copied from) default handler

// no need to use fmt.Printf since we don't need any format
func formatHead(level log.Level, time time.Time, msg string) []byte {
	b := make([]byte, 0, 18+4+len(defaultTimeStampFormat)+len(msg))
	b = append(b, level.ColoredString()...)
	b = append(b, ' ')
	b = time.AppendFormat(b, defaultTimeStampFormat)
	b = append(b, ' ')
	b = append(b, msg...)
	return b
}

// we have a new function because source sits between time and msg in output, instead of after msg
// i.e. info 2018-02-04T21:03:20-08:00 main.go:18 show me the line
func formatHeadWithSource(level log.Level, time time.Time, msg string, source string) []byte {
	b := make([]byte, 0, 18+4+len(defaultTimeStampFormat)+len(msg))
	b = append(b, level.ColoredString()...)
	b = append(b, ' ')
	b = time.AppendFormat(b, defaultTimeStampFormat)
	b = append(b, ' ')
	b = append(b, color.CyanStart...)
	b = append(b, source...)
	b = append(b, color.End...)
	b = append(b, ' ')
	b = append(b, msg...)
	return b
}

// it has an extra tailing space, which can be updated inplace to a \n
func formatFields(b []byte, fields log.Fields) []byte {
	for _, f := range fields {
		b = append(b, color.CyanStart...)
		b = append(b, f.Key...)
		b = append(b, color.End...)
		b = append(b, '=')
		switch f.Type {
		case log.IntType:
			b = strconv.AppendInt(b, f.Int, 10)
		case log.StringType:
			b = append(b, f.Str...)
		}
		b = append(b, ' ')
	}
	return b
}
