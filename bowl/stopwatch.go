package bowl

import (
	"fmt"
	"time"

	"github.com/DustinMeyer1010/ramen/keys"
)

// Default keys for stopwatch
var defaultcfg = stopWatchConfiguration{
	start: keys.KeyOptions{keys.Enter},
	pause: keys.KeyOptions{keys.Space},
	reset: keys.KeyOptions{keys.LowerR},
	exit:  keys.KeyOptions{keys.ControlC, keys.Esc},
}

type stopWatchConfiguration struct {
	start keys.KeyOptions
	pause keys.KeyOptions
	reset keys.KeyOptions
	exit  keys.KeyOptions
}

func NewStopWatchCfg(start, pause, reset, exit keys.KeyOptions) stopWatchConfiguration {
	cfg := defaultcfg

	if start.NotEmpty() {
		cfg.start = start
	}
	if pause.NotEmpty() {
		cfg.pause = pause
	}
	if reset.NotEmpty() {
		cfg.reset = reset
	}

	if exit.NotEmpty() {
		cfg.exit = exit
	}

	return cfg
}

type stopwatch struct {
	start      time.Time
	end        time.Time
	started    bool
	timelapsed time.Duration
}

func NewStopWatch() stopwatch {
	return stopwatch{started: false}
}

// Main loop for stopwatch functionality
func (s *stopwatch) Render(cfg stopWatchConfiguration) {
	s.start = time.Now()
	cur.Hide()
	cur.ClearTerminal()
	if rawMode == -1 {
		startRawMode()
	}
	s.drawControls(cfg)
	s.drawStopWatch()
	s.controlHandler(cfg)
}

func (s *stopwatch) controlHandler(cfg stopWatchConfiguration) {

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
			case cfg.exit.Contains(key):
				cur.ClearTerminal()
				return
			case cfg.start.Contains(key):

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
			case cfg.reset.Contains(key):
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

func (s *stopwatch) drawControls(cfg stopWatchConfiguration) {
	cur.OriginBottom()
	cur.DrawText("Exit: ")
	cur.DrawText(keys.KeyAlias[cfg.exit[0]])
}

func (s *stopwatch) GetTime() {

}
