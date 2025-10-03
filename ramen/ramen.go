package ramen

import (
	"log"
	"os"

	"github.com/DustinMeyer1010/ramen/terminal"
	"golang.org/x/term"
)

var cur = terminal.NewCursor(0, 0)
var rawMode = -1
var PrevTermState *term.State

func startRawMode() {
	var err error
	PrevTermState, err = term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		log.Fatal(err)
	}

}

func Exit() {
	term.Restore(int(os.Stdin.Fd()), PrevTermState)
}
