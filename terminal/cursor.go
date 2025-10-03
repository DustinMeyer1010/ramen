package terminal

import (
	"fmt"
	"log"

	"github.com/DustinMeyer1010/ramen/keys"
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
	ClearLine          string     = "\033[2K"
)

var print = fmt.Print

// Create cursor on the screen
func NewCursor(x int, y int) cursor {
	c := cursor{x: x, y: y}
	c.DrawContainer()
	err := terminal.GetDimensions()

	if err != nil {
		log.Fatal("Unable to get terminal Dimensions")
	}
	c.Origin()
	c.DrawContainer()

	return c
}

// Moves cursor Down n times
//
// Cursor will not move outside of configured terminal height
func (c *cursor) Down(n int) {
	for range n {
		if terminal.Height <= c.y {
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

		if terminal.Width <= c.x {
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
	c.Up(terminal.Height)
	c.Left(terminal.Width)
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
	c.Down(terminal.Height)
}

// Reset entire terminal to blank
func (c *cursor) ClearTerminal() {
	x := c.x
	y := c.y
	c.Origin()
	for range terminal.Height {
		print(ClearLine)
		c.Down(1)
	}
	c.MoveTo(x, y)
}

// Draws the height and width of terminal configuration
func (c *cursor) DrawContainer() {
	c.Origin()
	for range 100 {
		c.Down(terminal.Height)
	}
	for range 100 {
		c.Right(terminal.Width)
	}
	c.Origin()
}

// Draws text from location of cursor
//
// Text will not wrap any text greater than terminal width will throw error
func (c *cursor) DrawText(text string) error {
	if c.x+len(text) > terminal.Width {
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
	print(ClearLine)
}

func (c *cursor) Hide() {
	print(HIDE)
}

func (c *cursor) Show() {
	print(SHOW)
}
