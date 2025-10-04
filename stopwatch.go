package ramen

import (
	"fmt"
	"time"

	"github.com/DustinMeyer1010/ramen/keys"
	"github.com/DustinMeyer1010/ramen/terminal"
)

/*
TODO:
1. Render the help menu when pressing the help button
2. Render the logs when logs are press
3. When someone press exit must the let the caller know

BUGS:
1. Fix time reset when other buttons are pressed
*/

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

// Used to create custom controls for the stop watch, if nil or empty array is given for option it will use the default controls
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

// Creates the main stopwatch component which is the entry point to render the stopwatch UI
func NewStopWatch(controls stopWatchControls) stopwatch {
	return stopwatch{isStarted: false, controls: controls, helpIsShown: false, start: time.Now()}
}

// Renders the stopwatch
func (s *stopwatch) Render() ExitCode {
	terminal.NewTerminal(10, -1)
	CURSOR.ClearTerminal()
	CURSOR.Hide()
	s.drawHelpSection()
	s.drawStopWatch()
	return s.eventHandler()
}

// Main loop for handle the events that happen with the stop watch
func (s *stopwatch) eventHandler() ExitCode {

	startEventSources() // Starts all go routes

	for {
		if s.isStarted {
			s.drawStopWatch()
		}
		// Key checking in go route to prevent blocking
		select {
		case key := <-KEYSCHANNEL:
			switch {
			case s.controls.exit.Contains(key):
				CURSOR.ClearTerminal()
				fmt.Print("\033[?1049l")
				return QUIT
			case s.controls.start.Contains(key):
				s.StartStop()
			case s.controls.reset.Contains(key):
				s.reset()
			case s.controls.logtime.Contains(key):
				s.logTime()
			case s.controls.help.Contains(key):
				s.showHelpMenu()
			default:
				// Skip all other keys
			}
		case size := <-RESIZECHANNEL:

			CURSOR.ClearTerminal()
			CURSOR.Origin()
			terminal.TERMINAL.SetDimensions(terminal.TERMINAL.GetHeight(), size[0])
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
	elapsed := time.Since(s.start) + s.timelapsed
	displayTime := fmt.Sprintf("%02d:%02d:%02d",
		int(elapsed.Hours()),
		int(elapsed.Minutes())%60,
		int(elapsed.Seconds())%60,
	)
	CURSOR.DrawTextCenterX(displayTime)
}

// Draws the help section which display to key to press to show help
func (s *stopwatch) drawHelpSection() {
	CURSOR.OriginBottom()
	CURSOR.ClearLine()
	CURSOR.DrawText("Help: ")
	CURSOR.DrawText(string(s.controls.help[0]))
}

// Starts and Stops the stopwatch
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

// Reset the Stopwatch to 0
func (s *stopwatch) reset() {
	s.end = time.Now()
	s.start = time.Now()
	s.isStarted = false
	s.timelapsed = 0
	s.drawStopWatch()
}

// Logs the current time on the stop watch and will draw a log under the stop watch
func (s *stopwatch) logTime() {
	s.logs = append(s.logs, log{duration: time.Since(s.start) + s.timelapsed, start: time.Now()})
	s.drawLogs()
}

func (s *stopwatch) drawLogs() {
	line := 2
	CURSOR.MoveTo(0, line)
	for i := len(s.logs) - 1; i > 0; i-- {
		if line >= terminal.TERMINAL.GetHeight()-2 {
			break
		}
		line += 1
		CURSOR.DrawTextCenterX(s.logs[i].toString())
		CURSOR.MoveTo(0, line)
	}
}

func (l *log) toString() string {
	displayTime := fmt.Sprintf("%02d:%02d:%02d",
		int(l.duration.Hours()),
		int(l.duration.Minutes())%60,
		int(l.duration.Seconds())%60,
	)
	startTime := fmt.Sprintf("%02d:%02d:%02d",
		int(l.start.Hour()),
		int(l.start.Minute()),
		int(l.start.Second()),
	)
	return fmt.Sprintf("Duration: %s\tStart Time: %s", displayTime, startTime)
}

// Draws all the controls
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
