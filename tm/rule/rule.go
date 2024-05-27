package rule

import (
	"fmt"
)

type Rule struct {
	State string
	Match rune
	Next  string
	Write rune
	Move  string
}

func New(st string, ma rune, ne string, wr rune, mo string) Rule {
	r := Rule{
		State: st,
		Match: ma,
		Next:  ne,
		Write: wr,
		Move:  mo,
	}
	if !r.IsMoveLeft() && !r.IsMoveRight() {
		panic("unknown move direction")
	}
	return r
}

func (r Rule) Key() string {
	return CalculateKey(r.State, r.Match)
}

func CalculateKey(state string, match rune) string {
	return fmt.Sprintf("%v/%v", state, match)
}

func (r Rule) IsMoveLeft() bool {
	switch r.Move {
	case "left", "LEFT", "Left", "L":
		return true
	default:
		return false
	}
}

func (r Rule) IsMoveRight() bool {
	switch r.Move {
	case "right", "RIGHT", "Right", "R":
		return true
	default:
		return false
	}
}
