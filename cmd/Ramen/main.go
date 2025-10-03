package main

import "github.com/DustinMeyer1010/ramen/ramen"

func main() {
	s, _ := ramen.NewSelectionMenu([]string{"hello", "world", "option"}, ramen.NewSelectionMenuControls(nil, nil, nil, nil))
	s.Render()
	/*
		cfg := bowl.NewStopWatchCfg(nil, nil, nil, nil)
		stopwatch := bowl.NewStopWatch()
		stopwatch.Render(cfg)
		bowl.Exit()
	*/
}
