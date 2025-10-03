package ramen

import (
	"fmt"

	"github.com/DustinMeyer1010/ramen/keys"
)

type selectionMenu struct {
	itemCount        int
	items            []string
	currentSelection int
	controls         selectionMenuControls
}

var defaultSelectMenuControls = selectionMenuControls{
	up:       keys.KeyOptions{keys.UpArrow},
	down:     keys.KeyOptions{keys.DownArrow},
	selector: keys.KeyOptions{keys.Enter},
	exit:     keys.KeyOptions{keys.ControlC, keys.Esc},
}

type selectionMenuControls struct {
	up       keys.KeyOptions
	down     keys.KeyOptions
	selector keys.KeyOptions
	exit     keys.KeyOptions
}

// Create new controls for selection menu
//
// nil or empty values result in default being used
func NewSelectionMenuControls(up, down, selector, exit keys.KeyOptions) selectionMenuControls {
	controls := defaultSelectMenuControls

	if up.HasElements() {
		controls.up = up
	}
	if down.HasElements() {
		controls.down = down
	}
	if selector.HasElements() {
		controls.selector = selector
	}
	if exit.HasElements() {
		controls.exit = exit
	}

	return controls
}

func NewSelectionMenu(items []string, controls selectionMenuControls) (selectionMenu, error) {
	if len(items) == 0 {
		return selectionMenu{itemCount: 0, items: []string{}}, fmt.Errorf("selection menu must have at least one item")
	}

	return selectionMenu{itemCount: len(items), items: items, currentSelection: 0}, nil
}

func (s *selectionMenu) Render() {
	cur.ClearTerminal()
	if rawMode == -1 {
		startRawMode()
	}
	cur.Origin()
	for i, item := range s.items {
		cur.DrawText(fmt.Sprintf("%d: %s", i+1, item))
		cur.MoveTo(0, i+1)
	}
	cur.Origin()
	s.controlHandler()
}

func (s *selectionMenu) controlHandler() {
	s.drawCurrentSelection()

	for {
		pressedKey := keys.GetKeyPressed()
		switch pressedKey {
		case keys.ControlC, keys.Esc:
			{
				cur.ClearTerminal()
				return
			}
		case keys.UpArrow:
			s.handleUpSelect()
		case keys.DownArrow:
			s.handleDownSelect()
		case keys.Enter:
			s.currentSelection = cur.GetY()
			cur.ClearTerminal()
			return
		default:
			// ignore all other keys
		}
	}
}

func (s *selectionMenu) handleUpSelect() {
	cur.Up(1)
	s.currentSelection = cur.GetY()
	s.drawCurrentSelection()
}

func (s *selectionMenu) handleDownSelect() {
	if cur.GetY() < s.itemCount-1 {
		cur.Down(1)
		s.currentSelection = cur.GetY()
		s.drawCurrentSelection()
	}
}

func (s *selectionMenu) GetSelectedItem() int {
	return s.currentSelection + 1
}

func (s *selectionMenu) drawCurrentSelection() {
	cur.Hide()
	cur.MoveTo(0, s.itemCount+2)
	cur.ClearLine()
	cur.DrawText(fmt.Sprintf("Selected %d", s.currentSelection+1))
	cur.MoveTo(0, s.currentSelection)
	cur.Show()
}
