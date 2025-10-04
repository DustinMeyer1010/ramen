package ramen

import (
	"github.com/DustinMeyer1010/ramen/keys"
	"github.com/DustinMeyer1010/ramen/terminal"
)

type ExitCode int

const (
	QUIT ExitCode = iota
	SELECT
	UNKNOWN
)

var CURSOR = terminal.TERMINAL.Cursor
var KEYSCHANNEL = keys.NewKeyChannel()
var RESIZECHANNEL = terminal.NewResizeChannel()

func startEventSources() {
	go terminal.ResizeCheck(RESIZECHANNEL)
	go keys.KeyChannel(KEYSCHANNEL)
}
