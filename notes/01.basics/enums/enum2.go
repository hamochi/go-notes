package enums

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type NamedColor Color

func (n NamedColor) String() string {
	switch n {
	case Yellow:
		return "yellow"
	case Cyan:
		return "cyan"
	case Magenta:
		return "magenta"
	case Silver:
		return "silver"
	}
	return "unknown color"
}

var (
	Yellow  = NamedColor{255, 255, 0}
	Cyan    = NamedColor{0, 255, 255}
	Magenta = NamedColor{255, 0, 255}
	Silver  = NamedColor{192, 192, 192}
)

func ReturnMagenta() NamedColor {
	return Magenta
}
