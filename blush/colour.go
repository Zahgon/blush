package blush

// BgLevel is the colour value of R, G, or B when the colour is shown in the
// background.
const BgLevel = 70

// These are colour settings. NoRGB results in no colouring in the terminal.
var (
	NoRGB     = RGB{-1, -1, -1}
	FgRed     = RGB{255, 0, 0}
	FgBlue    = RGB{0, 0, 255}
	FgGreen   = RGB{0, 255, 0}
	FgBlack   = RGB{0, 0, 0}
	FgWhite   = RGB{255, 255, 255}
	FgCyan    = RGB{0, 255, 255}
	FgMagenta = RGB{255, 0, 255}
	FgYellow  = RGB{255, 255, 0}
	BgRed     = RGB{BgLevel, 0, 0}
	BgBlue    = RGB{0, 0, BgLevel}
	BgGreen   = RGB{0, BgLevel, 0}
	BgBlack   = RGB{0, 0, 0}
	BgWhite   = RGB{BgLevel, BgLevel, BgLevel}
	BgCyan    = RGB{0, BgLevel, BgLevel}
	BgMagenta = RGB{BgLevel, 0, BgLevel}
	BgYellow  = RGB{BgLevel, BgLevel, 0}
)

// Some stock colours. There will be no colouring when NoColour is used.
var (
	NoColour = Colour{NoRGB, NoRGB}
	Red      = Colour{FgRed, NoRGB}
	Blue     = Colour{FgBlue, NoRGB}
	Green    = Colour{FgGreen, NoRGB}
	Black    = Colour{FgBlack, NoRGB}
	White    = Colour{FgWhite, NoRGB}
	Cyan     = Colour{FgCyan, NoRGB}
	Magenta  = Colour{FgMagenta, NoRGB}
	Yellow   = Colour{FgYellow, NoRGB}
)

// DefaultColour is the default colour if no colour is set via arguments.
var DefaultColour = Blue

// RGB represents colours that can be printed in terminals. R, G and B should be
// between 0 and 255.
type RGB struct {
	R, G, B int
}

// Colour is a pair of RGB colours for foreground and background.
type Colour struct {
	Foreground RGB
	Background RGB
}

// Colourise wraps the input between colours.
func Colourise(input string, c Colour) string { _ = "STUB: not implemented"; return "" }

func foreground(c RGB) string { _ = "STUB: not implemented"; return "" }

func background(c RGB) string { _ = "STUB: not implemented"; return "" }

func unformat() string { _ = "STUB: not implemented"; return "" }

func colour(red, green, blue int) int { _ = "STUB: not implemented"; return 0 }

func baseColor(value, factor int) int { _ = "STUB: not implemented"; return 0 }

func colorFromArg(colour string) Colour { _ = "STUB: not implemented"; return *new(Colour) }

func colourGroup(colour string) Colour { _ = "STUB: not implemented"; return *new(Colour) }

func stockColour(colour string) Colour { _ = "STUB: not implemented"; return *new(Colour) }

// nolint:misspell // it's ok.

func hexColour(colour string) Colour { _ = "STUB: not implemented"; return *new(Colour) }

// getInt returns a number between 0-255 from a hex code. If the hex is not
// between 00 and ff, it returns -1.
func getInt(hex string) int { _ = "STUB: not implemented"; return 0 }
