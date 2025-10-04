package main

import "github.com/DustinMeyer1010/ramen"

func main() {
	stopwatch := ramen.NewStopWatch(ramen.DefaultStopWatchControls)
	stopwatch.Render()
}
