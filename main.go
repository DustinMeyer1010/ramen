package main

import (
	"github.com/DustinMeyer1010/Ramen/terminal"
)

type TermianlState int

var CurrentTerminalState TermianlState = OptionSelectionMenu
var PreviousTerminalState TermianlState = -1

const (
	OptionSelectionMenu TermianlState = iota
	ProjectSelectionMenu
	AnotherState
)

func main() {
	//database.Init()
	//handleTerminalState()
	stopwatch := terminal.NewStopWatch()
	cfg := terminal.NewStopWatchCfg(nil, nil, nil, nil)
	stopwatch.Draw(cfg)
	terminal.Exit()
	//RunTimer()
}

func handleTerminalState() {
	for {
		if CurrentTerminalState == PreviousTerminalState {
			continue
		}
		switch CurrentTerminalState {
		case OptionSelectionMenu:
			{
				PreviousTerminalState = CurrentTerminalState
				selectionMenu, _ := terminal.NewSelectionMenu([]string{"Create Project", "Delete Project"})
				selectionMenu.Draw()

				if selectionMenu.GetSelectedItem() == 1 {
					return
				}
				if selectionMenu.GetSelectedItem() == 2 {
					CurrentTerminalState = ProjectSelectionMenu
				}
			}
		case ProjectSelectionMenu:
			{
				PreviousTerminalState = CurrentTerminalState
				selectionMenu, _ := terminal.NewSelectionMenu([]string{"Project 1", "Project 2", "Project 3", "Project 4", "Project 5"})
				selectionMenu.Draw()
				return
			}
		}
	}
}
