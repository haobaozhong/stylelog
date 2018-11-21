package stylelog

import (
	"log"
	"os"
	"testing"
)

func TestColorLogger(t *testing.T) {
	baseLogger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	baseLogger.Print("hello, world")

	logger := New(baseLogger)
	logger.SetStyle(Bold, Underline, Italic, Green, Faint)
	logger.Print("hello, world")
	logger.Reset()

	logger.Print("hello, world")
}
