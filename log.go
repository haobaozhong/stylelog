package stylelog

import (
	"fmt"
	"log"
)

type StyleLogger struct {
	*log.Logger
}

// Set Style
func (sl *StyleLogger) SetStyle(styleCodes ...int) {
	seq := ""
	for _, code := range styleCodes {
		seq += fmt.Sprintf("%d;", code)
	}
	fmt.Printf("%s%d;%s%s", CSI, Reset, seq, Normal)
}

// Reset / Normal
func (sl *StyleLogger) Reset() {
	fmt.Printf("%s%d%s", CSI, Reset, Normal)
}

func New(logger *log.Logger) *StyleLogger {
	return &StyleLogger{
		Logger: logger,
	}
}
