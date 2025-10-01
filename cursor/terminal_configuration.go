package cursor

type TerminalCfg struct {
	Height int
	Width  int
}

var TerminalConfig = TerminalCfg{Height: 10, Width: 50}

func UpdateTerminalConfig(height, width int) {
	TerminalConfig = TerminalCfg{Height: height, Width: width}
}
