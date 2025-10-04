package terminal

import (
	"fmt"

	"github.com/DustinMeyer1010/ramen/keys"
)

var print = fmt.Print

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
	CLEARLINE          string     = "\033[2K"
)

// Create cursor on the screen
func NewCursor() cursor {
	c := cursor{x: 0, y: 0}
	return c
}

// Moves cursor Down n times
//
// Cursor will not move outside of configured terminal height
func (c *cursor) Down(n int) {
	for range n {
		if TERMINAL.height <= c.y {
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

		if TERMINAL.width <= c.x {
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
	c.Up(TERMINAL.height)
	c.Left(TERMINAL.width)
	c.y = 0
	c.x = 0
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

// Move cursor to 0, Height cordinate
func (c *cursor) OriginBottom() {
	c.Origin()
	c.Down(TERMINAL.height)
}

// Reset entire terminal to blank
func (c *cursor) ClearTerminal() {
	c.Origin()
	for range TERMINAL.height + 1 {
		c.ClearLine()
		c.Down(1)
	}
}

// Draws the height and width of terminal configuration
func (c *cursor) DrawContainer() {
	c.Origin()
	for range TERMINAL.height {
		c.Down(TERMINAL.height)
	}
	for range TERMINAL.width {
		c.Right(TERMINAL.width)
	}
	c.Origin()
}

// Draws text from location of cursor
//
// Text will not wrap any text greater than terminal width will throw error
func (c *cursor) DrawText(text string) error {
	if c.x+len(text) > TERMINAL.width {
		return fmt.Errorf("text too long for width of page")
	}

	print(text)
	c.x += len(text)

	return nil
}

// Draws text in the center of current link. Keeps original Y cordinate
func (c *cursor) DrawTextCenterX(text string) {
	c.MoveTo(0, c.y)
	c.MoveTo(int(TERMINAL.width/2)-len(text)/2, c.y)

	print(text)
	c.x += len(text)
}

// Draws text in teh center of the height of terminal. Keeps orginal X cordinate
func (c *cursor) DrawTextCenterY(text string) {
	c.MoveTo(c.x, int(TERMINAL.height/2))
	print(text)
	c.x += len(text)
}

// Draws text in the center of the terminal
func (c *cursor) DrawTextCenter(text string) {
	c.MoveTo(int(TERMINAL.width/2)-len(text), int(TERMINAL.height/2))
	print(text)
	c.x += len(text)
}

// Draws text on right side of current line
func (c *cursor) DrawTextRight(text string) {
	c.MoveTo(TERMINAL.width-len(text), c.y)
	print(text)
	c.x += len(text)
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
	print(CLEARLINE)
}

// Hids the cursor
func (c *cursor) Hide() {
	print(HIDE)
}

// Shows the cursor
func (c *cursor) Show() {
	print(SHOW)
}
