// https://en.wikipedia.org/wiki/ANSI_escape_code
package stylelog

const (
	// ESC (27 / hex 0x1B / octal 033)
	ESC = "\033"

	// CSI - Control Sequence Introducer
	CSI = "\033["
)

// Styles
const (
	Reset  = 0 // Reset / Normal	all attributes off
	Normal = "m"

	Bold       = 1 // Bold or increased intensity
	Faint      = 2 // Faint (decreased intensity)
	Italic     = 3 // Italic	Not widely supported. Sometimes treated as inverse.
	Underline  = 4 // Underline
	SlowBlink  = 5 // Slow Blink	less than 150 per minute
	RapidBlink = 6 // Rapid Blink	MS-DOS ANSI.SYS; 150+ per minute; not widely supported

	ReverseVideo = 7  // reverse video	swap foreground and background colors
	Conceal      = 8  // Conceal	Not widely supported.
	CrossedOut   = 9  // Crossed-out	Characters legible, but marked for deletion.
	Primary      = 10 // Primary(default) font
	// 11–19	Alternative font	Select alternative font {\displaystyle n-10} {\displaystyle n-10}
	Fraktur                = 20 //	Fraktur	Rarely supported
	DoublyUnderline        = 21 // Doubly underline or Bold off	Double-underline per ECMA-48.[20] See discussion
	NormalColorOrIntensity = 22 // Normal color or intensity	Neither bold nor faint
	NotItalicNotFraktur    = 23 // Not italic, not Fraktur
	UnderlineOff           = 24 // Underline off	Not singly or doubly underlined
	BlinkOff               = 25 // Blink off
	InverseOff             = 27 // Inverse off
	RevealConcealOff       = 28 // Reveal conceal off
	NotCrossedOut          = 29 // Not crossed out
	// 30–37	Set foreground color
	SetForegroundColor     = 38 // Set foreground color
	DefaultForegroundColor = 39 // Default foreground color	implementation defined (according to standard)
	// 40–47	Set background color
	SetBackgroundColor     = 48 // Set background color
	DefaultBackgroundColor = 49 //	Default background color implementation defined (according to standard)
	Framed                 = 51 // Framed
	Encircled              = 52 // Encircled
	Overlined              = 53 // Overlined
	NotFramedOrEncircled   = 54 // Not framed or encircled
	NotOverlined           = 55 // Not overlined
)

// Colors
const (
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
