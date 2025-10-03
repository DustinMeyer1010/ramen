package keys

import (
	"os"
)

type key string

var (
	ControlC   key = "\x1b"
	Esc        key = "\x03"
	LowerQ     key = "q"
	UpperQ     key = "Q"
	LowerR     key = "r"
	Space      key = " "
	UpArrow    key = "\033[A"
	DownArrow  key = "\033[B"
	RightArrow key = "\033[C"
	LeftArrow  key = "\033[D"
	Enter      key = "\r"
	Empty      key = ""
)

var KeyAlias map[key]string = map[key]string{
	ControlC:   "Control-C",
	Esc:        "Escape",
	LowerQ:     "q",
	UpperQ:     "Q",
	LowerR:     "r",
	UpArrow:    "↑",
	RightArrow: "→",
	LeftArrow:  "←",
	DownArrow:  "↓",
	Enter:      "Enter",
	Empty:      "Empty-Character",
}

func NewKeyChannel() chan key {
	return make(chan key)
}

func GetKeyPressed() key {
	var buf = make([]byte, 3)
	n, err := os.Stdin.Read(buf)

	if err != nil {
		panic(err)
	}

	//fmt.Print(string(buf[:n]))

	if n == 0 {
		return key("")
	}

	return key(buf[:n])
}

type KeyOptions []key

func (k KeyOptions) Contains(pressedKey key) bool {
	for _, key := range k {
		if key == pressedKey {
			return true
		}
	}
	return false
}

func (k KeyOptions) IsEmpty() bool {
	return k == nil
}

func (k KeyOptions) HasElements() bool {
	return k != nil
}

func NewKeyOptions(keys ...key) KeyOptions {
	return KeyOptions(keys)
}
