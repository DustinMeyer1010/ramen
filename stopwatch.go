package ramen

import (
	"fmt"
	"time"

	"github.com/DustinMeyer1010/ramen/keys"
	"github.com/DustinMeyer1010/ramen/terminal"
)

// Default keys for stopwatch
var DefaultStopWatchControls = stopWatchControls{
	start:   keys.KeyOptions{keys.Enter},
	pause:   keys.KeyOptions{keys.Space},
	reset:   keys.KeyOptions{keys.LowerR},
	logtime: keys.KeyOptions{keys.LowerL},
	help:    keys.KeyOptions{keys.LowerH},
	exit:    keys.KeyOptions{keys.ControlC, keys.Esc},
}

type log struct {
	start    time.Time
	end      time.Time
	duration time.Duration
}

type stopWatchControls struct {
	start   keys.KeyOptions
	pause   keys.KeyOptions
	reset   keys.KeyOptions
	logtime keys.KeyOptions
	exit    keys.KeyOptions
	help    keys.KeyOptions
}

func NewStopWatchControls(start, pause, reset, exit, logtime, help keys.KeyOptions) stopWatchControls {
	controls := DefaultStopWatchControls

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
	if logtime.HasElements() {
		controls.logtime = logtime
	}
	if help.HasElements() {
		controls.help = help
	}

	return controls
}

type stopwatch struct {
	start       time.Time
	end         time.Time
	isStarted   bool
	timelapsed  time.Duration
	logs        []log
	controls    stopWatchControls
	helpIsShown bool
}

func NewStopWatch(controls stopWatchControls) stopwatch {
	return stopwatch{isStarted: false, controls: controls, helpIsShown: false}
}

// Main loop for stopwatch functionality
func (s *stopwatch) Render() {
	terminal.NewTerminal(10, -1)
	s.start = time.Now()
	CURSOR.ClearTerminal()
	CURSOR.Hide()
	s.drawHelpSection()
	s.drawStopWatch()
	s.controlHandler()
}

func (s *stopwatch) controlHandler() {

	resizeChan := make(chan [2]int)
	keysChan := keys.NewKeyChannel()
	go terminal.ResizeCheck(resizeChan)
	go keys.KeyChannel(keysChan)

	for {
		if s.isStarted {
			s.drawStopWatch()
		}
		// Key checking in go route to prevent blocking
		select {
		case key := <-keysChan:
			switch {
			case s.controls.exit.Contains(key):
				CURSOR.ClearTerminal()
				fmt.Print("\033[?1049l")
				return
			case s.controls.start.Contains(key):
				s.StartStop()
			case s.controls.reset.Contains(key):
				s.reset()
			case s.controls.logtime.Contains(key):
				s.logTime()
			case s.controls.help.Contains(key):
				s.showHelpMenu()
			default:
				// ignore other keys
			}
		case size := <-resizeChan:

			CURSOR.ClearTerminal()
			CURSOR.Origin()
			terminal.TERMINAL.SetDimensions(terminal.TERMINAL.GetHeight(), size[0])
			time.Sleep(time.Second * 1)
			s.drawHelpSection()
			s.drawStopWatch()

		default:
			// no key ready, continue printing
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// Draws stop watch to screen based on the time passed from when it was started
func (s *stopwatch) drawStopWatch() {

	CURSOR.Origin()
	CURSOR.ClearLine()
	elapsed := time.Since(s.start) + s.timelapsed
	displayTime := fmt.Sprintf("%02d:%02d:%02d",
		int(elapsed.Hours()),
		int(elapsed.Minutes())%60,
		int(elapsed.Seconds())%60,
	)
	CURSOR.DrawTextCenterX(displayTime)
}

func (s *stopwatch) drawHelpSection() {
	CURSOR.OriginBottom()
	CURSOR.ClearLine()
	CURSOR.DrawText("Help: ")
	CURSOR.DrawText(string(s.controls.help[0]))
}

func (s *stopwatch) StartStop() {
	if s.isStarted {
		// functionality when stopwatch is STOPPED
		s.timelapsed += time.Since(s.start)
		s.isStarted = false
		s.end = time.Now()
	} else {
		// functionality when stopwatch is STARTED
		s.start = time.Now()
		s.isStarted = true
	}
}

func (s *stopwatch) reset() {
	s.end = time.Now()
	s.start = time.Now()
	s.timelapsed = 0
	s.drawStopWatch()
}

func (s *stopwatch) logTime() {
	i := 0
	CURSOR.MoveTo(0, 3)
	i += 1
	CURSOR.DrawText(fmt.Sprintf("%d", i))
}

func (s *stopwatch) showHelpMenu() {
	if s.helpIsShown {
		s.start = time.Now()
		s.drawHelpSection()
		s.drawStopWatch()
	} else {
		CURSOR.Origin()
		CURSOR.ClearTerminal()
	}
	s.helpIsShown = !s.helpIsShown

}
