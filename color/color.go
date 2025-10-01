package color

type colorCode struct {
	FG int
	BG int
}

type textCode struct {
	Color colorCode
	font  font
}

type font int

type color string

func NewColorCode(fg, bg color) colorCode {
	return colorCode{FG: colors[fg].FG, BG: colors[bg].BG}
}

func NewTextCode(color colorCode, font font) textCode {
	return textCode{Color: color, font: font}
}

// Colors are \033[<n>m
// <n> = color code
var colors = map[color]colorCode{
	"Black":          {FG: 30, BG: 40},
	"RED":            {FG: 31, BG: 41},
	"GREEN":          {FG: 32, BG: 42},
	"YELLOW":         {FG: 33, BG: 43},
	"BLUE":           {FG: 34, BG: 44},
	"MAGENTA":        {FG: 35, BG: 45},
	"CYAN":           {FG: 36, BG: 46},
	"WHITE":          {FG: 37, BG: 47},
	"BRIGHT BLACK":   {FG: 90, BG: 100},
	"BRIGHT RED":     {FG: 91, BG: 101},
	"BRIGHT GREEN":   {FG: 92, BG: 102},
	"BRIGHT YELLOW":  {FG: 93, BG: 103},
	"BRIGHT BLUE":    {FG: 94, BG: 104},
	"BRIGHT MAGENTA": {FG: 95, BG: 105},
	"BRIGHT CYAN":    {FG: 96, BG: 106},
	"BRIGHT WHITE":   {FG: 97, BG: 107},
}

const (
	BLACK          color = "BLACK"
	RED            color = "RED"
	GREEN          color = "GREEN"
	YELLOW         color = "YELLOW"
	BLUE           color = "BLUE"
	MAGENTA        color = "MAGENTA"
	CYAN           color = "CYAN"
	WHITE          color = "WHITE"
	BRIGHT_BLACK   color = "BRIGHT BLACK"
	BRIGHT_RED     color = "BRIGHT RED"
	BRIGHT_GREEN   color = "BRIGHT GREEN"
	BRIGHT_BLUE    color = "BRIGHT BLUE"
	BRIGHT_MAGENTA color = "BRIGHT MAGENTA"
	BRIGHT_CYAN    color = "BRIGHT CYAN"
	BRIGHT_WHITE   color = "BRIGHT WHITE"
)

// Font are \033[f;cm
// f = font
// c = color
const (
	BOLD          font = 1
	DIM           font = 2
	ITALIC        font = 3
	UNDERLINE     font = 4
	BLINKSLOW     font = 5
	BLINKFAST     font = 6
	REVERSE       font = 7
	HIDDEN        font = 8
	STRIKETHROUGH font = 9
)

const RESET int = 0
