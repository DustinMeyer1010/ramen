package terminal

import (
	"fmt"
	"os"
	"time"

	"github.com/DustinMeyer1010/Ramen/keys"
)

type keyOptions []keys.Key

func (k keyOptions) contains(pressedKey keys.Key) bool {
	for _, key := range k {
		if key == pressedKey {
			return true
		}
	}
	return false
}

func (k keyOptions) isEmpty() bool {
	return len(k) == 0
}

func (k keyOptions) notEmpty() bool {
	return len(k) > 0
}

func NewKeyOptions(keys ...keys.Key) keyOptions {
	return keyOptions(keys)
}

// Default keys for stopwatch
var defaultcfg = stopWatchConfiguration{
	start: keyOptions{keys.Enter},
	pause: keyOptions{keys.Space},
	reset: keyOptions{keys.LowerR},
	exist: keyOptions{keys.ControlC, keys.Esc},
}

type stopWatchConfiguration struct {
	start keyOptions
	pause keyOptions
	reset keyOptions
	exist keyOptions
}

func NewStopWatchCfg(start, pause, reset, exist keyOptions) stopWatchConfiguration {
	cfg := defaultcfg

	if !start.notEmpty() {
		cfg.start = start
	}
	if !pause.notEmpty() {
		cfg.pause = pause
	}
	if !reset.notEmpty() {
		cfg.reset = reset
	}
	if !exist.notEmpty() {
		cfg.exist = exist
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
func (s *stopwatch) Draw(cfg stopWatchConfiguration) {
	s.start = time.Now()
	cur.Hide()
	cur.ClearTerminal()
	if rawMode == -1 {
		startRawMode()
	}

	s.controlHandler(cfg)
}

func (s *stopwatch) controlHandler(cfg stopWatchConfiguration) {

	keysChan := make(chan keys.Key)

	go func() {
		for {
			key := keys.GetKeyPressed()
			if key != keys.Key("") {
				keysChan <- key
			}

		}
	}()
	s.drawstopwatch()
	for {
		if s.started {
			s.drawstopwatch()
		}
		// Key checking in go route to prevent blocking
		select {
		case key := <-keysChan:
			switch {
			case cfg.exist.contains(key):
				cur.ClearTerminal()
				os.Exit(1)
			case cfg.start.contains(key):

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
			case cfg.reset.contains(key):
				s.end = time.Now()
				s.start = time.Now()
				s.timelapsed = 0
				s.drawstopwatch()
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
func (s *stopwatch) drawstopwatch() {
	cur.Origin()
	elapsed := time.Since(s.start) + s.timelapsed
	cur.DrawText(fmt.Sprintf("\rElapsed time: %02d:%02d:%02d",
		int(elapsed.Hours()),
		int(elapsed.Minutes())%60,
		int(elapsed.Seconds())%60,
	))
}

func (s *stopwatch) GetTime() {

}
