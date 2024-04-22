package copy

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type LoggerFactory struct {
	StartTime string
	Status    int
	Duration  time.Duration
	Hostname  string
	Method    string
	Path      string
	Request   *http.Request
}

var LoggerDefaultFormat = "{{.StartTime}} | {{.Status}} | \t {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}}"

var LoggerDefaultDateFormat = time.RFC3339

type ALogger interface {
	PrintLn(v ...interface{})
	Printf(format string, v ...interface{})
}

type Logger struct {
	ALogger
	dateFormat string
	template   *template.Template
}

func NewLogger() *Logger {
	logger := &Logger{
		ALogger:    log.New(os.Stdout, "[negroni]", 0),
		dateFormat: LoggerDefaultDateFormat,
	}
	logger.SetFormat(LoggerDefaultFormat)
	return logger
}

func (l Logger) SetFormat(format string) {
	l.template = template.Must(template.New("negroni_parse").Parse(format))
}

func (l Logger) SetDateFormat(format string) {
	l.dateFormat = format
}

func (l Logger) ServerHttp(rw http.ResponseWriter, r *http.Request,next http.HandlerFunc)  {
	start := time.Now()
	next(rw, r)

	res := rw.()
}
