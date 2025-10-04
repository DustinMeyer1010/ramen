package terminal

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/term"
)

var TERMINAL terminal

type terminal struct {
	height        int
	width         int
	Cursor        cursor
	prevTermState *term.State
}

func NewTerminal(height, width int) {
	fmt.Print("\033[?1049h")
	h, w, err := GetDimensions()

	if err != nil {
		log.Fatal("failed to get terminal dimensions")
	}
	if height <= 0 {
		height = h
	}
	if width <= 0 {
		width = w
	}

	TERMINAL = terminal{height: height, width: width, prevTermState: rawMode(), Cursor: NewCursor()}
	TERMINAL.Cursor.Origin()
	TERMINAL.Cursor.ClearTerminal()
}

// Get the current dimensions of the terminal
func GetDimensions() (int, int, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		return 0, 0, err
	}

	return height, width, err

}

// Set the height and width of the current terminal
//
// If height and width exceed terminal dimension it will default to max dimensions
// Putting in negative values default to max dimensions
func (t *terminal) SetDimensions(height, width int) {
	t.height = height
	t.width = width
}

func rawMode() *term.State {
	prevTermState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		log.Fatal(err)
	}
	return prevTermState
}

func (t *terminal) Exit() {
	term.Restore(int(os.Stdin.Fd()), t.prevTermState)
}

func (t *terminal) GetHeight() int {
	return t.height
}

func (t *terminal) GetWidth() int {
	return t.width
}

func ResizeCheck(ch chan<- [2]int) {
	lastW, lastH, _ := term.GetSize(int(os.Stdout.Fd()))
	for {
		time.Sleep(time.Millisecond * 200)
		w, h, _ := term.GetSize(int(os.Stdout.Fd()))
		if w != lastW || h != lastH {
			ch <- [2]int{w, h} // send new size
			lastW, lastH = w, h
		}
	}
}

func NewResizeChannel() chan [2]int {
	return make(chan [2]int)
}
