// Code generated by gommon from log/logger_generated.go.tmpl DO NOT EDIT.
package log

import (
	"fmt"
	"time"
)

func (l *Logger) IsTraceEnabled() bool {
	return l.level >= TraceLevel
}

func (l *Logger) Trace(args ...interface{}) {
	if l.level >= TraceLevel {
		l.h.HandleLog(TraceLevel, time.Now(), fmt.Sprint(args...))
	}
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	if l.level >= TraceLevel {
		l.h.HandleLog(TraceLevel, time.Now(), fmt.Sprintf(format, args...))
	}
}

func (l *Logger) TraceF(msg string, fields Fields) {
	if l.level >= TraceLevel {
		l.h.HandleLogWithFields(TraceLevel, time.Now(), msg, fields)
	}
}

func (l *Logger) IsDebugEnabled() bool {
	return l.level >= DebugLevel
}

func (l *Logger) Debug(args ...interface{}) {
	if l.level >= DebugLevel {
		l.h.HandleLog(DebugLevel, time.Now(), fmt.Sprint(args...))
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level >= DebugLevel {
		l.h.HandleLog(DebugLevel, time.Now(), fmt.Sprintf(format, args...))
	}
}

func (l *Logger) DebugF(msg string, fields Fields) {
	if l.level >= DebugLevel {
		l.h.HandleLogWithFields(DebugLevel, time.Now(), msg, fields)
	}
}

func (l *Logger) IsInfoEnabled() bool {
	return l.level >= InfoLevel
}

func (l *Logger) Info(args ...interface{}) {
	if l.level >= InfoLevel {
		l.h.HandleLog(InfoLevel, time.Now(), fmt.Sprint(args...))
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level >= InfoLevel {
		l.h.HandleLog(InfoLevel, time.Now(), fmt.Sprintf(format, args...))
	}
}

func (l *Logger) InfoF(msg string, fields Fields) {
	if l.level >= InfoLevel {
		l.h.HandleLogWithFields(InfoLevel, time.Now(), msg, fields)
	}
}

func (l *Logger) IsWarnEnabled() bool {
	return l.level >= WarnLevel
}

func (l *Logger) Warn(args ...interface{}) {
	if l.level >= WarnLevel {
		l.h.HandleLog(WarnLevel, time.Now(), fmt.Sprint(args...))
	}
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.level >= WarnLevel {
		l.h.HandleLog(WarnLevel, time.Now(), fmt.Sprintf(format, args...))
	}
}

func (l *Logger) WarnF(msg string, fields Fields) {
	if l.level >= WarnLevel {
		l.h.HandleLogWithFields(WarnLevel, time.Now(), msg, fields)
	}
}

func (l *Logger) IsErrorEnabled() bool {
	return l.level >= ErrorLevel
}

func (l *Logger) Error(args ...interface{}) {
	if l.level >= ErrorLevel {
		l.h.HandleLog(ErrorLevel, time.Now(), fmt.Sprint(args...))
	}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.level >= ErrorLevel {
		l.h.HandleLog(ErrorLevel, time.Now(), fmt.Sprintf(format, args...))
	}
}

func (l *Logger) ErrorF(msg string, fields Fields) {
	if l.level >= ErrorLevel {
		l.h.HandleLogWithFields(ErrorLevel, time.Now(), msg, fields)
	}
}
