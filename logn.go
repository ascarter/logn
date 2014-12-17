package logn

import (
	"io"
	"log"
	"runtime"
)

type Logger struct {
	t *log.Logger
	i *log.Logger
	w *log.Logger
	e *log.Logger
	f *log.Logger
	p *log.Logger
}

// New creates a new Logger. The out variable sets the destination to which
// log data will be written.
func New(out io.Writer) *Logger {
	return &Logger{
		t: log.New(out, "TRACE ", log.LstdFlags|log.Lshortfile),
		i: log.New(out, "INFO  ", log.LstdFlags),
		w: log.New(out, "WARN  ", log.LstdFlags),
		e: log.New(out, "ERROR ", log.LstdFlags),
		f: log.New(out, "FATAL ", log.LstdFlags),
		p: log.New(out, "PANIC ", log.LstdFlags),
	}
}

// Trace prints msg and then the current stack trace
func (l *Logger) Trace(msg string) {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	l.t.Printf("%s\n%s", msg, buf)
}

// Info prints informational message
func (l *Logger) Info(v ...interface{}) {
	l.i.Print(v...)
}

// Infof prints informational message as in fmt.Printf
func (l *Logger) Infof(fmt string, v ...interface{}) {
	l.i.Printf(fmt, v...)
}

// Infoln prints informational message as in fmt.Println
func (l *Logger) Infoln(v ...interface{}) {
	l.i.Println(v...)
}

// Warning prints warning message
func (l *Logger) Warning(v ...interface{}) {
	l.w.Print(v...)
}

// Warningf prints warning message as in fmt.Printf
func (l *Logger) Warningf(fmt string, v ...interface{}) {
	l.w.Printf(fmt, v...)
}

// Warningln prints warning message as in fmt.Println
func (l *Logger) Warningln(v ...interface{}) {
	l.w.Println(v...)
}

// Error prints error message
func (l *Logger) Error(v ...interface{}) {
	l.e.Print(v...)
}

// Errorf prints error message as in fmt.Printf
func (l *Logger) Errorf(fmt string, v ...interface{}) {
	l.e.Printf(fmt, v...)
}

// Errorln prints error message as in fmt.Println
func (l *Logger) Errorln(v ...interface{}) {
	l.e.Println(v...)
}

// Fatal prints warning message using log.Fatal
func (l *Logger) Fatal(v ...interface{}) {
	l.f.Fatal(v...)
}

// Fatalf prints warning message using log.Fatalf
func (l *Logger) Fatalf(fmt string, v ...interface{}) {
	l.f.Fatalf(fmt, v...)
}

// Fatalln prints warning message using log.Fatalln
func (l *Logger) Fatalln(v ...interface{}) {
	l.f.Fatalln(v...)
}

// Panic prints warning message using log.Panic
func (l *Logger) Panic(v ...interface{}) {
	l.p.Panic(v...)
}

// Panicf prints warning message using log.Panicf
func (l *Logger) Panicf(fmt string, v ...interface{}) {
	l.p.Panicf(fmt, v...)
}

// Panicln prints warning message using log.Panicln
func (l *Logger) Panicln(v ...interface{}) {
	l.p.Panicln(v...)
}
