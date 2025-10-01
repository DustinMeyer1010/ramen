package cursor

import (
	"fmt"

	"github.com/DustinMeyer1010/Ramen/keys"
)

type cursor struct {
	x int
	y int
}

type cursorCode int

// \033[n q
// n = number for cursor code
const (
	BLINKING           cursorCode = 1
	STEADY             cursorCode = 2
	BLINKING_UNDERLINE cursorCode = 3
	UNDERLINE          cursorCode = 4
	BLINKING_BAR       cursorCode = 5
	STEADY_BAR         cursorCode = 6
	HIDE               string     = "\033[?25l"
	SHOW               string     = "\033[?25h"
)

var print = fmt.Print

// Create cursor on the screen
func NewCursor(x int, y int) cursor {
	fmt.Printf("\033[%d q", STEADY)
	c := cursor{x: x, y: y}
	c.Origin()
	c.DrawContainer()

	return c
}

// Moves cursor Down n times
//
// Cursor will not move outside of configured terminal height
func (c *cursor) Down(n int) {
	for range n {
		if TerminalConfig.Height <= c.y {
			return
		}
		c.y += 1
		print(keys.DownArrow)
	}
}

// Moves cursor Up n times
//
// Cursor will not move ouside of configured terminal height
func (c *cursor) Up(n int) {
	for range n {
		if 0 >= c.y {
			return
		}
		c.y -= 1
		print(keys.UpArrow)
	}
}

// Move cursor Right n times
//
// Cursor will not move ouside of configured terminal width
func (c *cursor) Right(n int) {

	for range n {

		if TerminalConfig.Width <= c.x {
			return
		}
		c.x += 1
		print(keys.RightArrow)
	}
}

// Move cursor Left n times
//
// Cursor will not move ouside of configured terminal width
func (c *cursor) Left(n int) {
	for range n {
		if 0 >= c.x {
			return
		}
		c.x -= 1
		print(keys.LeftArrow)
	}
}

// Move cursor back to 0,0
func (c *cursor) Origin() {
	c.Up(TerminalConfig.Height)
	c.Left(TerminalConfig.Width)
}

// Move cursor to x, y cordinate
//
// If coordinates outside configured terminal height and width
// cursor will be put at edge/max height or width
func (c *cursor) MoveTo(x, y int) {
	c.Origin()
	c.Right(x)

	c.Down(y)
}

// Reset entire terminal to blank
func (c *cursor) ClearTerminal() {
	x := c.x
	y := c.y
	c.Origin()
	for range TerminalConfig.Height {
		print(keys.ClearLine)
		c.Down(1)
	}
	c.MoveTo(x, y)
}

// Draws the height and width of terminal configuration
func (c *cursor) DrawContainer() {
	c.Origin()
	c.Down(TerminalConfig.Height)
	c.Right(TerminalConfig.Height)
	c.Origin()
}

// Draws text from location of cursor
//
// Text will not wrap any text greater than terminal width will throw error
func (c *cursor) DrawText(text string) error {
	if c.x+len(text) > TerminalConfig.Width {
		return fmt.Errorf("text too long for width of page")
	}

	print(text)
	c.x += len(text)

	return nil
}

// Returns the x coordinate of cursor
func (c *cursor) GetX() int {
	return c.x
}

// Returns the y coordinate of curosor
func (c *cursor) GetY() int {
	return c.y
}

// Clears the entire line the cursor is on
func (c *cursor) ClearLine() {
	print(keys.ClearLine)
}

func (c *cursor) Hide() {
	print(HIDE)
}

func (c *cursor) Show() {
	print(SHOW)
}
