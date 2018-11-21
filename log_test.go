package stylelog

import (
	"log"
	"os"
	"testing"
)

func TestColorLogger(t *testing.T) {
	baseLogger := log.New(os.Stdout, "STYLE LOGGER: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	logger := New(baseLogger)

	logger.Style(Bold, Underline, Green, Faint, SlowBlink)
	logger.Print("hello,")
	logger.Print("world")
	logger.Reset()
	baseLogger.Print("hello,")
	baseLogger.Print("world,")

	logger.Danger("hello, danger")
	logger.Dangerln("hello, danger")
}
