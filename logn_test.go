package logn

import (
	"os"
	"testing"
)

func TestInfo(t *testing.T) {
	logger := New(os.Stderr)
	logger.Info("test")
}
