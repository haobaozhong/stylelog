// https://en.wikipedia.org/wiki/ANSI_escape_code
package stylelog

const (
	// ESC (27 / hex 0x1B / octal 033)
	ESC = "\033"

	// CSI - Control Sequence Introducer
	CSI = "\033["
)

// style
const (
	// TODO: add more
	Reset  = 0 // Reset / Normal	all attributes off
	Normal = "m"

	Bold       = 1 // Bold or increased intensity
	Faint      = 2 // Faint (decreased intensity)
	Italic     = 3 // Italic	Not widely supported. Sometimes treated as inverse.
	Underline  = 4 // Underline
	SlowBlink  = 5 // Slow Blink	less than 150 per minute
	RapidBlink = 6 // Rapid Blink	MS-DOS ANSI.SYS; 150+ per minute; not widely supported

	// foreground
	Black   = 30
	Red     = 31
	Green   = 32
	Yellow  = 33
	Blue    = 34
	Magenta = 35
	Cyan    = 36
	White   = 37

	// background
	BlackBG   = 40
	RedBG     = 41
	GreenBG   = 42
	YellowBG  = 43
	BlueBG    = 44
	MagentaBG = 45
	CyanBG    = 46
	WhiteBG   = 47
)
