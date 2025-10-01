package keys

import (
	"os"
)

type Key string

var (
	ControlC   Key = "\x1b"
	Esc        Key = "\x03"
	LowerQ     Key = "q"
	UpperQ     Key = "Q"
	LowerR     Key = "r"
	Space      Key = " "
	UpArrow    Key = "\033[A"
	DownArrow  Key = "\033[B"
	RightArrow Key = "\033[C"
	LeftArrow  Key = "\033[D"
	ClearLine  Key = "\033[2K"
	Enter      Key = "\r"
)

func GetKeyPressed() Key {
	var buf = make([]byte, 3)
	n, err := os.Stdin.Read(buf)

	if err != nil {
		panic(err)
	}

	//fmt.Print(string(buf[:n]))

	if n == 0 {
		return Key("")
	}

	return Key(buf[:n])
}
