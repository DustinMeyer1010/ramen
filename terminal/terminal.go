package terminal

import (
	"log"
	"os"

	"github.com/DustinMeyer1010/Ramen/cursor"
	"golang.org/x/term"
)

var cur = cursor.NewCursor(0, 0)
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
