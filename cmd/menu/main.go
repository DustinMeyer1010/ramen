package main

import "github.com/DustinMeyer1010/ramen"

func main() {

	s, _ := ramen.NewMenu([]string{"hello", "world", "option"}, ramen.NewMenuControls(nil, nil, nil, nil))
	s.Render()
	/*
		cfg := bowl.NewStopWatchCfg(nil, nil, nil, nil)
		stopwatch := bowl.NewStopWatch()
		stopwatch.Render(cfg)
		bowl.Exit()
	*/
}
