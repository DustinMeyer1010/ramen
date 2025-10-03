package ramen

import (
	"fmt"
	"time"

	"github.com/DustinMeyer1010/ramen/keys"
)

// Default keys for stopwatch
var defaultStopWatchControls = stopWatchControls{
	start: keys.KeyOptions{keys.Enter},
	pause: keys.KeyOptions{keys.Space},
	reset: keys.KeyOptions{keys.LowerR},
	exit:  keys.KeyOptions{keys.ControlC, keys.Esc},
}

type stopWatchControls struct {
	start keys.KeyOptions
	pause keys.KeyOptions
	reset keys.KeyOptions
	exit  keys.KeyOptions
}

func NewStopWatchControls(start, pause, reset, exit keys.KeyOptions) stopWatchControls {
	controls := defaultStopWatchControls

	if start.HasElements() {
		controls.start = start
	}
	if pause.HasElements() {
		controls.pause = pause
	}
	if reset.HasElements() {
		controls.reset = reset
	}

	if exit.HasElements() {
		controls.exit = exit
	}

	return controls
}

type stopwatch struct {
	start      time.Time
	end        time.Time
	started    bool
	timelapsed time.Duration
	controls   stopWatchControls
}

func NewStopWatch(controls stopWatchControls) stopwatch {
	return stopwatch{started: false}
}

// Main loop for stopwatch functionality
func (s *stopwatch) Render() {
	s.start = time.Now()
	cur.Hide()
	cur.ClearTerminal()
	if rawMode == -1 {
		startRawMode()
	}
	s.drawControls()
	s.drawStopWatch()
	s.controlHandler()
}

func (s *stopwatch) controlHandler() {

	keysChan := keys.NewKeyChannel()

	go func() {
		for {
			key := keys.GetKeyPressed()
			if key != keys.Empty {
				keysChan <- key
			}

		}
	}()

	for {
		if s.started {
			s.drawStopWatch()
		}
		// Key checking in go route to prevent blocking
		select {
		case key := <-keysChan:
			switch {
			case s.controls.exit.Contains(key):
				cur.ClearTerminal()
				return
			case s.controls.start.Contains(key):

				if s.started {
					// functionality when stopwatch is STOPPED
					s.timelapsed += time.Since(s.start)
					s.started = false
					s.end = time.Now()
				} else {
					// functionality when stopwatch is STARTED
					s.start = time.Now()
					s.started = true
				}
			case s.controls.reset.Contains(key):
				s.end = time.Now()
				s.start = time.Now()
				s.timelapsed = 0
				s.drawStopWatch()
			default:
				// ignore other keys
			}
		default:
			// no key ready, continue printing
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// Draws stop watch to screen based on the time passed from when it was started
func (s *stopwatch) drawStopWatch() {
	cur.OriginBottom()
	cur.Up(1)
	elapsed := time.Since(s.start) + s.timelapsed
	cur.DrawText(fmt.Sprintf("%02d:%02d:%02d",
		int(elapsed.Hours()),
		int(elapsed.Minutes())%60,
		int(elapsed.Seconds())%60,
	))
}

func (s *stopwatch) drawControls() {
	cur.OriginBottom()
	cur.DrawText("Exit: ")
	cur.DrawText(keys.KeyAlias[s.controls.exit[0]])
}
