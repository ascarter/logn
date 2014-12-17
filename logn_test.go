package logn

import (
	"bytes"
	"regexp"
	"testing"
)

const (
	pattern = `(\S+)\s+(\d{4}[/]\d{2}[/]\d{2}\s\d{2}:\d{2}:\d{2})\s(.*)`
)

var validLogMsg = regexp.MustCompile(pattern)

var examples = []struct {
	level string
	msg   string
}{
	{"INFO", "info test message"},
	{"WARN", "warning test message"},
	{"ERROR", "error test message"},
}

func checkLogMessages(t *testing.T, result string) {
	matches := validLogMsg.FindAllStringSubmatch(result, -1)
	for i, v := range matches {
		lm := examples[i]
		if v[1] != lm.level {
			t.Errorf("%v != %v", v[1], lm.level)
		}
		if v[3] != lm.msg {
			t.Errorf("%v != %v", v[3], lm.msg)
		}
	}
}

func TestLog(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	for _, v := range examples {
		switch v.level {
		default:
			t.Logf("Unknown level: %s\n", v.level)
		case "INFO":
			logger.Info(v.msg)
		case "WARN":
			logger.Warning(v.msg)
		case "ERROR":
			logger.Error(v.msg)
		}

	}
	checkLogMessages(t, buf.String())
}

func TestLogf(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	for _, v := range examples {
		switch v.level {
		default:
			t.Logf("Unknown level: %s\n", v.level)
		case "INFO":
			logger.Infof("%s", v.msg)
		case "WARN":
			logger.Warningf("%s", v.msg)
		case "ERROR":
			logger.Errorf("%s", v.msg)
		}

	}
	checkLogMessages(t, buf.String())
}

func TestLogln(t *testing.T) {
	var buf bytes.Buffer
	logger := New(&buf)
	for _, v := range examples {
		switch v.level {
		default:
			t.Logf("Unknown level: %s\n", v.level)
		case "INFO":
			logger.Infoln(v.msg)
		case "WARN":
			logger.Warningln(v.msg)
		case "ERROR":
			logger.Errorln(v.msg)
		}

	}
	checkLogMessages(t, buf.String())
}

func TestDefaultLog(t *testing.T) {
	var buf bytes.Buffer
	std = New(&buf)
	for _, v := range examples {
		switch v.level {
		default:
			t.Logf("Unknown level: %s\n", v.level)
		case "INFO":
			Info(v.msg)
		case "WARN":
			Warning(v.msg)
		case "ERROR":
			Error(v.msg)
		}

	}
	checkLogMessages(t, buf.String())
}

func TestDefaultLogf(t *testing.T) {
	var buf bytes.Buffer
	std = New(&buf)
	for _, v := range examples {
		switch v.level {
		default:
			t.Logf("Unknown level: %s\n", v.level)
		case "INFO":
			Infof("%s", v.msg)
		case "WARN":
			Warningf("%s", v.msg)
		case "ERROR":
			Errorf("%s", v.msg)
		}

	}
	checkLogMessages(t, buf.String())
}

func TestDefaultLogln(t *testing.T) {
	var buf bytes.Buffer
	std = New(&buf)
	for _, v := range examples {
		switch v.level {
		default:
			t.Logf("Unknown level: %s\n", v.level)
		case "INFO":
			Infoln(v.msg)
		case "WARN":
			Warningln(v.msg)
		case "ERROR":
			Errorln(v.msg)
		}

	}
	checkLogMessages(t, buf.String())
}
