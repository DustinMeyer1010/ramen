package ramen

import (
	"fmt"

	"github.com/DustinMeyer1010/ramen/keys"
)

type Menu struct {
	itemCount        int
	items            []string
	currentSelection int
	controls         MenuControls
}

var defaultMenuControls = MenuControls{
	up:       keys.KeyOptions{keys.UpArrow},
	down:     keys.KeyOptions{keys.DownArrow},
	selector: keys.KeyOptions{keys.Enter},
	exit:     keys.KeyOptions{keys.ControlC, keys.Esc},
}

type MenuControls struct {
	up       keys.KeyOptions
	down     keys.KeyOptions
	selector keys.KeyOptions
	exit     keys.KeyOptions
}

// Create new controls for selection menu
//
// nil or empty values result in default being used
func NewMenuControls(up, down, selector, exit keys.KeyOptions) MenuControls {
	controls := defaultMenuControls

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

func NewMenu(items []string, controls MenuControls) (Menu, error) {
	if len(items) == 0 {
		return Menu{itemCount: 0, items: []string{}}, fmt.Errorf("selection menu must have at least one item")
	}

	return Menu{itemCount: len(items), items: items, currentSelection: 0}, nil
}

func (m *Menu) Render() {
	CURSOR.ClearTerminal()
	CURSOR.Origin()
	for i, item := range m.items {
		CURSOR.DrawText(fmt.Sprintf("%d: %s", i+1, item))
		CURSOR.MoveTo(0, i+1)
	}
	CURSOR.Origin()
	m.controlHandler()
}

func (m *Menu) controlHandler() {
	m.drawCurrentSelection()

	for {
		pressedKey := keys.GetKeyPressed()
		switch pressedKey {
		case keys.ControlC, keys.Esc:
			{
				CURSOR.ClearTerminal()
				return
			}
		case keys.UpArrow:
			m.handleUpSelect()
		case keys.DownArrow:
			m.handleDownSelect()
		case keys.Enter:
			m.currentSelection = CURSOR.GetY()
			CURSOR.ClearTerminal()
			return
		default:
			// ignore all other keys
		}
	}
}

func (m *Menu) handleUpSelect() {
	CURSOR.Up(1)
	m.currentSelection = CURSOR.GetY()
	m.drawCurrentSelection()
}

func (m *Menu) handleDownSelect() {
	if CURSOR.GetY() < m.itemCount-1 {
		CURSOR.Down(1)
		m.currentSelection = CURSOR.GetY()
		m.drawCurrentSelection()
	}
}

func (m *Menu) GetSelectedItem() int {
	return m.currentSelection + 1
}

func (m *Menu) drawCurrentSelection() {
	CURSOR.Hide()
	CURSOR.MoveTo(0, m.itemCount+2)
	CURSOR.ClearLine()
	CURSOR.DrawText(fmt.Sprintf("Selected %d", m.currentSelection+1))
	CURSOR.MoveTo(0, m.currentSelection)
	CURSOR.Show()
}
