package terminal

import (
	"os"

	"golang.org/x/term"
)

var terminal = Terminal{}

type Terminal struct {
	Height int
	Width  int
}

func (t *Terminal) GetDimensions() error {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	t.Height = height
	t.Width = width

	return nil

}
