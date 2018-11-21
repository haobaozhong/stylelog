package stylelog

import (
	"fmt"
	"log"
)

const ()

type StyleLogger struct {
	csiSequences string // CSI sequences
	*log.Logger
}

func (sl *StyleLogger) Danger(v ...interface{}) {
	sl.Style(Red)
	sl.Logger.Print(v...)
	sl.Reset()
}

func (sl *StyleLogger) Dangerln(v ...interface{}) {
	sl.Style(Red)
	sl.Logger.Println(v...)
	sl.Reset()
}

// Set Style
func (sl *StyleLogger) Style(styleCodes ...int) {
	seq := ""
	for _, code := range styleCodes {
		seq += fmt.Sprintf("%d;", code)
	}
	csiSequences := fmt.Sprintf("%s%d;%s%s", CSI, Reset, seq, Normal)
	fmt.Print(csiSequences)
	// fmt.Printf("%s%d;%d;%d;%d;%s", CSI, Reset, Yellow, Bold, Underline, Normal)
	sl.setCSISequences(csiSequences)
}

// Style And Reset
func (sl *StyleLogger) StyleAndReset(styleCodes ...int) {
	sl.Style(styleCodes...)
	sl.Reset()
}

// Get CSI sequences
func (sl *StyleLogger) CSISequences() string {
	return sl.csiSequences
}

// set CSI sequences
func (sl *StyleLogger) setCSISequences(csi string) {
	sl.csiSequences = csi
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
